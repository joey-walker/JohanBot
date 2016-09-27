package main

import (
	"os"
	"testing"
)

const testToken string = "TEST_TOKEN"

func TestNewJoBot(t *testing.T) {
	testBot := NewJoBot()

	t.Log("Discord Token: " + testBot.DiscordToken())
	t.Log("Slack Token: " + testBot.SlackToken())

	if testBot.DiscordToken() == "" {
		t.Error("Discord Token came back empty, got ", testBot.DiscordToken())
	}

	if testBot.SlackToken() == "" {
		t.Error("Slack Token came back empty, got ", testBot.DiscordToken())
	}

	testBot.SetDiscordToken(testToken)

	if testBot.DiscordToken() != testToken {
		t.Error("Expected: "+testToken+" got: ", testBot.DiscordToken())
	}

	testBot.SetSlackToken(os.Getenv("FAKE_ENVIRONMENT_VARIABLE"))

	if testBot.SlackToken() != "" {
		t.Error("Slack Token came back not empty, got ", testBot.SlackToken())
	}

}
