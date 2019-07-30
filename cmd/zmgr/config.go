package main

import (
	"github.com/Urethramancer/cross"
	"github.com/Urethramancer/signor/files"
	"github.com/Urethramancer/signor/log"
)

// Config for server connection.
type Config struct {
	// Host is the domain/UP address part of the address.
	Host string `json:"host"`
	// Port defaults to 389.
	Port string `json:"port"`
	// Password for the admin.
	Password string `json:"password"`
	// LMTPPort is used in configuring the zimbraMailTransport attribute on accounts.
	LMTPPort string `json:"lmtpport"`
}

var configName string

func init() {
	cross.SetConfigPath("zmgr")
	configName = cross.ConfigName("config.json")
}

func loadConfig() *Config {
	cfg := Config{
		Host:     "localhost",
		Port:     "389",
		LMTPPort: "7025",
	}

	if files.FileExists(configName) {
		err := files.LoadJSON(configName, &cfg)
		if err != nil {
			log.Default.Msg("Couldn't load configuration %s: %s", configName, err.Error())
		}
	}

	return &cfg
}

func saveConfig(data interface{}) error {
	return files.SaveJSON(configName, data)
}
