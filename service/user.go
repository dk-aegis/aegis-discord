package service

import (
	"discord/global"
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func MemberJoin(s *discordgo.Session, m *discordgo.GuildMemberUpdate) {

	msg := fmt.Sprintf("Welcome To Aegis Server %s!\n서버가 자동으로 역할을 설정했어요!", m.Nick)
	s.ChannelMessageSend(global.Discord.WelcomeChannelID, msg)
	err := s.GuildMemberRoleAdd(
		global.Discord.GuildID,
		m.Member.User.ID,
		global.Discord.StudentRoleID,
	)
	if err != nil {
		s.ChannelMessageSend(global.Discord.WelcomeChannelID, "ㅠㅠ (재학생) 역할 설정을 실패했어요!")
		fmt.Println(err)
	}
}

func MemberUpdate(s *discordgo.Session, m *discordgo.GuildMemberUpdate) {

}

func Level(s *discordgo.Session, m *discordgo.MessageCreate) {
	s.ChannelMessageSend(m.ChannelID, "현재 해당 기능은 구현되지 않았습니다.")
}
