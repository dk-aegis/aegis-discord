package service

import (
	"discord/global"

	"fmt"

	"github.com/bwmarrin/discordgo"
)

func MemberJoin(s *discordgo.Session, m *discordgo.GuildMemberAdd) { //GuildMemberAdd : 서버에 사람 들어오는거 감지
	msg := fmt.Sprintf("Welcome To Aegis Server %s!\n\n\n재학생이면 \"재학생\"을 입력하시고 졸업생이면 \"졸업생\"을 입력하세요.", s.State.User.Username)
	s.ChannelMessageSend(global.Discord.WelcomeChannelID, msg)
	

	//member join 후에 역할부여 handler 추가
}

func RoleSelect(s *discordgo.Session, m *discordgo.MessageUpdate) {
	if 
	

	

	if m.Content == "재학생" {

		s.GuildMemberRoleAdd(m.GuildID, m.Author.ID, global.Discord.StudentRoleID)
		s.ChannelMessageSend(global.Discord.WelcomeChannelID, "설정 완료")

	} else if m.Content == "졸업생" {

		s.GuildMemberRoleAdd(m.GuildID, m.Author.ID, global.Discord.GraduateRoleID)
		s.ChannelMessageSend(global.Discord.WelcomeChannelID, "설정 완료")

	} else {
		return
	} // 예외의 입력이 들어왔을때인데 날리는 메시지, 근데 서버 들어오는 메세지 보고 이거 날려버림
	// 그리고 핸들러 비활성화 시키는 방법을 아직 못찾음...

}

func MemberUpdate(s *discordgo.Session, m *discordgo.GuildMemberUpdate) {

}

func Level(s *discordgo.Session, m *discordgo.MessageCreate) {
	s.ChannelMessageSend(m.ChannelID, "현재 해당 기능은 구현되지 않았습니다.")
}
