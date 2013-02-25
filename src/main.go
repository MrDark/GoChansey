package main

import (
	"bot"
)

func main() {
	ircBot, err := bot.NewIrcBot("config.json")
	if err == nil {
		ircBot.Start()
	} else {
		println(err.Error())
	}
}
