package main

/*
이 파일에서는 slash command 를 정의합니다. package main 에 둠으로써 바로 쓸수있게 했습니다.
*/

import (
	"discord/service"
	"fmt"

	"github.com/bwmarrin/discordgo"
)

var commands = []*discordgo.ApplicationCommand{

	//general commands
	{
		//모든 command 와 option 은 description 을 가져야한다고함.
		Name:        "ping",
		Description: "Responds with pong",
	},
	{
		Name:        "pong",
		Description: "Responds with ping",
	},
	{
		Name:        "help",
		Description: "명령어들을 출력합니다",
	},

	//동방에 사람이 얼마나 있는지 확인하는 명령어
	{
		Name:        "좌석상황",
		Description: "현재 동아리방의 좌석 상황을 보여줍니다. 버튼으로 상호작용 가능합니다.",
	},
	// 문 관련 명령어들

	{
		Name:        "문등록",
		Description: "문을 등록합니다",
	},
}

var (
	Chan_id string //문등록으로 설정된 채널의 ID
	Door_id string //문등록으로 보낸 메세지의 ID 를 전역변수로 선언
)

var commandHandlers = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){

	//대충 이런 형식으로 함수를 만들어야 해서 남겨둠. Interaction 을 받으면 무조건 response 해야함.(디스코드 권장사양)
	"ping": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "pong",
			},
		})
		if err != nil {
			fmt.Println("error response", err)
			return
		}
	},
	"pong": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "ping",
			},
		})
		if err != nil {
			fmt.Println("error response", err)
			return
		}
	},
	"help": service.HelpMessage,
	"좌석상황": service.CheckSeatState,
	"문등록":  service.CreateDoor,
}
