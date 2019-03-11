package lib

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"log"
	"github.com/spf13/viper"
	"sync"
	"time"
)

type ConnectionConfig struct {
	retryCount   int
	retryTimeOut time.Duration
}

var (
	onceDB sync.Once
	db     *gorm.DB
	config = &ConnectionConfig{1, time.Second * 1}
)

func Connect() *gorm.DB {
	onceDB.Do(dbConnect)
	return db
}

func dbConnect() {
	tryCount := 1
	for {
		log.Println("Connecting to DB...")
		var err error
		db, err = gorm.Open("postgres", viper.Get("postgres"))
		if err != nil {
			log.Println("dbConnect failed: " + err.Error())
			db = nil
		} else if db != nil {
			db.LogMode(true)
			db.DB().SetMaxOpenConns(90)
			return
		}
		if tryCount++; tryCount > config.retryCount {
			panic(err)
		}
		<-time.After(config.retryTimeOut)
	}
}
