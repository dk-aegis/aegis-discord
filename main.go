package main

import (
	"discord/global"
	"discord/service"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

type TokenConfig struct {
	Token   string `json:"token"`
	GuildID string `json:"guild_id"`
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

	dg, err := discordgo.New("Bot " + tc.Token) //session 을 생성합니다. 이 session 으로 discordbot 의 동작이나 상태를 관리할 수 있습니다.
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}

	err = service.InitDatabase() //DB on
	if err != nil {
		fmt.Println(err)
		return
	}

	err = global.InitDiscordConfig() //config file on
	if err != nil {
		fmt.Println(err)
		return
	}

	// Register the messageCreate func as a callback for MessageCreate events.
	//AddHandler 는 인자가 2개인 함수를 인자로 받음. 첫번째는 세션, 두번째는 이벤트...

	dg.AddHandler(InteractionHandler)
	dg.AddHandler(messageCreate)
	dg.AddHandler(service.MemberJoin)

	// 디스코드봇의 권한을 설정하는것같은..
	dg.Identify.Intents = discordgo.IntentsGuildMessages | discordgo.IntentsGuildMembers | discordgo.IntentsGuilds | discordgo.IntentsDirectMessages | discordgo.IntentsGuildPresences

	// Open a websocket connection to Discord and begin listening.
	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	fmt.Println("Adding Commands!")
	for _, cmd := range commands { //cmd가 명령어 같긴 한데.. 뭘 무시하고 cmd 를 받는건지 잘 모르겠네
		_, err := dg.ApplicationCommandCreate(dg.State.User.ID, tc.GuildID, cmd) //SlashCommands 를 추가합니다. 봇의 ID / 길드ID / 명령어
		if err != nil {
			log.Fatalf("Cannot create Command.. Error at : %s", cmd.Name)
		} else {
			fmt.Printf("adding command success %s \n", cmd.Name)
		}
	}

	dg.UpdateWatchStatus(1, "/도움말") //디코봇 상태 메세지 설정

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	// Cleanly close down the Discord session.
	dg.Close()

	service.DBclose()
}
