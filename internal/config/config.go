package config

import (
	"encoding/json"
	"log"
	"os"
)

type Config struct {
	DbUrl string `json:"db_url"`
	CurrUser string `json:"current_user_name"`
}

const configFileName = "/.gatorconfig.json"

func Read() *Config {
	raw, err := os.ReadFile(getConfigPath())
	if err != nil {
		log.Fatal(err)
	}

	var cfg Config
	err = json.Unmarshal(raw, &cfg)
	if err != nil {
		log.Fatal(err)
	}

	return &cfg
}

func (c *Config)SetUser(usr string) {
	c.CurrUser = usr
	write(c)
}

func getConfigPath() string {
	homePath, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}
	return homePath+configFileName
}

func write(cfg *Config) error {
	raw, err := json.Marshal(*cfg)
	if err != nil {
		return err
	}
	os.WriteFile(getConfigPath(), raw, 0777)
	return nil
}
