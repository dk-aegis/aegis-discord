package service

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	_ "github.com/go-sql-driver/mysql"
)

func HelpMessage(s *discordgo.Session, i *discordgo.InteractionCreate) {
	msg := `
		/좌석상황 : 현재 동아리방의 좌석 상황을 보여줍니다. 버튼으로 상호작용 가능합니다.
		
		/출석 : 출석을 진행합니다.

		/슬롯머신 : 슬롯머신을 돌립니다. 돈 10원이 필요합니다.

		/정보 : 유저에 대한 정보를 보여줍니다.
	`
	embed := &discordgo.MessageEmbed{
		Title:       "commands",
		Description: msg,
		Color:       0x00ff00,
	}
	

	err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Embeds: []*discordgo.MessageEmbed{embed},
		},
	})

	if err != nil {
		fmt.Println("error response", err)
		return
	}
}
