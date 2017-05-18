// Package main provides
package main

import (
	"fmt"
	"os"

	"./utils"
	"./utils/password"

	"gopkg.in/telegram-bot-api.v4"
)

// SalesKeyboard the main keyboard with buttons for listing status of things
var SalesKeyboard = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Routers"),
		tgbotapi.NewKeyboardButton("Switches"),
		tgbotapi.NewKeyboardButton("Wires"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("close"),
		tgbotapi.NewKeyboardButton("end"),
	),
)

// this is the main
func main() {

	var helpMessage = "\n/start 			 : starts the bot\n	/help 			: to see the CommandArguments\n/login PASSWORD  : to gain admin access"

	tgbot := os.Getenv("TGBOT")

	bot, _ := tgbotapi.NewBotAPI(tgbot)

	// newbot.Debug = true

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 30

	updates, err := bot.GetUpdatesChan(u)
	utils.Check(err)

	// this is where everything happens      main loop
	for update := range updates {
		if update.Message == nil {
			continue
		}
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)

		// what every Button does
		switch update.Message.Text {
		case "open":
			msg.ReplyMarkup = SalesKeyboard
		case "close":
			msg.ReplyMarkup = tgbotapi.NewHideKeyboard(true)
			// msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
		case "end":
			msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
		default:
			msg.Text = "Sorry unknown commmand :("
		}

		// checking if we received a  Document
		// if update.Message.Document.FileID != "" {
		// 	// Download it

		// 	msgtime := update.Message.Time()
		// 	url, _ := bot.GetFileDirectURL(update.Message.Document.FileID)
		// 	msg.Text = "got a doc at : " + msgtime.Format("Mon Jan 2 15:04:05 MST 2006") + "\n" + url
		// utils.url2File(url, update.Message.Document.FileName)

		// }

		if update.Message.IsCommand() {
			switch update.Message.Command() {
			case "start":
				msg.Text = "Welcome to our shop!"
				msg.ReplyMarkup = SalesKeyboard
			case "login":
				tgpass := update.Message.CommandArguments()
				fmt.Printf("password entered: = %+v\n", tgpass)
				password.Pass_checker(&msg.Text, tgpass)

			case "help":
				msg.Text = helpMessage
			}
		}
		// msg.Text = "grrr XD"
		bot.Send(msg)
	}
}
