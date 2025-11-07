package config

import (
	"os"

	"github.com/pelletier/go-toml/v2"
)

type Server struct {
	Host string `toml:"host"`
	Port int    `toml:"port"`
}

type Config struct {
	Server Server `toml:"server"`
}

func readConfigFile(filename string) []byte {
	data, err := os.ReadFile(filename)
	if err != nil {
		panic(err.Error())
	}
	return data
}

func GetConfig(filename string) *Config {
	var cfg Config
	err := toml.Unmarshal(readConfigFile(filename), &cfg)
	if err != nil {
		panic(err.Error())
	}
	return &cfg
}
