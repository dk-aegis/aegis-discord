package service

import (
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
)

type Event struct {
	FromDate  string
	UntilDate string
	Title     string
	Content   string
}

// 엠베드 메세지로 다 출력할거임. 엠베드 메세지 하는법 좀 공부 ㄱㄱ

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

// input form !prefix Title Content 원본글URL fromdate untildate
func CreateEvent(s *discordgo.Session, m *discordgo.MessageCreate) error {

	check, err := CheckAdmin(Hashstring(m.Author.ID)) //권한체크  권한이 없다는것도 출력해야됨.
	if (err != nil) || (!check) {
		fmt.Println(err)
		return err
	}

	args := strings.Split(m.Content[18:], "|") //'!이벤트 등록' 이 8문자를 죽이고 시작 , 구분자는 | (shift+\)

	if len(args) < 5 {
		msg := "인수가 부족해요"
		_, err = s.ChannelMessageSend(m.ChannelID, msg) //인수가 부족한가??
		if err != nil {
			return err
		}
		return nil
	}

	query := `INSERT INTO notice (title, content, noticeurl, fromdate, untildate)
	VALUES (?,?,?,?,?)`

	sql, err := db.Exec(query, strings.TrimSpace(args[0]), strings.TrimSpace(args[1]), strings.TrimSpace(args[2]), strings.TrimSpace(args[3]), strings.TrimSpace(args[4]))

	if err != nil {
		fmt.Println(err)
		return err
	}

	affected, err := sql.RowsAffected()

	check = (affected == 0)

	if check || (err != nil) {
		msg := "공지 삭제에 실패했습니다!"
		_, err = s.ChannelMessageSend(m.ChannelID, msg)
		if err != nil {
			return err
		}

		return (err)
	}

	msg := "이벤트 등록 완료."
	_, err = s.ChannelMessageSend(m.ChannelID, msg)

	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

// input form : !prefix notice_id
func RemoveEvent(s *discordgo.Session, m *discordgo.MessageCreate) error {
	check, err := CheckAdmin(Hashstring(m.Author.ID))
	if (err != nil) || (!check) {
		fmt.Println(err)
		return err
	}

	event := strings.TrimPrefix(m.Content, "!이벤트 삭제 ")

	args := strings.TrimSpace(event)

	query := `DELETE FROM notice WHERE notice_id = ?`

	result, err := db.Exec(query, args)

	if err != nil {
		return err
	}

	affected, err := result.RowsAffected()

	check = (affected == 0)

	if check || (err != nil) {
		msg := "공지 삭제에 실패했습니다!"
		_, err = s.ChannelMessageSend(m.ChannelID, msg)
		if err != nil {
			return err
		}

		return (err)
	}

	msg := "공지를 성공적으로 삭제하였습니다."
	_, err = s.ChannelMessageSend(m.ChannelID, msg)
	if err != nil {
		return err
	}

	return nil
}

func ShowEvents(s *discordgo.Session, m *discordgo.MessageCreate) {

	embed := &discordgo.MessageEmbed{
		URL:         "https://discord.com/channels/1223177363722862612/1223177363722862616/1239501454348128287",
		Title:       "이벤트 목록",
		Description: "위 링크를 클릭해주세요",
		Color:       0x1f1e33,
		Fields: []*discordgo.MessageEmbedField{
			{
				Name:   "시작날짜",
				Value:  "1999-92-92",
				Inline: true,
			},
			{
				Name:   "종료날짜",
				Value:  "2000-02-02",
				Inline: true,
			},
		},
	}

	_, err := s.ChannelMessageSendEmbed(m.ChannelID, embed)
	if err != nil {
		fmt.Println(err)
		return
	}

}

func ShowEventLists(s *discordgo.Session, m *discordgo.MessageCreate) {

	embed := &discordgo.MessageEmbed{
		Title:       "이벤트 목록",
		Description: "!이벤트 {id} 를 해보세요",
		Color:       0x1f1e33,
		Fields: []*discordgo.MessageEmbedField{
			{
				Name:   "시작날짜",
				Value:  "1999-92-92",
				Inline: true,
			},
			{
				Name:   "종료날짜",
				Value:  "2000-02-02",
				Inline: true,
			},
		},
	}

	_, err := s.ChannelMessageSendEmbed(m.ChannelID, embed)
	if err != nil {
		fmt.Println(err)
		return
	}

}
