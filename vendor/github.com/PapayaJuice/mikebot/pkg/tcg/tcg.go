package tcg

import (
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
	log "github.com/sirupsen/logrus"
)

const (
	tcgRespTemplate = "```\nName: %s\nRarity: %s\nOracle Text: %s\n```\n"
)

// Inbound ...
func Inbound(session *discordgo.Session, message *discordgo.MessageCreate) {
	// Ignore messages by bot
	if message.Author.ID == session.State.User.ID {
		return
	}

	var cards []string
	parts := strings.Split(message.Content, "[[")
	if len(parts) < 2 {
		return
	}
	for _, part := range parts {
		p := strings.Split(part, "]]")
		if len(p) != 2 {
			continue
		}
		cards = append(cards, p[0])
	}

	for _, card := range cards {
		resp, img, err := searchTCG(card)
		if err != nil {
			log.WithError(err).Error("error searching tcgplayer API")
			session.ChannelMessageSend(message.ChannelID, err.Error())
			continue
		}

		msg := discordgo.MessageSend{
			Content: resp,
			Embed: &discordgo.MessageEmbed{
				Image: &discordgo.MessageEmbedImage{
					URL: img,
				},
			},
		}

		_, err = session.ChannelMessageSendComplex(message.ChannelID, &msg)
		if err != nil {
			log.WithError(err).Error("error sending message")
		}
	}
}

func searchTCG(text string) (string, string, error) {
	cardID, err := searchCard(text)
	if err != nil {
		return "", "", err
	}
	cardInfo, err := requestCard(cardID)
	if err != nil {
		return "", "", err
	}

	rarity := "Not Found"
	oracle := "Not Found"
	for _, ex := range cardInfo.ExtendedData {
		switch ex.Name {
		case "Rarity":
			rarity = ex.Value
		case "OracleText":
			oracle = ex.Value
		}
	}
	return fmt.Sprintf(tcgRespTemplate, cardInfo.Name, rarity, oracle), cardInfo.ImageURL, nil
}
