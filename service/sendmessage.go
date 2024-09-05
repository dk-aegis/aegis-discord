package service

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func SendInteractionMessage (s *discordgo.Session, i *discordgo.InteractionCreate, msg string){
	err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: msg,
		},
	})

	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
}
