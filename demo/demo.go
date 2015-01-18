package main

import (
	"flag"

	"github.com/localhots/confection"
)

type (
	Config struct {
		AppName        string         `json:"app_name" attrs:"required" title:"Application Name"`
		BuildNumber    int            `json:"build_number" attrs:"readonly" title:"Build Number"`
		EnableSignIn   bool           `json:"enable_sign_in" title:"Enable Sign-In"`
		DatabaseDriver string         `json:"database_driver" title:"Database Driver" options:"mysql,postgresql,mssql"`
		DatabaseConfig DatabaseConfig `json:"database_config"`
	}
	DatabaseConfig struct {
		Hostname string `json:"hostname"`
		Port     int    `json:"port"`
		Username string `json:"username"`
		Password string `json:"password"`
		Database string `json:"database" attrs:"required"`
	}
)

func main() {
	confection.Setup()
	flag.Parse()

	conf := Config{
		DatabaseConfig: DatabaseConfig{},
	}
	manager := confection.New(conf)
	manager.StartServer()
}
