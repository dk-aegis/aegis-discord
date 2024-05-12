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

// 엠베드 메세지로 다 출력할거임. 엠베드 메세지 하는법 좀 공부 ㄱㄱ
var eventList []Event

func ShowEvent(s *discordgo.Session, m *discordgo.MessageCreate) {

	if len(eventList) == 0 { //이벤트가 3개 이상이면 뭘 볼지 선택하세요
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

func CreateEvent(s *discordgo.Session, m *discordgo.MessageCreate) {

}

func RemoveEvent(s *discordgo.Session, m *discordgo.MessageCreate) {

}

func ShowEventLists(m *discordgo.MessageCreate) {

}
