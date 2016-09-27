package slack

import (
	"fmt"

	"github.com/nlopes/slack"
)

func StartSlack(slackToken string) {

	api := slack.New(slackToken)

	slackChannels, err := api.GetChannels(true)

	if err != nil {
		fmt.Println("failure to connect to slack")
	}

	params := slack.PostMessageParameters{}

	for _, channel := range slackChannels {
		api.PostMessage(channel.ID, "johanbot sending message", params)
	}

}
