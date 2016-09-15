package main

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"os"
)

// derp
var (
	BotID string
	BotMention string
	exit bool = false
)

func main() {

	discord, err := discordgo.New(os.Getenv("JOHAN_TOKEN"))

	userChannels, err := discord.UserGuilds()

	for _, userChannel := range userChannels {
		discord.ChannelMessageSend(userChannel.ID, "Johan Bot Starting")
	}

	if err != nil {
		fmt.Println("Failed to connect, ", err)
		return
	}

	botInfo, err := discord.Application("@me")
	if err != nil {
		fmt.Println("Failed to get johanbot info", err)
		return
	}

	BotID = botInfo.ID
	BotMention = "<@" + BotID + ">"
	discord.AddHandler(createMessage)

	err = discord.Open()
	if err != nil {
		fmt.Println("Failed to open connection,", err)
		return
	}

	fmt.Println("JohanBot is now running.  Press CTRL-C to exit.")

	defer discord.Close()

	defer discordShutdown(discord, userChannels)
	for exit != true {}

	return

}

func createMessage(s *discordgo.Session, m *discordgo.MessageCreate) {

	// ignore messages by self
	if m.Author.ID == BotID {
		return
	}

	if len(m.Mentions) > 0 && m.Mentions[0].ID == BotID {

		if m.Content == BotMention + " speak" {
			_, err := s.ChannelMessageSend(m.ChannelID, "Which one of us was unwanted?")
			fmt.Println("Sent message johan")
			if err != nil {
				fmt.Println("failed to speak: " + err.Error())
			}
			return
		}

		if m.Content == BotMention + " off" {
			exit = true
			return
		}
	}
}

func discordShutdown(s *discordgo.Session, userChannels []*discordgo.Guild) {
	for _, userChannel := range userChannels {
		s.ChannelMessageSend(userChannel.ID, "Johan Bot Stopping")
	}
}
