package main

/*
이 파일에서는 slash command 를 정의합니다.
*/

import (
	"github.com/bwmarrin/discordgo"
)

var commands = []*discordgo.ApplicationCommand{
	{
		Name: "ping",
		//모든 command 와 option 은 description 을 가져야한다고함.
		Description: "Responds with pong",
	},
	{
		Name:        "pong",
		Description: "Responds with ping",
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
	"ping": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "pong",
			},
		})
	},
	"pong": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "ping",
			},
		})
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
