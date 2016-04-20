package memes

import (
	"github.com/parsley42/gobot-chatops/bot"
)

var (
	gobot   bot.Robot
	botName string
)

func memegen(bot bot.Robot, channel, user, command string, args ...string) error {
	switch command {
	case "start":
		gobot = bot
		botName = user
	case "simply":
		bot.SendChannelMessage(channel, "Yeah, you're right about that!")
	}
	return nil
}

func init() {
	bot.RegisterPlugin("memes", memegen)
}