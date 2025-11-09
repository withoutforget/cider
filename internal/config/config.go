package config

import (
	"bytes"
	"encoding/json"
	"os"

	"github.com/pelletier/go-toml/v2"
)

type Server struct {
	Host             string   `toml:"host"`
	Port             int      `toml:"port"`
	AllowAllOrigins  bool     `toml:"allow_origins"`
	AllowOrigins     []string `toml:"origins"`
	AllowMethods     []string `toml:"methods"`
	AllowHeaders     []string `toml:"headers"`
	AllowCredentials bool     `toml:"allow_credentials"`
	AllowWildcard    bool     `toml:"allow_wildcard"`
	AllowWebSockets  bool     `toml:"allow_websockets"`
}

type Logging struct {
	HumanReadable bool   `toml:"human_readable"`
	Level         string `toml:"level"`
}

type Session struct {
	Timeout int `toml:"timeout"`
}

type Config struct {
	Server  *Server  `toml:"server"`
	Logging *Logging `toml:"logging"`
	Session *Session `toml:"session"`
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

	reader := bytes.NewReader(readConfigFile(filename))
	decoder := toml.NewDecoder(reader)
	decoder = decoder.DisallowUnknownFields()

	err := decoder.Decode(&cfg)

	if err != nil {
		panic(err.Error())
	}
	return &cfg
}

func (c *Config) String() string {
	v, err := json.MarshalIndent(c, "", " ")
	if err != nil {
		return ""
	}
	return string(v)
}
