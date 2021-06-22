package speak

import (
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
	log "github.com/sirupsen/logrus"
)

// Inbound ...
func Inbound(session *discordgo.Session, message *discordgo.MessageCreate) {
	// Ignore messages by bot
	if message.Author.ID == session.State.User.ID {
		return
	}
	if message.Content[0] != '!' {
		return
	}

	parts := strings.Split(message.Content, " ")
	command := parts[0]
	body := strings.Join(parts[1:], " ")

	var response string
	var err error
	switch command {
	case "!love":
		response = fmt.Sprintf("%s holds %s closely and kisses their cheek.", message.Author.Username, body)
	case "!slap":
		response = fmt.Sprintf("%s slaps %s around with a large trout.", message.Author.Username, body)
	default:
		log.Warnf("unknown command %s", command)
		return
	}

	if response != "" {
		_, err = session.ChannelMessageSend(message.ChannelID, response)
		if err != nil {
			log.WithError(err).Errorf("error sending message")
		}
	}

	return
}
