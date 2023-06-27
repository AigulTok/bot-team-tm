package delivery

import (
	"bot-team-tm/internal/service"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"gopkg.in/telebot.v3"
)

type Handler struct {
	service *service.Service
	bot     *telebot.Bot
}

type HandlerBot struct {
	service  *service.Service
	msg      chan string
	commands []string
}

func NewHandlerBot(s *service.Service) *HandlerBot {
	return &HandlerBot{
		service:  s,
		msg:      make(chan string, 1),
		commands: []string{"/start"},
	}
}

func NewHandlerAPI(s *service.Service, bot *telebot.Bot) *Handler {
	return &Handler{
		service: s,
		bot:     bot,
	}
}

func (hb *HandlerBot) InitRoutesBot(bot *telebot.Bot) {
	bot.Handle("/start", hb.StartHandler)
}

func (h *Handler) InitRoutesAPI() {
	router := mux.NewRouter()

	router.HandleFunc("/health", h.sendStatus)
	router.HandleFunc("/sendMessage", h.sendMesssage)

	go func() {
		err := http.ListenAndServe(":8080", router)
		if err != nil {
			log.Fatalf("Не удалось запустить сервер. Ошибка: %v", err)
		}
	}()
}
