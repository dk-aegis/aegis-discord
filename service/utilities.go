package service
//잡다하게 많이 쓸거같은 기능들 모아놓음
import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func SendInteractionMessage(s *discordgo.Session, i *discordgo.InteractionCreate, msg string) {
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

func CheckRole(roleSlice []string, role string) bool {
	for _, r := range roleSlice {
		if r == role {
			return true
		}
	}
	return false
}
