package config

import "github.com/spf13/viper"

type AppConfig struct {
	*viper.Viper
	Config
}

type Config struct {
	Env    string                  `yaml:"env"`
	Log    LogConfig               `json:"log"`
	Server ServerConfig            `yaml:"server"`
	DB     map[string]*DbConfig    `yaml:"db"`
	Redis  map[string]*RedisConfig `yaml:"redis"`
}

type LogConfig struct {
	Level    string `json:"level"`
	Appender string `json:"appender"`
	LogDir   string `json:"log-dir"`
	MaxSize  int    `json:"max-size"`
	MaxDays  int    `json:"max-days"`
}

type ServerConfig struct {
	Listen string `yaml:"listen"`
}

type DbConfig struct {
	Type    string `yaml:"type"`
	Dsn     string `yaml:"dsn"`
	MaxIdle int    `yaml:"max-idle"`
	MaxConn int    `yaml:"max-conn"`
	ShowSQL bool   `yaml:"show-sql"`
}

type RedisConfig struct {
	Addr     string `yaml:"addr"`
	Password string `yaml:"password"`
	DB       int    `yaml:"db"`
}
