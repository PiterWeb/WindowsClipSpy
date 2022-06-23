package builder

import (
	"io/fs"
	"log"
	"os"

	"github.com/pelletier/go-toml"
)

type Config struct {
	Filename string `json:"filename"`
	Url string `json:"url"`
	Method string `json:"method"`
}

var config Config

func GetConfig() Config {
	
	configFile, err := os.Open("../builder/env.toml")

	if err != nil {
		log.Fatal(err)
	}

	if err := toml.NewDecoder(configFile).Decode(&config); err != nil {
		log.Fatal(err)
	}

	return config

}

func SetConfig(c Config) {
	
	configFile, err := os.OpenFile("./builder/env.toml", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, fs.ModeAppend)

	if err != nil {
		log.Fatal(err)
	}

	if err := toml.NewEncoder(configFile).Encode(c); err != nil {
		log.Fatal(err)
	}

	if err := configFile.Close(); err != nil {
		log.Fatal(err)
	}

}