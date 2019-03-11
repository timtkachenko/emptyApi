package lib

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"sync"
	"testing"
	"time"
)

func TestConnect(t *testing.T) {
	a := assert.New(t)
	dbmock, _, err := sqlmock.New()
	if err != nil {
		t.Errorf(err.Error())
	}
	viper.Set("postgres", dbmock)
	onceDB = sync.Once{}
	config = &ConnectionConfig{1, 1}
	a.NotPanics(func() {
		a.NotNil(Connect())
	})
	a.Equal(db, Connect())
}
func TestConnectTimeout(t *testing.T) {
	a := assert.New(t)
	dbmock, _, err := sqlmock.New()
	if err != nil {
		t.Errorf(err.Error())
	}
	viper.Set("postgres", "")
	go func() {
		<-time.After(time.Millisecond)
		viper.Set("postgres", dbmock)
	}()
	onceDB = sync.Once{}
	config = &ConnectionConfig{3, time.Millisecond * 2}
	a.NotPanics(func() {
		a.NotNil(Connect())
	})
}
func TestConnectPanics(t *testing.T) {
	a := assert.New(t)
	viper.Set("postgres", "wrong=")
	onceDB = sync.Once{}
	config = &ConnectionConfig{1, 1}
	a.Panics(func() {
		Connect()
	})
}
