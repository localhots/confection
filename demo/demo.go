package main

import (
	"encoding/json"
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
		SensitiveData  string         `json:"sensitive_data" attrs:"ignored"`
	}
	DatabaseConfig struct {
		Hostname string `json:"hostname"`
		Port     int    `json:"port"`
		Username string `json:"username"`
		Password string `json:"password"`
		Database string `json:"database" attrs:"required"`
	}
)

func init() {
	confection.SetupFlags()
	flag.Parse()
}

func main() {
	conf := Config{
		DatabaseConfig: DatabaseConfig{},
	}
	manager := confection.New(conf, func(b []byte) interface{} {
		var newConf Config
		if err := json.Unmarshal(b, &newConf); err != nil {
			panic(err)
		}
		return newConf
	})
	manager.StartServer()
}
