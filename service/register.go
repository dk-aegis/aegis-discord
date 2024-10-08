package service

import (
	"discord/global"
	"fmt"
	"time"

	"github.com/bwmarrin/discordgo"
)

func MemberJoin(s *discordgo.Session, m *discordgo.GuildMemberAdd) {

	//뜸들이기
	time.Sleep(time.Second * 5)

	msg := fmt.Sprintf("Welcome To Aegis Server <@%s>! 별명을 '학번이름(재학)' 으로 설정해주세요!", m.User.ID)
	_, err := s.ChannelMessageSend(global.Discord.WelcomeChannelID, msg)
	if err != nil {
		fmt.Println(err)
		return
	}

	//역할 추가되는거 좀 모르겠네
	err = s.GuildMemberRoleAdd(global.Discord.GuildID, m.User.ID, global.Discord.Role.StudentRoleID)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = Regist_user(s, m.User.ID)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func ForkallGuild(s *discordgo.Session, i *discordgo.InteractionCreate) {

	if !CheckRole(i.Member.Roles, global.Discord.Role.ModRoleID) {
		SendInteractionMessage(s, i, "권한이 없습니다")
		return
	}

	MemList, err := s.GuildMembers(i.GuildID, "", 1000)

	if err != nil {
		fmt.Printf("err: %v\n", err)
	}

	for index, member := range MemList {
		err := Regist_user(s, member.User.ID)

		if err != nil {
			fmt.Printf("err: %v\n", err)
			continue
		}

		msg := fmt.Sprintf("%d : success %s ", index, member.Nick)
		fmt.Println(msg)
	}
}

// user ID 를 받아서 db 에 등록합니다.
func Regist_user(s *discordgo.Session, userID string) error {

	tx, err := db.Begin() //transaction on.

	if err != nil {
		fmt.Println(err)
		return err
	}

	var count int
	query := `SELECT COUNT(*) 
	FROM attendance 
	WHERE id = ?`
	err = tx.QueryRow(query, userID).Scan(&count)

	if err != nil {
		fmt.Println(err)
		return err
	}

	if count != 0 {
		tx.Rollback()
		fmt.Println("이미 등록된 회원.", userID)
	}

	wallet_query := "INSERT INTO wallet (id,money,exp) VALUES (?,10000,0)"
	attend_query := "INSERT INTO attendance (id,attend_count,last_seen,conseq_count) VALUES (?,1,CURRENT_DATE,1)"

	_, err = tx.Exec(attend_query, userID)
	if err != nil {
		tx.Rollback()
		fmt.Println(err)
		return err
	}

	_, err = tx.Exec(wallet_query, userID)
	if err != nil {
		tx.Rollback()
		fmt.Println(err)
		return err
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		fmt.Println(err)
		return err
	}

	return nil
}
