package config

import (
	"flag"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Env      string        `yaml:"env" env-requered:"true"`
	TokenTTL time.Duration `yaml:"token_ttl" env-requered:"true"`
	Database struct {
		Host     string `yaml:"host" env-requered:"true"`
		Port     int    `yaml:"port" env-requered:"true"`
		User     string `yaml:"user" env-requered:"true"`
		Password string `yaml:"password" env-requered:"true"`
		DBName   string `yaml:"dbname" env-requered:"true"`
	} `yaml:"database" env-requered:"true"`
	GRPC struct {
		Port    int           `yaml:"port" env-requered:"true"`
		Timeout time.Duration `yaml:"timeout" env-requered:"true"`
	} `yaml:"grpc" env-requered:"true"`
}

func MustLoadConfig() *Config {
	path := fetchConfigPath()

	if path == "" {
		panic("config path is empty")
	}

	if _, err := os.Stat(path); os.IsNotExist(err) {
		panic("config file does not exist: " + path)
	}

	var cfg Config

	if err := cleanenv.ReadConfig(path, &cfg); err != nil {
		panic("failed to read config: " + err.Error())
	}

	return &cfg
}

func fetchConfigPath() string {
	var res string

	flag.StringVar(&res, "config", "", "path to config file")
	flag.Parse()

	if res == "" {
		res = os.Getenv("CONFIG_PATH")
	}

	return res
}
