package discord

import (
	"github.com/bwmarrin/discordgo"
	"github.com/joey-walker/johanbot/common"
)

var (
	BotID      string
	BotMention string
)

func StartDiscord(discordToken string) {

	common.LogInfoAPI(common.DISCORD, "apiDiscord.go", "Opening discord Session")
	//Open connection with token obtained from system environment.
	discord, err := discordgo.New(discordToken)

	//All error messages need to be transferred to a proper logging mechanism
	if err != nil {
		common.LogErrorAPIFatal(common.DISCORD, "apiDiscord.go", err)
		return
	}

	//Acquire all channels we have access to
	userChannels, err := discord.UserGuilds()

	if err != nil {
		common.LogErrorAPI(common.DISCORD, "apiDiscord.go", err)
	}

	common.LogDebugAPI(common.DISCORD, "apiDiscord.go", "User Channels available", userChannels)

	//Debug message, tell all channels we are starting.
	/*for _, userChannel := range userChannels {
		discord.ChannelMessageSend(userChannel.ID, "Johan Bot Starting")
	}*/

	//Get own info
	botInfo, err := discord.Application("@me")
	if err != nil {
		common.LogErrorAPI(common.DISCORD, "apiDiscord.go", err)
	}

	common.LogDebugAPI(common.DISCORD, "apiDiscord.go", "Bot info", botInfo)

	// move these to bot struct
	//store bot id
	BotID = botInfo.ID

	// eg. <@4332425245224>
	BotMention = "<@" + BotID + ">"

	//Add function for actions on listening to channel
	discord.AddHandler(createMessage)

	//open the channel
	common.LogInfoAPI(common.DISCORD, "apiDiscord.go", "Connecting to discord channels")
	err = discord.Open()
	if err != nil {
		common.LogErrorAPI(common.DISCORD, "apiDiscord.go", err)
	}

	// close when we are done.
	defer discord.Close()

	// tell everyone we are turning off.
	//defer discordShutdown(discord, userChannels)
}
