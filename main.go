package main

import (
	"discord/global"
	"discord/service"
	"encoding/json"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

type TokenConfig struct {
	Token string `json:"token"`
}

func getToken() (TokenConfig, error) {
	var tc TokenConfig
	file, err := os.Open("./config/token.json")

	if err != nil {
		return TokenConfig{}, err
	}

	defer file.Close()

	jsonParser := json.NewDecoder(file)
	jsonParser.Decode(&tc)

	return tc, nil
}

func main() {

	// Create a new Discord session using the provided bot token.
	tc, err := getToken()
	if err != nil {
		fmt.Println(err)
		return
	}
	dg, err := discordgo.New("Bot " + tc.Token)
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}

	service.InitEvent()
	err = service.InitDatabase()
	if err != nil {
		fmt.Println(err)
		return
	}

	err = global.InitDiscordConfig()
	if err != nil {
		fmt.Println(err)
		return
	}

	// Register the messageCreate func as a callback for MessageCreate events.
	dg.AddHandler(messageCreate)
	dg.AddHandler(service.MemberJoin)

	// In this example, we only care about receiving message events.
	dg.Identify.Intents = discordgo.IntentsGuildMessages | discordgo.IntentsGuildMembers | discordgo.IntentsGuilds | discordgo.IntentsDirectMessages

	// Open a websocket connection to Discord and begin listening.
	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	// Cleanly close down the Discord session.
	dg.Close()
}
