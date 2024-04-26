package service

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"

	"github.com/bwmarrin/discordgo"
	_ "github.com/go-sql-driver/mysql"
)

/**
Nickname을 Hash화
DB Table

*/

type At struct {
	NickHash string
	Date     string
	SeqCount int
}

func Attendance(s *discordgo.Session, m *discordgo.MessageCreate) {
	user_id := m.Author.ID
	hash := sha256.Sum256([]byte(user_id)) //해시부분 미완성.(아직코드이해안함)
	hashString := hex.EncodeToString(hash[:])

	/* 닉네임해쉬로부터 출석 정보를 불러옴 */
	at, err := loadAttendance(string(hashString))
	if err != nil {
		msg := fmt.Sprintf("%s님의 출석에 문제가 생겼어요!", m.Member.Nick)
		s.ChannelMessageSend(m.ChannelID, msg)
		return
	}

	_ = at

	/* 중복 출석 체크*/

	/* 연속 출석 체크 */
	setAttendance()

	/* 출석 진행 */
	updateAttendance()
}

/*
*	사용자의 출석 정보를 불러온다.
 */
func loadAttendance(nick string) (At, error) {
	query := fmt.Sprintf("select * from attendance where nick=%s", nick)
	at := At{}
	err := db.QueryRow(query).Scan(at.NickHash, at.Date, at.SeqCount)

	if err != nil {
		return at, err
	}

	return at, err
}

func setAttendance() {

}

func updateAttendance() {

}
