package delivery

import (
	"bot-team-tm/models"
	"fmt"

	"gopkg.in/telebot.v3"
)

func (hb *HandlerBot) StartHandler(c telebot.Context) error {
	user := models.User{
		Name:       c.Sender().Username,
		TelegramId: c.Sender().ID,
		FirstName:  c.Sender().FirstName,
		LastName:   c.Sender().LastName,
		ChatId:     c.Chat().ID,
	}

	if err := hb.service.User.AddNewUser(user); err != nil {
		return fmt.Errorf("не удалось добавить пользователя. Ошибка %s", err)

	}

	return c.Send(fmt.Sprintf("Привет, %s. Список всех команд: /start", user.FirstName))
}
