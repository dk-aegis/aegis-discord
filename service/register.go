package service

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func Hashstring(str string) string { //hasinghasing
	hash := sha256.Sum256([]byte(str))
	hashString := hex.EncodeToString(hash[:])
	return string(hashString)
}

func Regist_user(s *discordgo.Session, m *discordgo.MessageCreate) error {

	msg := "등록중..."
	_, err := s.ChannelMessageSend(m.ChannelID, msg)

	if err != nil {
		fmt.Println(err)
		return err
	}

	hashString := Hashstring(m.Author.ID)

	ta, err := db.Begin() //transaction on.

	if err != nil {
		ta.Rollback()
		fmt.Println(err)
		return err
	}

	var count int
	query := `SELECT COUNT(*) 
	FROM account 
	WHERE user_id = ?`
	err = ta.QueryRow(query, hashString).Scan(&count)

	if err != nil {
		fmt.Println(err)
		return err
	}

	if count != 0 {
		ta.Rollback()
		msg := "이미 등록되어있습니다."
		_, err := s.ChannelMessageSend(m.ChannelID, msg)

		if err != nil {
			fmt.Println(err)
			return err
		}

	}

	acc_query := "INSERT INTO account (user_id, regist_day) VALUES (?,CURRENT_DATE)"
	play_query := "INSERT INTO players (player_id,exp,money) VALUES (?,0,10000)"
	attend_query := "INSERT INTO attendance (attend_id,attend_count,last_seen) VALUES (?,1,CURRENT_DATE)"

	_, err = ta.Exec(acc_query, hashString)
	if err != nil {
		ta.Rollback()
		fmt.Println(err)
		return err
	}

	_, err = ta.Exec(play_query, hashString)
	if err != nil {
		ta.Rollback()
		fmt.Println(err)
		return err
	}

	_, err = ta.Exec(attend_query, hashString)
	if err != nil {
		ta.Rollback()
		fmt.Println(err)
		return err
	}

	err = ta.Commit()
	if err != nil {
		ta.Rollback()
		fmt.Println(err)
		return err
	}

	msg = "등록완료!"
	_, err = s.ChannelMessageSend(m.ChannelID, msg)

	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}
