package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
)

// Message report
type Message struct {
	ChatID string `json:"chat_id,omitempty"`
	Text   string `json:"text"`
}

// Pass token and sensible APIs through environment variables
const telegramApiBaseUrl string = "https://api.telegram.org/bot"
const telegramApiSendMessage string = "/sendMessage"

var telegramChatID string = os.Getenv("TELEGRAM_CHAT_ID")
var telegramToken string = os.Getenv("TELEGRAM_API_KEY")

func main() {
	m := Message{
		ChatID: telegramChatID,
		Text:   "*Telegram Bot* \nStarted",
	}

	_, err := sendTextToTelegramChat(m.ChatID, m.Text)
	if err != nil {
		log.Printf("Error sending text to telegram chat %s: %s", m.ChatID, err)
	}
}

// sendTextToTelegramChat sends a text message to the Telegram chat identified by its chat Id
func sendTextToTelegramChat(chatId string, text string) (string, error) {

	log.Printf("Sending '%s' text to chat_id: %s", text, chatId)
	response, err := http.PostForm(fmt.Sprintf("%s%s%s", telegramApiBaseUrl, telegramToken, telegramApiSendMessage),
		url.Values{
			"chat_id":    {chatId},
			"text":       {text},
			"parse_mode": {"Markdown"},
		})

	if err != nil {
		log.Printf("error when posting text to the chat: %s", err.Error())
		return "", err
	}
	defer response.Body.Close()

	var bodyBytes, errRead = ioutil.ReadAll(response.Body)
	if errRead != nil {
		log.Printf("error in parsing telegram answer %s", errRead.Error())
		return "", err
	}
	bodyString := string(bodyBytes)
	log.Printf("Body of Telegram Response: %s", bodyString)

	return bodyString, nil
}
