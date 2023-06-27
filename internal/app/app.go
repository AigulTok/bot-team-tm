package app

import (
	"log"
	"time"

	"bot-team-tm/internal/config"
	delivery "bot-team-tm/internal/delivery"
	"bot-team-tm/internal/repository"
	"bot-team-tm/internal/service"
	"bot-team-tm/models"

	"gopkg.in/telebot.v3"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type App struct {
	config *config.Config
	bot    *telebot.Bot
}

func NewApp(cfg *config.Config) *App {
	log.Println("init telebot")

	pref := telebot.Settings{
		Token:  cfg.BotToken,
		Poller: &telebot.LongPoller{Timeout: 10 * time.Second},
	}

	b, err := telebot.NewBot(pref)
	if err != nil {
		log.Fatalf("Не удалось инициализировать бота. Ошибка %s", err)
	}

	return &App{
		config: cfg,
		bot:    b,
	}
}

func (a *App) Run() {
	log.Println("init db")

	db, err := gorm.Open(sqlite.Open(a.config.DB), &gorm.Config{})
	if err != nil {
		log.Fatalf("Не удалось инициализировать базу данных. Ошибка %s", err)
	}

	if err := db.AutoMigrate(models.User{}); err != nil {
		log.Fatalf("Не удалось мигрировать модель в базу. Ошибка %s", err)
	}

	repository := repository.NewRepository(db)
	service := service.NewService(repository)
	handlerBot := delivery.NewHandlerBot(service)

	log.Println("init endpoints for bot")
	handlerBot.InitRoutesBot(a.bot)

	log.Println("start routing")
	handler := delivery.NewHandlerAPI(service, a.bot)
	handler.InitRoutesAPI()

	log.Println("start bot")
	a.bot.Start()

}
