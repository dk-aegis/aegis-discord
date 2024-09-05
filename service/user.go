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

func GiveMoney(userID string, money int) error { //돈주는함수

	query := `UPDATE players
	SET money = money + ?
	WHERE player_id = ?`

	_, err := db.Exec(query, money, userID)

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
