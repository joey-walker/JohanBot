package main

import (
	"github.com/joey-walker/johanbot/discord"
	"github.com/joey-walker/johanbot/common"
	"github.com/joey-walker/johanbot/slack"
)

var ItsTimeToStop bool = true

// Bot information should placed in own bot Struct

func main() {

	// New Bot
	johanBot := NewJoBot()

	common.LogInfo("main.go", "Starting Bot activation on IRC channels")

	//TODO: probably pass bot struct through.
	//Start listening on discord
	discord.StartDiscord(johanBot.DiscordToken())
	slack.StartSlack(johanBot.SlackToken())


	// go forever till termination
	for ItsTimeToStop {

	}
}
