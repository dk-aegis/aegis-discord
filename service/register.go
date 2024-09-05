package service

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func ForkallGuild(s *discordgo.Session, i *discordgo.InteractionCreate) {
	Mem, err := s.GuildMembers(i.GuildID, "", 1000)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
	for _, member := range Mem {
		fmt.Println(member.Nick)
	}
}

// user ID 를 받아서 db 에 등록합니다.
func Regist_user(s *discordgo.Session, userID string) error {

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
	err = ta.QueryRow(query, userID).Scan(&count)

	if err != nil {
		fmt.Println(err)
		return err
	}

	if count != 0 {
		ta.Rollback()

		if err != nil {
			fmt.Println("이미 등록된 유저입니다.")
			return err
		}

	}

	wallet_query := "INSERT INTO wallet (player_id,exp,money) VALUES (?,0,10000)"
	attend_query := "INSERT INTO attendance (attend_id,attend_count,last_seen) VALUES (?,1,CURRENT_DATE)"

	_, err = ta.Exec(attend_query, userID)
	if err != nil {
		ta.Rollback()
		fmt.Println(err)
		return err
	}

	_, err = ta.Exec(wallet_query, userID)
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

	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}
