package config

import (
	"flag"
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type HttpServer struct {
	Adr string `yaml:"address"`
}
type Config struct {
	ENV         string `yaml:"env" env:"ENV" env-required:"true"`
	StoragePath string `yaml:"storage_path" env-required:"true"`
	HttpServer  `yaml:"http_server"`
}

func MustLoad() *Config{
	var configPath string

	configPath = os.Getenv("CONFIG_PATH")

	if (configPath) == "" {
		flags := flag.String("config", "", "path to config file")
		flag.Parse()

		configPath = *flags
		if configPath == "" {
			log.Fatal("Config Path Is Not Set")
		}

	}

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("config file does not exist %s ", configPath)
	}

	var cfg Config 

	err := cleanenv.ReadConfig(configPath , &cfg)

	if err!=nil {
		log.Fatalf("cannot read config file %s" , err.Error())
	}

	return &cfg

}
