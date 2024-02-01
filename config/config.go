package config

import (
	"fmt"
	"github.com/spf13/viper"
)

type AppConfiguration struct {
	Server          Server
	DbConfig        DbConfig
	MigrationConfig Migration
}

type Server struct {
	Host string
	Port string
}
type DbConfig struct {
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
		DbConfig: DbConfig{
			Host:     viper.GetString("DbConfig.Host"),
			DbName:   viper.GetString("DbConfig.DbName"),
			DbPort:   viper.GetInt("DbConfig.DbPort"),
			Username: viper.GetString("DbConfig.Username"),
			Password: viper.GetString("DbConfig.Password"),
		},
		MigrationConfig: Migration{
			MigrationDir:   viper.GetString("Migration.MigrationDir"),
			MigrationTable: viper.GetString("Migration.MigrationTable"),
		},
	}
}
