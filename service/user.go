package service

import (
	"discord/global"

	"fmt"

	"github.com/bwmarrin/discordgo"
)

func MemberJoin(s *discordgo.Session, m *discordgo.GuildMemberAdd) {
	msg := fmt.Sprintf("Welcome To Aegis Server <@%s>!", m.User.ID)
	_, err := s.ChannelMessageSend(global.Discord.WelcomeChannelID, msg)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = s.GuildMemberRoleAdd(global.Discord.GuildID, m.User.ID, global.Discord.StudentRoleID)
	if err != nil {
		fmt.Println(err)
		return
	}

}

func GiveMoney(hashString string, money int) error { //돈주는함수

	query := `UPDATE players
	SET money = money + ?
	WHERE player_id = ?`

	_, err := db.Exec(query, money, hashString)

	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

func GiveExp(hashString string, exp int) error {

	query := `UPDATE players
	SET exp = exp + ?
	WHERE player_id = ?`

	_, err := db.Exec(query, exp, hashString)

	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

func CheckMoney(s *discordgo.Session, m *discordgo.MessageCreate) {

	id := m.Author.ID
	hashString := Hashstring(id)
	play, err := LoadPlayers(hashString)
	if err != nil {
		fmt.Println(err)
		return
	}
	msg := fmt.Sprintf("<@%s>님이 소지중인 money(은)는 %d 입니다!", id, play.money)
	_, err = s.ChannelMessageSend(m.ChannelID, msg)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func CheckExp(s *discordgo.Session, m *discordgo.MessageCreate) {

	id := m.Author.ID
	hashString := Hashstring(id)
	play, err := LoadPlayers(hashString)
	if err != nil {
		fmt.Println(err)
		return
	}
	msg := fmt.Sprintf("<@%s>님이 소지중인 exp(은)는 %d 입니다!", id, play.exp)
	_, err = s.ChannelMessageSend(m.ChannelID, msg)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func Checkattend(s *discordgo.Session, m *discordgo.MessageCreate) {

	id := m.Author.ID
	hashString := Hashstring(id)
	att, err := LoadAttendance(hashString)
	if err != nil {
		fmt.Println(err)
		return
	}

	msg := fmt.Sprintf("<@%s>님의 출석일수(은)는 %d일 입니다!", id, att.attendacne)
	_, err = s.ChannelMessageSend(m.ChannelID, msg)
	if err != nil {
		fmt.Println(err)
		return
	}
}
