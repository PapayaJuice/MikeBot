package slap

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

// Slap naively slaps a member.
func Slap(body string, message *discordgo.MessageCreate) string {
	response := fmt.Sprintf("%s slaps %s with a big fish.", message.Author.Username, body)
	return response
}
