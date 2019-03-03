package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	log "github.com/sirupsen/logrus"

	"github.com/PapayaJuice/mikebot/pkg/tcg"
)

var (
	debug = flag.Bool("debug", false, "Turn on debug logging")
	token = flag.String("token", "", "Token for discord bot")
)

func init() {
	if *debug {
		log.SetLevel(log.DebugLevel)
	} else {
		log.SetLevel(log.ErrorLevel)
	}
}

func main() {
	flag.Parse()

	// Start refresh service for tcgplayer API
	go tcg.TokenRefresh()

	bot, err := discordgo.New(fmt.Sprintf("Bot %s", *token))
	if err != nil {
		log.Fatalf("Error initiating bot: %v\n", err)
	}

	bot.AddHandler(routeInbound)

	err = bot.Open()
	if err != nil {
		log.Fatalf("Error opening websocket to Discord: %v\n", err)
	}

	log.Info("Listening for messages...")

	killChan := make(chan os.Signal, 1)
	signal.Notify(killChan, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-killChan

	log.Info("Goodbye")
	bot.Close()
}
