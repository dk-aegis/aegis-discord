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
		fmt.Println("in MemberJoin: ", err)
		return
	}

	err = s.GuildMemberRoleAdd(global.Discord.GuildID, m.User.ID, global.Discord.StudentRoleID)
	if err != nil {
		fmt.Println("In MemberJoin: ", err)
		return
	}

}

func Level(s *discordgo.Session, m *discordgo.MessageCreate) {
	s.ChannelMessageSend(m.ChannelID, "현재 해당 기능은 구현되지 않았습니다.")
}
