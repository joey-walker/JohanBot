package main

import "os"

// Our bot type, holds only onto to new fields right now:
// SlackToken, to access slack channel
// DiscordToken, to access discord channels
type Bot struct {
	slackToken   string
	discordToken string
}

// Need pointer in order to modify struct field.
// SlackToken Setter
func (johanBot *Bot) SetSlackToken(slackToken string) {
	johanBot.slackToken = slackToken
}

// No pointer means it just receives a copy.
// SlackToken getter
func (johanBot Bot) SlackToken() string {
	return johanBot.slackToken
}

// DiscordToken setter
func (johanBot *Bot) SetDiscordToken(discordToken string) {
	johanBot.discordToken = discordToken
}

// DiscordToken getter
func (johanBot Bot) DiscordToken() string {
	return johanBot.discordToken
}

// basic constructor
func NewJoBot() *Bot {

	johanBot := Bot{
		os.Getenv("JOHAN_SLACK"),
		os.Getenv("JOHAN_DISCORD"),
	}

	return &johanBot
}
