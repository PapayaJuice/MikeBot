package roll

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/bwmarrin/discordgo"
)

var (
	message = discordgo.MessageCreate{
		Message: &discordgo.Message{
			Author: &discordgo.User{
				Username: "KimChiPls",
			},
		},
	}
)

func TestCoinFlip(t *testing.T) {
	resp := CoinFlip(int64(0))
	assert.Equal(t, "Flips a coin... It's heads!", resp)
}
