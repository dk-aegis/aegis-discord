package service

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

var (
	closedoor string = "https://cdn.discordapp.com/attachments/1276518219414503547/1276518281263710249/20240823_212621.jpg?ex=66d8fb0a&is=66d7a98a&hm=01b3d9c5410efe748f16c29e7b546dda7772d580442b3768e07ccf208f3e408c&"
	opendoor  string = "https://cdn.discordapp.com/attachments/1276518219414503547/1276518291816448111/20240823_212612.jpg?ex=66d8fb0d&is=66d7a98d&hm=412b3217584c9a50a303ee6971ec43f7b9f8b8bd0e9779d80ec146a153dde08d&"

	doorState = &discordgo.InteractionResponseData{
		Embeds: []*discordgo.MessageEmbed{
			{
				Title: "State of the door",
				Image: &discordgo.MessageEmbedImage{
					URL: opendoor,
				},
				Color: 0x00FF00,
			},
		},
		Components: []discordgo.MessageComponent{
			discordgo.ActionsRow{
				Components: []discordgo.MessageComponent{
					discordgo.Button{
						Label:    "열기",
						Style:    discordgo.SuccessButton,
						CustomID: "opendoor_btn",
						Emoji: &discordgo.ComponentEmoji{
							Name: "⚜️", //
						},
					},
					discordgo.Button{
						Label:    "닫기",
						Style:    discordgo.DangerButton,
						CustomID: "closedoor_btn",
						Emoji: &discordgo.ComponentEmoji{
							Name: "🚪",
						},
					},
				},
			},
		},
	}
)

func OpentheDoor(s *discordgo.Session, i *discordgo.InteractionCreate) {

	doorState.Embeds[0].Image.URL = opendoor
	doorState.Embeds[0].Color = 0x00FF00

	err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseUpdateMessage,
		Data: doorState,
	})

	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
}

func ClosetheDoor(s *discordgo.Session, i *discordgo.InteractionCreate) {

	doorState.Embeds[0].Image.URL = closedoor
	doorState.Embeds[0].Color = 0xa5ea89

	err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseUpdateMessage,
		Data: doorState,
	})

	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
}

func CreateDoor(s *discordgo.Session, i *discordgo.InteractionCreate) {

	err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: doorState,
	})

	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
}
