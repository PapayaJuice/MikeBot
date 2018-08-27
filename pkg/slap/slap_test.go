package slap

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

func TestSlap(t *testing.T) {
	resp := Slap("himself", &message)
	assert.Equal(t, "KimChiPls slaps himself with a big fish.", resp)
}
