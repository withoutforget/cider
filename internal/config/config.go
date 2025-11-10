package config

import (
	"bytes"
	"encoding/json"
	"fmt"
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

type Postgres struct {
	Host     string `toml:"host"`
	Port     int    `toml:"port"`
	Username string `toml:"username"`
	Password string `toml:"password"`
	Database string `toml:"database"`
	Driver   string `toml:"driver"`
}

func (p *Postgres) Dsn() string {
	url := "host=%v port=%v dbname=%v user=%v password=%v sslmode=disable"
	return fmt.Sprintf(url, p.Host, p.Port, p.Database, p.Username, p.Password)
}

type Logging struct {
	HumanReadable bool   `toml:"human_readable"`
	Level         string `toml:"level"`
}

type Session struct {
	Timeout int `toml:"timeout"`
}

type Config struct {
	Server   *Server   `toml:"server"`
	Postgres *Postgres `toml:"postgres"`
	Logging  *Logging  `toml:"logging"`
	Session  *Session  `toml:"session"`
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
