package tgbot

import (
	"encoding/json"
	"log"
	"os"
	"strconv"
	"testing"
)

func bot() *Bot {
	cfg, err := os.Open("config_test.json")
	if err != nil {
		panic(err)
	}
	defer cfg.Close()
	config := new(Config)
	if err := json.NewDecoder(cfg).Decode(config); err != nil {
		panic(err)
	}
	os.Setenv("HTTPS_PROXY", "socks5://127.0.0.1:1080")
	return New(config)
}
func TestGetMe(t *testing.T) {
	me := new(User)
	err := bot().Run(&GetMe{}, me)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(me)
}
func TestUpdate(t *testing.T) {
	var msgs []Update
	err := bot().Run(&GetUpdates{Timeout: 30}, &msgs)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(msgs)
}
func TestSend(t *testing.T) {
	bot := bot()
	var msgs []Update
	err := bot.Run(&GetUpdates{Timeout: 30}, &msgs)
	if err != nil {
		log.Println(err)
		return
	}
	msg := new(Message)
	err = bot.Run(&SendMessage{ChatId: strconv.Itoa(msgs[0].Message.Chat.Id), Text: "hello tgbot"}, msg)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(msg)
}
