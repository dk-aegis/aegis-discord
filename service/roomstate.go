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
	//임베드 메세지의 필드로서 자리는 한 9개 정도 해놓음.
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

func countEmpty(table []*discordgo.MessageEmbedField) int {
	var count int = 0
	for _ , now := range table {
		if now.Name == ":x:" && now.Value == "공석" {
			count++
		}
	}
	return count
}

func existOnTable(table []*discordgo.MessageEmbedField, userName string) bool {
	var exist bool = false
	for _ , now := range table {
		if now.Name != ":x:" && now.Value == "userName" {
			exist = true
		}
	}
	return exist
}


/*
좌석을 하나 차지하고, 하나 빼는 작업을 수행해야 하는데 어떻게 구현할까.
슬라이스는 처음 x공석으로 초기화 되어있음. 일단 공석의 수를 세는 코드를 만듭시다.
*/
func TakeaSeat(s *discordgo.Session, i *discordgo.InteractionCreate) {

	m := i.Member //Member 접근할 때 nil 검사하는게 좋다고 해서 그냥 해줌.
	if  m != nil {
		fmt.Println("user info load error", m)
	}

	userName := m.User.Username  //username 을 함부로 못바꾸게 해야할지도 

	if countEmpty(RoomStateEmbed.Fields) == 0 {  //빈자리 없다!
		s.ChannelMessageSend(i.ChannelID,"차지할 수 있는 좌석이 없어요!") // Out!
		return
	}
	
	if existOnTable(RoomStateEmbed.Fields,userName) {
		s.ChannelMessageSend(i.ChannelID,"욕심이 많으시네요") // Out!
		return
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
