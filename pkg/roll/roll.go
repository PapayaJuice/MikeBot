package roll

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
	log "github.com/sirupsen/logrus"
)

// Inbound ...
func Inbound(session *discordgo.Session, message *discordgo.MessageCreate) {
	// Ignore messages by bot
	if message.Author.ID == session.State.User.ID {
		return
	}
	if message.Content == "" || message.Content[0] != '!' {
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
		response = CoinFlip(seed, message)
	case "!roll":
		response, err = Dice(body, message)
	default:
		log.Warnf("unknown command %s", command)
		return
	}
	if err != nil {
		log.WithError(err).Errorf("error performing %s", command)
		return
	}

	if response != "" {
		_, err = session.ChannelMessageSend(message.ChannelID, response)
		if err != nil {
			log.WithError(err).Errorf("Error sending message")
		}
	}

	return
}

// CoinFlip tosses a coin in the air and returns the result.
func CoinFlip(seed int64, message *discordgo.MessageCreate) string {
	states := []string{
		"heads",
		"tails",
	}
	rand.Seed(seed)

	return fmt.Sprintf("%s flips a coin... It's %s!", message.Author.Username, states[rand.Intn(len(states))])
}

// Dice rolls a dice with modifiers and returns the result.
func Dice(body string, message *discordgo.MessageCreate) (string, error) {
	return "", nil
}
