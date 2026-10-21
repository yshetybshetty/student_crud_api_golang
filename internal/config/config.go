package config

import (
	"flag"
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type HttpServerConfig struct {
	Addr string `yaml:"adress" env-required:"true"`
}

type Config struct {
	Env              string `yaml:"env" env:"ENV" env-required:"true"`
	StoragePath      string `yaml:"storage-path"`
	HttpServerConfig `yaml:"http_server"`
}

func MustLoad() *Config {
	var configPath string

	configPath = os.Getenv("CONFIG_PATH")

	if configPath == "" {
		flags := flag.String("config", "", "path for the configuration files")
		flag.Parse()

		configPath = *flags

		if configPath == "" {
			log.Fatal("config path not found")
		}
	}

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("path dosenot exist: %s", configPath)
	}

	var config Config

	err := cleanenv.ReadConfig(configPath, &config)

	if err != nil {
		log.Fatalf("can't read config file %s", err.Error())
	}

	return &config
}
