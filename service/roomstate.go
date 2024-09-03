package service

import (
	"fmt"
	"regexp"


	"github.com/bwmarrin/discordgo"
)

/*
	{
		Name:   ":white_check_mark:",
		Value:  "24누군가(상태)",          여기서 졸업생이면 숫자가 없을 수도 있음.
		Inline: true,
	},
	{
		Name:   ":x:",
		Value:  "공석",
		Inline: true,
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
	for _, now := range table {
		if now.Name == ":x:" && now.Value == "공석" {
			count++
		}
	}
	return count
}

func existOnTable(table []*discordgo.MessageEmbedField, userName string) bool {
	for _, now := range table {
		if now.Value == userName {
			return true
		}
	}
	return false
}

func sliceName(name string) string{
	re := regexp.MustCompile(`^\d*|\(.*?\)$`)
	return re.ReplaceAllString(name, "")
}

/*
좌석을 하나 차지하고, 하나 빼는 작업을 수행해야 하는데 어떻게 구현할까.
슬라이스는 처음 x공석으로 초기화 되어있음. 일단 공석의 수를 세는 코드를 만듭시다.
*/
func TakeaSeat(s *discordgo.Session, i *discordgo.InteractionCreate) {

	Nickname := i.Member.Nick

	if countEmpty(RoomStateEmbed.Fields) <= 0 { //빈자리 없다!
		err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "빈 좌석이 없어요",
			},
		})
		if err != nil {
			fmt.Println("error response", err)
		}
		return
	}

	if existOnTable(RoomStateEmbed.Fields, Nickname) {
		err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "이미 좌석을 차지하고 계십니다",
			},
		})
		if err != nil {
			fmt.Println("error response", err)
		}
		return
	}

	//상태 바꿈
	for index, now := range RoomStateEmbed.Fields {
		if now.Name == ":x:" {
			RoomStateEmbed.Fields[index].Name = ":white_check_mark:"
			RoomStateEmbed.Fields[index].Value = Nickname
			break
		}
	}

	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseUpdateMessage,
		Data: &discordgo.InteractionResponseData{
			Embeds: []*discordgo.MessageEmbed{RoomStateEmbed},
		},
	})

}

func Standup(s *discordgo.Session, i *discordgo.InteractionCreate) {

	Nickname := i.Member.Nick

	if existOnTable(RoomStateEmbed.Fields, Nickname) {

		fmt.Println("ㅎㅇ?")

		for index, now := range RoomStateEmbed.Fields {
			if now.Value == Nickname {
				RoomStateEmbed.Fields[index].Name = ":x:"
				RoomStateEmbed.Fields[index].Value = "공석"
				break
			}
		}

	} else {
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "좌석에 없습니다",
			},
		})
	}

	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseUpdateMessage,
		Data: &discordgo.InteractionResponseData{
			Embeds: []*discordgo.MessageEmbed{RoomStateEmbed},
		},
	})
}

func CheckSeatState(s *discordgo.Session, i *discordgo.InteractionCreate) {

	err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Embeds: []*discordgo.MessageEmbed{RoomStateEmbed},
			Components: []discordgo.MessageComponent{
				discordgo.ActionsRow{
					Components: []discordgo.MessageComponent{
						discordgo.Button{
							Label:    "착석",
							Style:    discordgo.SuccessButton,
							CustomID: "sitdown_btn",
							Emoji: discordgo.ComponentEmoji{
								Name: "🧘", // Unicode 이모지가 들어가야함 window + . 으로 하는 이모지만 들어갈 수 있음 :x: 이런식이면 에러남.
							},
						},
						discordgo.Button{
							Label:    "기립",
							Style:    discordgo.DangerButton,
							CustomID: "standup_btn",
							Emoji: discordgo.ComponentEmoji{
								Name: "🏃",
							},
						},
						discordgo.Button{
							Label:    "Exit",
							Style:    discordgo.DangerButton,
							CustomID: "exit_btn",
							Emoji: discordgo.ComponentEmoji{
								Name: "❌",
							},
						},
					},
				},
			},
		},
	})

	if err != nil {
		fmt.Println("error response", err)
		return
	}
}
