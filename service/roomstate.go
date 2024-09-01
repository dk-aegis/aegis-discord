package service

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

/*
	{
		Name:   ":white_check_mark:",
		Value:  "안모씨",
		Inline: false,
	},
	{
		Name:   ":x:",
		Value:  "공석",
		Inline: false,
	},

*/

var RoomStateEmbed *discordgo.MessageEmbed = &discordgo.MessageEmbed{
	Title:       "현재 좌석 상황",
	Description: "",
	Color:       0x00ff00,
	Fields: []*discordgo.MessageEmbedField{
		{
			Name:   ":x:",
			Value:  "공석",
			Inline: true,
		},
		{
			Name:   ":x:",
			Value:  "공석",
			Inline: true,
		},
		{
			Name:   ":x:",
			Value:  "공석",
			Inline: true,
		},
		{
			Name:   ":x:",
			Value:  "공석",
			Inline: true,
		},
		{
			Name:   ":x:",
			Value:  "공석",
			Inline: true,
		},
		{
			Name:   ":x:",
			Value:  "공석",
			Inline: true,
		},
		{
			Name:   ":x:",
			Value:  "공석",
			Inline: true,
		},
		{
			Name:   ":x:",
			Value:  "공석",
			Inline: true,
		},
		{
			Name:   ":x:",
			Value:  "공석",
			Inline: true,
		},
	},
}

func TakeaSeat(s *discordgo.Session, i *discordgo.InteractionCreate) {
	user := i.User
	if  user != nil {
		
	}
	
}

func Standup(s *discordgo.Session, i *discordgo.InteractionCreate) {

}

func CheckSeatState(s *discordgo.Session, i *discordgo.InteractionCreate) {

	err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Embeds: []*discordgo.MessageEmbed{RoomStateEmbed},
		},
	})

	if err != nil {
		fmt.Println("error response", err)
		return
	}
}
