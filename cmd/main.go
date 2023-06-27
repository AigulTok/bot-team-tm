package main

import (
	"bot-team-tm/internal/app"
	"bot-team-tm/internal/config"
	"log"
)

func main() {
	log.Println("start prog")

	cfg, err := config.GetConfig("./config.json")
	if err != nil {
		log.Fatalf("Не удалось получить конфигурации приложенияю Ошибка %s", err)
	}

	app := app.NewApp(cfg)

	app.Run()
}
