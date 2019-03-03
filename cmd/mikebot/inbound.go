package main

import (
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
	log "github.com/sirupsen/logrus"

	"github.com/PapayaJuice/mikebot/pkg/roll"
	"github.com/PapayaJuice/mikebot/pkg/speak"
	"github.com/PapayaJuice/mikebot/pkg/tcg"
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

	// Split for tcg player
	// TODO: Clean this entire function it's so gross
	var tcgCard string
	tcgP := strings.Split(message.Content, "[[")
	if len(tcgP) == 2 {
		tcgP = strings.Split(tcgP[1], "]]")
		tcgCard = tcgP[0]
	}

	var response string
	var err error
	var tcgResp string
	var tcgImg string
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
		if strings.HasPrefix(command, "![[") && tcgCard != "" {
			tcgResp, tcgImg, err = tcg.SearchTCG(tcgCard)
		} else {
			log.Warnf("Unknown command %s\n", command)
			return
		}
	}
	if err != nil {
		log.Errorf("Error performing %s: %v\n", command, err)
		return
	}

	if tcgResp != "" {
		msg := discordgo.MessageSend{
			Content: tcgResp,
			Embed: &discordgo.MessageEmbed{
				Image: &discordgo.MessageEmbedImage{
					URL: tcgImg,
				},
			},
		}
		_, err = session.ChannelMessageSendComplex(message.ChannelID, &msg)
		if err != nil {
			log.Errorf("Error sending message: %v\n", err)
		}
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
