package config

import "github.com/spf13/viper"

func Init()  {
	viper.Set("postgres", "host=127.0.0.1 port=5432 user= password= "+
		"dbname=demo sslmode=disable fallback_application_name=emptyApi")
}