package service

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	_ "github.com/go-sql-driver/mysql"
)

func Attendance(s *discordgo.Session, m *discordgo.MessageCreate) {
	hashString := Hashstring(m.Author.ID)
	at, err := LoadAttendance(hashString)
	if err != nil {
		fmt.Println(err)
		msg := "출석에 문제가 생겼어요!"
		s.ChannelMessageSend(m.ChannelID, msg)
		return
	}

	query := `
	UPDATE attendance
	SET attend_count = attend_count + 1, last_seen = CURRENT_DATE
	WHERE attend_id = ? AND ? != CURRENT_DATE`

	sqlresult, err := db.Exec(query, hashString, at.lastseen)
	if err != nil {
		fmt.Println(err)
		return
	}

	update, err := sqlresult.RowsAffected()
	if err != nil {
		fmt.Println(err)
		return
	}

	if update == 0 {
		msg := "이미 출석을 하셨습니다."
		_, err = s.ChannelMessageSend(m.ChannelID, msg)
		if err != nil {
			fmt.Println(err)
			return
		}

		return
	}

	err = GiveMoney(hashString, 1000)
	if err != nil {
		fmt.Println("money 지급에 문제가 생겼어요")
		return
	}

	err = GiveExp(hashString, 10)
	if err != nil {
		fmt.Println("exp 지급에 문제가 생겼어요")
		return
	}

	msg := "출석처리완료"
	_, err = s.ChannelMessageSend(m.ChannelID, msg)
	if err != nil {
		fmt.Println(err)
		return
	}

}
