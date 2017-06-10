// Package main provides
package main

import (
	"fmt"
	"os"

	"Maghaze_Bot/utils"
	"Maghaze_Bot/utils/password"
	// "./utils"
	// "./utils/password"
	// "github.com/pouya-barzegar/Maghaze_Bot/utils"
	// "github.com/pouya-barzegar/Maghaze_Bot/utils/password"

	"gopkg.in/telegram-bot-api.v4"
)

// this is the main
func main() {

	authorize_admin := false
	IsDocument := false

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

		// testing the disable option temp
		// msg.DisableNotification = true
		// what every Button does
		switch update.Message.Text {
		case "open":
			msg.ReplyMarkup = utils.Keyboard_page1
		case "close":
			msg.ReplyMarkup = tgbotapi.NewHideKeyboard(true)
			// msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
		case "بعدی":
			msg.ReplyMarkup = utils.Keyboard_page2
		case "قبلی":
			msg.ReplyMarkup = utils.Keyboard_page1
		case "روتر":
			msg.Text = utils.Tealeg_Excel2str("router.xlsx", data_dir)
			fmt.Println(msg.Text)
		case "سویچ":
			msg.Text = utils.Tealeg_Excel2str("switch.xlsx", data_dir)
			fmt.Println(msg.Text)
		case "کارت شبکه":
			msg.Text = utils.Tealeg_Excel2str("nic.xlsx", data_dir)
			fmt.Println(msg.Text)
		case "مودم":
			msg.Text = utils.Tealeg_Excel2str("modem.xlsx", data_dir)
			fmt.Println(msg.Text)
		case "دوربین":
			msg.Text = utils.Tealeg_Excel2str("camera.xlsx", data_dir)
			fmt.Println(msg.Text)
		case "اکسس پوینت":
			msg.Text = utils.Tealeg_Excel2str("accesspoint.xlsx", data_dir)
			fmt.Println(msg.Text)
		case "تجهیزات رادیویی":
			msg.Text = utils.Tealeg_Excel2str("radio.xlsx", data_dir)
			fmt.Println(msg.Text)
		case "سیسکو":
			msg.Text = utils.Tealeg_Excel2str("cisco.xlsx", data_dir)
			fmt.Println(msg.Text)
		case "هاب و KVM":
			msg.Text = utils.Tealeg_Excel2str("hub&kvm.xlsx", data_dir)
			fmt.Println(msg.Text)
		default:
			if IsDocument {
				fmt.Println("Waiting for a doc file")
				msgtime := update.Message.Time()
				fmt.Println("Getting the url for the message")
				url, err := bot.GetFileDirectURL(update.Message.Document.FileID)
				if err != nil {
					fmt.Println("Didn't get a doc after all")
					continue
				}

				fmt.Println("Got a doc at time:")
				msg.Text = "got a doc at : " + msgtime.Format("Mon Jan 2 15:04:05 MST 2006") + "\n" + url
				utils.Url2File(url, data_dir+"/"+update.Message.Document.FileName)

				IsDocument = false
			}
		}

		if update.Message.IsCommand() {
			switch update.Message.Command() {
			case "start":
				msg.Text = "به شبکه برتر خوش آمدید!"
				fmt.Println("data directory: " + data_dir)
				msg.ReplyMarkup = utils.Keyboard_page1
			case "login":
				tgpass := update.Message.CommandArguments()
				fmt.Printf("password entered: = %+v\n", tgpass)
				authorize_admin = password.Pass_checker(&msg.Text, tgpass)
			case "senddoc":
				if authorize_admin && !IsDocument {
					msg.Text = "You can Now send the doc"
					fmt.Println("Senddoc command entered")

					IsDocument = true

				} else {
					msg.Text = "You should first try to /login ."
				}
			case "logout":
				authorize_admin = false
				msg.Text = "you logged out\n switch to normal user"
			case "help":
				msg.Text = utils.HelpMessage
			}
		}
		bot.Send(msg)
	}
}
