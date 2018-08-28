package main

import (
	"strings"
	"time"

	"github.com/PapayaJuice/mikebot/pkg/roll"
	"github.com/PapayaJuice/mikebot/pkg/speak"
	"github.com/bwmarrin/discordgo"
	log "github.com/sirupsen/logrus"
)

func routeInbound(session *discordgo.Session, message *discordgo.MessageCreate) {
	// Ignore messages by bot
	if message.Author.ID == session.State.User.ID {
		return
	}
	if message.Content[0] != '!' {
		return
	}

	// Seed for commands which need one
	seed := time.Now().UnixNano()

	parts := strings.Split(message.Content, " ")
	command := parts[0]
	body := strings.Join(parts[1:], " ")

	var response string
	var err error
	switch command {
	case "!coinflip":
		response = roll.CoinFlip(seed, message)
	case "!love":
		response = speak.Love(body, message)
	case "!roll":
		response, err = roll.Dice(body, message)
	case "!slap":
		response = speak.Slap(body, message)
	default:
		log.Warnf("Unknown command %s\n", command)
		return
	}
	if err != nil {
		log.Errorf("Error performing %s: %v\n", command, err)
		return
	}

	err = session.ChannelMessageDelete(message.ChannelID, message.ID)
	if err != nil {
		log.Errorf("Error deleting message: %v\n", err)
	}

	if response != "" {
		_, err = session.ChannelMessageSend(message.ChannelID, response)
		if err != nil {
			log.Errorf("Error sending message: %v\n", err)
		}
	}

	return
}
