package service

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func generatesalt() (string, error) {

	salt := make([]byte, 8)
	_, err := rand.Read(salt)

	if err != nil {
		fmt.Println(err)
	}

	return hex.EncodeToString(salt), nil

}

func Regist_user(s *discordgo.Session, m *discordgo.MessageCreate) error {

	msg := "등록중..."
	_, err := s.ChannelMessageSend(m.ChannelID, msg)

	if err != nil {
		fmt.Println(err)
		return err
	}

	salt, err := generatesalt()

	id := m.Author.ID
	user_id := fmt.Sprintf("%s%s", salt, id)
	hash := sha256.Sum256([]byte(user_id)) //userid해싱
	hashString := hex.EncodeToString(hash[:])

	if err != nil {
		fmt.Println(err)
		return err
	}

	ta, err := db.Begin() //transaction on.

	if err != nil {
		ta.Rollback()
		fmt.Println(err)
		return err
	}

	var count int
	query := "SELECT COUNT(*) FROM account WHERE user_id = ?"
	err = ta.QueryRow(query, hashString).Scan(&count)

	if err != nil {
		fmt.Println(err)
		return err
	}

	if count != 0 {
		ta.Rollback() //transaction die.
		msg := "이미 등록되어있습니다."
		_, err := s.ChannelMessageSend(m.ChannelID, msg)

		if err != nil {
			fmt.Println(err)
			return err
		}

	}

	acc_query := "INSERT INTO account (user_id, salt, regist_day) VALUES (?,?,CURRENT_DATE)"
	play_query := "INSERT INTO players (exp,money) VALUES (0,10000)"
	attend_query := "INSERT INTO attendance (attend,last_date) VALUES (0,CURRENT_DATE)"

	_, err = ta.Exec(acc_query, hashString, salt)
	if err != nil {
		ta.Rollback()
		fmt.Println(err)
		return err
	}

	_, err = ta.Exec(play_query)
	if err != nil {
		ta.Rollback()
		fmt.Println(err)
		return err
	}

	_, err = ta.Exec(attend_query)
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
