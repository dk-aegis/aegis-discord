package main

/*
이 파일에서는 slash command 를 정의합니다. package main 에 둠으로써 바로 쓸수있게 했습니다.
*/

/*
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


			{

		Name:        "ping",
		Description: "Responds with pong",
	},
	{
		Name:        "pong",
		Description: "Responds with ping",
	},
*/

import (
	"discord/service"

	"github.com/bwmarrin/discordgo"
)

var commands = []*discordgo.ApplicationCommand{

	//general commands
	//모든 command 와 option 은 description 을 가져야한다고함.

	{
		Name:        "help",
		Description: "당신을 도와줄 명령어 모음입니다",
	},

	//동방에 사람이 얼마나 있는지 확인하는 명령어
	{
		Name:        "좌석상황",
		Description: "현재 동아리방의 좌석 상황을 보여줍니다. 버튼으로 상호작용 가능합니다.",
	},
	//출석
	{
		Name:        "출석",
		Description: "출석을 진행합니다.",
	},

	{
		Name:        "슬롯머신",
		Description: "슬롯머신을 돌립니다 돈 10원이 필요합니다.",
	},
	//관리자만 쓸 수 있도록 해야할듯.
	{
		Name:        "문등록",
		Description: "(권한 필요) 문을 등록합니다",
	},
	{
		Name:        "사용자등록",  //명령어 이름에 스페이스 들어가면 안됨
		Description: "(권한 필요) 사용자 전부 db에 올리기",
	},
	{
		Name: "정보",
		Description: "유저에 대한 정보를 보여줍니다",
	},
}

var commandHandlers = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate) {
	"help":  service.HelpMessage,
	"좌석상황":  service.CheckSeatState,
	"문등록":   service.CreateDoor,
	"출석":    service.DoAttendance,
	"사용자등록": service.ForkallGuild, 
	"슬롯머신":  service.Slotmachine,
}
