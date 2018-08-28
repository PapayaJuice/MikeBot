package roll

import (
	"fmt"
	"math/rand"

	"github.com/bwmarrin/discordgo"
)

// CoinFlip tosses a coin in the air and returns the result.
func CoinFlip(seed int64) string {
	states := []string{
		"heads",
		"tails",
	}
	rand.Seed(seed)

	return fmt.Sprintf("Flips a coin... It's %s!", states[rand.Intn(len(states))])
}

// Dice rolls a dice with modifiers and returns the result.
func Dice(body string, message *discordgo.MessageCreate) (string, error) {
	return "", nil
}
