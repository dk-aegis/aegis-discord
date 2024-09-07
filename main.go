package main

import (
	"discord/global"
	"discord/service"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "OK")
}

func main() {

	// Create a new Discord session using the provided bot token.
	err := global.InitDiscordConfig() //config file on
	if err != nil {
		fmt.Println(err)
		return
	}

	dg, err := discordgo.New("Bot " + global.Discord.Bot.Token) //session 을 생성합니다. 이 session 으로 discordbot 의 동작이나 상태를 관리할 수 있습니다.
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}

	err = service.InitDatabase() //DB on
	if err != nil {
		fmt.Println(err)
		return
	}

	http.HandleFunc("/health", HealthCheck)

	// Register the messageCreate func as a callback for MessageCreate events.
	//AddHandler 는 인자가 2개인 함수를 인자로 받음. 첫번째는 세션, 두번째는 이벤트...

	dg.AddHandler(InteractionHandler)
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
		_, err := dg.ApplicationCommandCreate(dg.State.User.ID, global.Discord.GuildID, cmd) //SlashCommands 를 추가합니다. 봇의 ID / 길드ID / 명령어
		if err != nil {
			log.Fatalf("Cannot create Command.. Error at : %s", cmd.Name)
		} else {
			fmt.Printf("adding command success %s \n", cmd.Name)
		}
	}

	dg.UpdateCustomStatus("https://dk-aegis.org  |  /help")

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	// Cleanly close down the Discord session.
	dg.Close()

	service.DBclose()
}
