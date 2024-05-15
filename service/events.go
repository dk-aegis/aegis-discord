package service

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/bwmarrin/discordgo"
)

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

// input form !prefix Title Content 원본글URL fromdate untildate
func CreateEvent(s *discordgo.Session, m *discordgo.MessageCreate) {

	check, err := CheckAdmin(Hashstring(m.Author.ID)) //권한체크  권한이 없다는것도 출력해야됨.
	if (err != nil) || (!check) {
		fmt.Println(err)
		return
	}

	args := strings.Split(m.Content[18:], "|") //'!이벤트 등록' 이 8문자를 죽이고 시작 , 구분자는 | (shift+\)

	if len(args) < 5 {
		msg := "인수가 부족해요"
		_, err = s.ChannelMessageSend(m.ChannelID, msg) //인수가 부족한가??
		if err != nil {
			return
		}
		return
	}

	query := `INSERT INTO notice (title, content, noticeurl, fromdate, untildate)
	VALUES (?,?,?,?,?)`

	sql, err := db.Exec(query, strings.TrimSpace(args[0]), strings.TrimSpace(args[1]), strings.TrimSpace(args[2]), strings.TrimSpace(args[3]), strings.TrimSpace(args[4]))

	if err != nil {
		fmt.Println(err)
		return
	}

	affected, err := sql.RowsAffected()

	check = (affected == 0)

	if check || (err != nil) {
		msg := "이벤트 등록을 실패했습니다!"
		_, err = s.ChannelMessageSend(m.ChannelID, msg)
		if err != nil {
			return
		}

		return
	}

	msg := "이벤트 등록 완료."
	_, err = s.ChannelMessageSend(m.ChannelID, msg)

	if err != nil {
		fmt.Println(err)
		return

	}
}

// input form : !prefix notice_id
func RemoveEvent(s *discordgo.Session, m *discordgo.MessageCreate) error {
	check, err := CheckAdmin(Hashstring(m.Author.ID))
	if (err != nil) || (!check) {
		fmt.Println(err)
		return err
	}

	args := strings.TrimSpace(m.Content[18:])

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

	var title, content, fromdate, untildate, url, query string
	var id int
	var count int
	var err error

	query = "SELECT COUNT(*) FROM notice"
	err = db.QueryRow(query).Scan(&count)

	if err != nil {
		fmt.Println(err)
		return
	}

	if count == 0 { //공지가 없어요 ㅠㅠ
		fmt.Println(err)
		return
	}

	args, _ := strings.CutPrefix(m.Content, "!이벤트")

	args = strings.ReplaceAll(args, " ", "")

	if args == "" {
		query = `SELECT title, content, fromdate, untildate, noticeurl FROM notice ORDER BY notice_id DESC LIMIT ?`
		id = 1
	} else {
		num, err := strconv.Atoi(args)

		if err != nil { //숫자를 입력해주세요~!!
			fmt.Println(err)
			return
		}

		query = `SELECT title, content, fromdate, untildate, noticeurl FROM notice WHERE notice_id = ?`
		id = num
	}

	err = db.QueryRow(query, id).Scan(&title, &content, &fromdate, &untildate, &url)

	if err != nil {
		fmt.Println(err)
		return
	}

	url = strings.ReplaceAll(url, "\u2060", "")

	embed := &discordgo.MessageEmbed{ //URL 아닌 에러 처리해야됨.
		Title:       title,
		Description: content,
		Color:       0x1f1e33,
		URL:         url,
		Fields: []*discordgo.MessageEmbedField{
			{
				Name:   "시작날짜",
				Value:  fromdate,
				Inline: true,
			},
			{
				Name:   "종료날짜",
				Value:  untildate,
				Inline: true,
			},
		},
	}

	_, err = s.ChannelMessageSendEmbed(m.ChannelID, embed)
	if err != nil {
		fmt.Println(err)
		return
	}

}

func ShowEventLists(s *discordgo.Session, m *discordgo.MessageCreate) {

	embed := &discordgo.MessageEmbed{
		Title:       "이벤트 목록",
		Description: "!이벤트 num 를 입력해보세요",
		Color:       0x1f1e33,
		Fields:      []*discordgo.MessageEmbedField{},
	}

	query := `SELECT notice_id, title, fromdate, untildate FROM notice`

	rows, err := db.Query(query)

	if err != nil {
		fmt.Println(err)
		return
	}

	for rows.Next() {
		var notice int
		var title, fromdate, untildate string

		err := rows.Scan(&notice, &title, &fromdate, &untildate)

		if err != nil {
			fmt.Println(err)
			return
		}
		name := fmt.Sprintf("%d: %s", notice, title)
		value := fmt.Sprintf("%s~%s", fromdate, untildate)

		field := &discordgo.MessageEmbedField{
			Name:   name,
			Value:  value,
			Inline: true,
		}

		embed.Fields = append(embed.Fields, field)

	}
	_, err = s.ChannelMessageSendEmbed(m.ChannelID, embed)
	if err != nil {
		fmt.Println(err)
		return
	}

}
