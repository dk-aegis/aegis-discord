package service

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

type Event struct {
	FromDate  string
	UntilDate string
	Title     string
	Content   string
}

var eventList []Event

func ShowEvent(s *discordgo.Session, m *discordgo.MessageCreate) {

	if len(eventList) == 0 {
		sendEmbedMessage(s, m.ChannelID, "ㅠㅠ", "현재 진행중인 이벤트가 없습니다.", 0x00ff00)
		return
	}

	for _, event := range eventList {
		sendEmbedMessage(s, m.ChannelID, event.Title, event.String(), 0x00ff00)
	}
}

func (e Event) String() string {
	s := fmt.Sprintf("**%s**\n\n시작 날짜: %s\n\n종료 날짜: %s\n", e.Content, e.FromDate, e.UntilDate)
	return s
}

func sendEmbedMessage(s *discordgo.Session, chid, title, content string, color int) {
	embed := &discordgo.MessageEmbed{
		Title:       title,
		Description: content,
		Color:       color, // Green color
	}

	// Send the embedded message to the channel where the command was received.
	_, err := s.ChannelMessageSendEmbed(chid, embed)
	if err != nil {
		fmt.Println("Error sending embedded message:", err)
		return
	}
}
