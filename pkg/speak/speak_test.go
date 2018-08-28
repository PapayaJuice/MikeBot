package speak

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
	assert.Equal(t, "KimChiPls slaps himself around with a large trout.", resp)
}

func TestLove(t *testing.T) {
	resp := Love("himself", &message)
	assert.Equal(t, "KimChiPls holds himself closely and kisses their cheek.", resp)
}
