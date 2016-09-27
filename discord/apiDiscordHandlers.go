package discord

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/joey-walker/johanbot/common"
)

// Send Message handler
func createMessage(discordSession *discordgo.Session, incomingMessage *discordgo.MessageCreate) {

	common.LogInfoAPI(common.DISCORD, "apiDiscordHandlers.go", "Adding message handler")

	// If we are mentioned then we care to respond.
	if len(incomingMessage.Mentions) > 0 && incomingMessage.Mentions[0].ID == BotID {

		if incomingMessage.Content == BotMention+" speak" {
			_, err := discordSession.ChannelMessageSend(incomingMessage.ChannelID,
				"Which one of us was unwanted?")
			fmt.Println("Sent message johan")
			if err != nil {
				fmt.Println("failed to speak: " + err.Error())
			}
			return
		}
	}
}

// Send message to discord we are shutting down.  More of a debug message.
func discordShutdown(discordSession *discordgo.Session, userChannels []*discordgo.Guild) {
	for _, userChannel := range userChannels {
		discordSession.ChannelMessageSend(userChannel.ID, "Johan Bot Stopping")
	}
}
