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
		Name:        "착석",
		Description: "자리 하나를 차지합니다",
	},
	{
		Name:        "기립",
		Description: "자리를 하나 내어줍니다",
	},
	{
		Name:        "좌석상황",
		Description: "현재 동아리방의 좌석 상황을 보여줍니다",
	},
	{
		Name:        "역할부여",
		Description: "역할부여하기도식이",
	},

	// 문 관련 명령어들

	{
		Name:        "문열기",
		Description: "동아리방 문을 열림 상태로 바꿉니다. 문등록이 된 채널에서만 사용 가능합니다.",
	},
	{
		Name:        "문닫기",
		Description: "동아리방 문을 닫힘 상태로 바꿉니다. 문등록이 된 채널에서만 사용 가능합니다.",
	},
	{
		Name:        "문등록",
		Description: "한 채널에 문의 상태를 보여주는 임베드메세지를 보냅니다. 문 관련된 명령어는 문등록 명령어를 실행한 채널에서만 가능합니다.",
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
	"착석":   service.TakeaSeat,
	"기립":   service.Standup,
	"좌석상황": service.CheckSeatState,
	"역할부여": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		err := s.GuildMemberNickname(i.GuildID, i.Member.User.ID, "doshik")
		if err != nil {
			return
		}
	},

	"문열기": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "Loading..",
			},
		})
		if Door_id != "" { // 메세지가 없으면
			if Chan_id != i.ChannelID { //등록한 채널과 다르면
				s.ChannelMessageSend(i.ChannelID, "이 채널에서 명령어를 사용할 수 없습니다.")
				s.InteractionResponseDelete(i.Interaction)
				return
			}
			embed := &discordgo.MessageEmbed{
				Title: "State of the door",
				Image: &discordgo.MessageEmbedImage{
					URL: "https://cdn.discordapp.com/attachments/1276518219414503547/1276518291816448111/20240823_212612.jpg?ex=66c9d1cd&is=66c8804d&hm=d47a842b833b442e3aff3f94ec3865cf4256318accae3fd97e5f4815778faa33&",
				},
				Color: 0xa5ea89,
			}
			s.ChannelMessageEditEmbed(i.ChannelID, Door_id, embed)
		} else {
			s.ChannelMessageSend(i.ChannelID, "문등록을 먼저 해주세요")
		}

		s.InteractionResponseDelete(i.Interaction)

	},
	"문닫기": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "Loading..",
			},
		})
		if Door_id != "" {
			if Chan_id != i.ChannelID { //등록한 채널과 다르면
				s.ChannelMessageSend(i.ChannelID, "이 채널에서 명령어를 사용할 수 없습니다.")
				s.InteractionResponseDelete(i.Interaction)
				return
			}
			embed := &discordgo.MessageEmbed{
				Title: "State of the door",
				Image: &discordgo.MessageEmbedImage{
					URL: "https://cdn.discordapp.com/attachments/1276518219414503547/1276518281263710249/20240823_212621.jpg?ex=66c9d1ca&is=66c8804a&hm=cdaac532d7c9782f373507f17a232c8bb8da1bf4e2d67941aa71c826cb2332b8&",
				},
				Color: 0xff8e7f,
			}
			s.ChannelMessageEditEmbed(i.ChannelID, Door_id, embed)
		} else {
			s.ChannelMessageSend(i.ChannelID, "문등록을 먼저 해주세요")
		}
		s.InteractionResponseDelete(i.Interaction)
	},
	"문등록": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "Loading..",
			},
		})
		embed := &discordgo.MessageEmbed{
			Title: "State of the door",
			Image: &discordgo.MessageEmbedImage{
				URL: "https://cdn.discordapp.com/attachments/1276518219414503547/1276518281263710249/20240823_212621.jpg?ex=66c9d1ca&is=66c8804a&hm=cdaac532d7c9782f373507f17a232c8bb8da1bf4e2d67941aa71c826cb2332b8&",
			},
		}
		msg, err := s.ChannelMessageSendEmbed(i.ChannelID, embed)
		if err != nil {
			s.ChannelMessageSend(i.ChannelID, "임베드 메세지 보내기에 실패했습니다.")
			return
		}

		// 등록할 때 메세지의 ID 랑 채널의 ID 를 저장
		Door_id = msg.ID
		Chan_id = i.ChannelID

		s.InteractionResponseDelete(i.Interaction)
	},
}
