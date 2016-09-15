package main

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"os"
)

// Bot information should placed in own bot Struct
var (
	BotID      string
	BotMention string
	exit       bool = false
)

func main() {

	//Open connection with token obtained from system environment.
	//We need to move this to its own channel so we can use both slack and discord.
	discord, err := discordgo.New(os.Getenv("JOHAN_TOKEN"))

	//All error messages need to be transferred to a proper logging mechanism
	if err != nil {
		fmt.Println("Failed to connect, ", err)
		return
	}

	//Acquire all channels we have access to
	userChannels, err := discord.UserGuilds()

	if err != nil {
		fmt.Println("Fail to acquire the channels we have access to.")
	}

	//Debug message, tell all channels we are starting.
	for _, userChannel := range userChannels {
		discord.ChannelMessageSend(userChannel.ID, "Johan Bot Starting")
	}



	//Get own info
	botInfo, err := discord.Application("@me")
	if err != nil {
		fmt.Println("Failed to get johanbot info", err)
		return
	}

	//store bot id
	BotID = botInfo.ID

	// eg. <@4332425245224>
	BotMention = "<@" + BotID + ">"

	//Add function for actions on listening to channel
	discord.AddHandler(createMessage)


	//open the channel
	err = discord.Open()
	if err != nil {
		fmt.Println("Failed to open connection,", err)
		return
	}

	// close when we are done.
	defer discord.Close()

	// tell everyone we are turning off.
	defer discordShutdown(discord, userChannels)

	// Until one of handlers tells use to shut down...
	for exit != true {
	}

	return

}

// Send Message handler
func createMessage(discordSession *discordgo.Session, incomingMessage *discordgo.MessageCreate) {

	/* ignore messages by self
	if m.Author.ID == BotID {
		return
	}*/

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

		// If we tell johan to go off, we should stop listening.
		if incomingMessage.Content == BotMention+" off" {
			exit = true
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
