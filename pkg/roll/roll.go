package roll

import (
	"fmt"
	"math/rand"

	"github.com/bwmarrin/discordgo"
)

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
