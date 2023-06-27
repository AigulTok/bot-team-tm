package delivery

import (
	"bot-team-tm/models"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"gopkg.in/telebot.v3"
)

const line = "============================="

func (h *Handler) sendStatus(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "OK")
}

func (h *Handler) sendMesssage(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	fmt.Println(line)
	log.Println("Получаем данные с формы...")

	var mes models.Message
	if err := json.NewDecoder(r.Body).Decode(&mes); err != nil {
		log.Printf("Не удалось получить данные. Ошибка %s", err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		fmt.Println(line)
		return
	}

	fmt.Fprint(w, "Форма была получена успешно!\n")

	chatsID, err := h.service.GetAllChatID()
	if err != nil {
		log.Printf("Не удалось получить ID всех чатов. Ошибка %s", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		fmt.Println(line)
		return
	}

	log.Printf("Отправляем сообщение...")

	for i := 0; i < len(chatsID); i++ {
		group := telebot.ChatID(chatsID[i])
		h.bot.Send(group, mes.Text)
	}

	log.Printf("Сообщения успешно отправлены...")
	fmt.Println(line)

	fmt.Fprint(w, "Отправка прошла успешна!\n")
}
