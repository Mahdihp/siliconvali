package config

import (
	"fmt"
	"github.com/spf13/viper"
)

type AppConfiguration struct {
	Server          Server
	DbConfig        Database
	MigrationConfig Migration
}

type Server struct {
	Host string
	Port string
}
type Database struct {
	Host     string
	DbPort   int
	DbName   string
	Username string
	Password string
}
type Migration struct {
	MigrationDir   string
	MigrationTable string
}

func Initialize() AppConfiguration {
	viper.SetConfigName("config")
	viper.AddConfigPath("./config")
	viper.SetConfigType("json")
	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file, %s", err)
	}

	return AppConfiguration{
		Server: Server{
			Host: viper.GetString("Server.Host"),
			Port: viper.GetString("Server.Port"),
		},
		DbConfig: Database{
			Host:     viper.GetString("Database.Host"),
			DbName:   viper.GetString("Database.DbName"),
			DbPort:   viper.GetInt("Database.DbPort"),
			Username: viper.GetString("Database.Username"),
			Password: viper.GetString("Database.Password"),
		},
		MigrationConfig: Migration{
			MigrationDir:   viper.GetString("Migration.MigrationDir"),
			MigrationTable: viper.GetString("Migration.MigrationTable"),
		},
	}
}
