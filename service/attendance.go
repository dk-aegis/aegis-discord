package service

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	_ "github.com/go-sql-driver/mysql"
)

func DoAttendance(s *discordgo.Session, i *discordgo.InteractionCreate) {
	userID := i.Member.User.ID

	at, err := LoadAttendance(userID)

	if err != nil {
		fmt.Println(err)
		msg := "출석에 문제가 생겼어요!"
		s.ChannelMessageSend(i.ChannelID, msg)
		return
	}

	query := `
	UPDATE attendance
	SET attend_count = attend_count + 1, last_seen = CURRENT_DATE
	WHERE attend_id = ? AND ? != CURRENT_DATE`

	sqlresult, err := db.Exec(query, userID, at.Lastseen)
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
		_, err = s.ChannelMessageSend(i.ChannelID, msg)
		if err != nil {
			fmt.Println(err)
			return
		}

		return
	}

	err = GiveMoney(userID, 1000)
	if err != nil {
		fmt.Println("money 지급에 문제가 생겼어요")
		return
	}

	err = GiveExp(userID, 25)
	if err != nil {
		fmt.Println("exp 지급에 문제가 생겼어요")
		return
	}

	msg := "출석처리완료"
	_, err = s.ChannelMessageSend(i.ChannelID, msg)
	if err != nil {
		fmt.Println(err)
		return
	}
}
