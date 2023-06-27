package config

import (
	"flag"
	"log"

	"github.com/BurntSushi/toml"
)

type Config struct {
	Env      string
	BotToken string
	DB       string
}

func GetConfig(path string) (*Config, error) {
	log.Println("init configs")

	configPath := flag.String("config", "", "Path to config file")
	flag.Parse()

	cfg := &Config{}
	_, err := toml.DecodeFile(*configPath, cfg)

	if err != nil {
		log.Fatalf("Ошибка декодирования файла конфигураций. Ошибка %v", err)
	}
	
	return cfg, nil
}
