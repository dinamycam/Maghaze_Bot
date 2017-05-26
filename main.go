// Package main provides
package main

import (
	"fmt"
	"os"

	// "./utils"
	// "./utils/password"
	"github.com/pouya-barzegar/Maghaze_Bot/utils"
	"github.com/pouya-barzegar/Maghaze_Bot/utils/password"

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

	authorize_admin := false
	IsDocument := false

	var helpMessage = `you can use these commands to control this bot
						/start  starts the bot
						/help  to see the CommandArguments
						/login PASSWORD  to gain admin access
						/logout turn back to a normal user
						/senddoc send the files to the bot in the next message`

	tgbot := os.Getenv("TGBOT")
	data_dir := os.Getenv("TGBOTDATA")

	bot, _ := tgbotapi.NewBotAPI(tgbot)

	// bot.Debug = true

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
		case "Routers":
			// msg.Text = utils.Excel2str(data_dir + "routers.xlsx")
			msg.Text = utils.Excel2str("routers.xlsx", data_dir)
			fmt.Println(msg.Text)
		case "Switches":
			msg.Text = utils.Excel2str("switches.xlsx", data_dir)
			fmt.Println(msg.Text)
		case "Wires":
			msg.Text = utils.Excel2str("wires.xlsx", data_dir)
			fmt.Println(msg.Text)
		default:
			if IsDocument {
				msgtime := update.Message.Time()
				url, _ := bot.GetFileDirectURL(update.Message.Document.FileID)
				msg.Text = "got a doc at : " + msgtime.Format("Mon Jan 2 15:04:05 MST 2006") + "\n" + url
				utils.Url2File(url, data_dir+"/"+update.Message.Document.FileName)

				IsDocument = false
			}
		}

		if update.Message.IsCommand() {
			switch update.Message.Command() {
			case "start":
				msg.Text = "Welcome to our shop!"
				fmt.Println("data directory: " + data_dir)
				msg.ReplyMarkup = SalesKeyboard
			case "login":
				tgpass := update.Message.CommandArguments()
				fmt.Printf("password entered: = %+v\n", tgpass)
				authorize_admin = password.Pass_checker(&msg.Text, tgpass)
			case "senddoc":
				if authorize_admin {
					msg.Text = "You may Now send the doc"
					IsDocument = true

				} else {
					msg.Text = "You should first try to /login ."
				}
			case "logout":
				authorize_admin = false
				msg.Text = "you logged out\n switch to normal user"
			case "help":
				msg.Text = helpMessage
			}
		}
		bot.Send(msg)
	}
}
