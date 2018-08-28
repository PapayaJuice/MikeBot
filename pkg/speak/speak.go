package speak

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

// Slap naively slaps a member.
func Slap(body string, message *discordgo.MessageCreate) string {
	response := fmt.Sprintf("%s slaps %s around with a large trout.", message.Author.Username, body)
	return response
}

// Love holds a member closely.
func Love(body string, message *discordgo.MessageCreate) string {
	response := fmt.Sprintf("%s holds %s closely and kisses their cheek.", message.Author.Username, body)
	return response
}
