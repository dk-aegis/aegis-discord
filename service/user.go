package service

import (
	"discord/global"

	"fmt"

	"github.com/bwmarrin/discordgo"
)

func Rolerole(s *discordgo.Session, m *discordgo.MessageCreate) {
	s.ChannelMessageSend(m.ChannelID, "!재학생 또는 !졸업생을 입력해주세요.")

}

func MemberJoin(s *discordgo.Session, m *discordgo.GuildMemberAdd) { //GuildMemberAdd : 서버에 사람 들어오는거 감지
	msg := fmt.Sprintf("Welcome To Aegis Server %s!", m.User.Username)
	s.ChannelMessageSend(global.Discord.WelcomeChannelID, msg)
	//member join 후에 역할부여 handler 추가
}

func Level(s *discordgo.Session, m *discordgo.MessageCreate) {
	s.ChannelMessageSend(m.ChannelID, "현재 해당 기능은 구현되지 않았습니다.")
}

func Graduaterole(s *discordgo.Session, m *discordgo.MessageCreate) {

	Dmchannel, err := s.UserChannelCreate(m.Author.ID)
	if err != nil {
		fmt.Println("실패", err)
		return
	}

	msg := " 1+1= ? "
	_, err = s.ChannelMessageSend(Dmchannel.ID, msg)
	if err != nil {
		fmt.Println("실패", err)
		return
	}

	s.AddHandlerOnce(func(s *discordgo.Session, m *discordgo.MessageCreate) {
		correctmsg := fmt.Sprintf("<@%s>님의 역할 설정 완료 되었습니다.", m.Author.ID)
		if Dmchannel.ID != m.ChannelID { //check the channel the message  received is  Direct Message.
			return
		} else if m.Author.ID == s.State.User.ID {
			return
		}
		if m.Content == global.Discord.ExcutivePrivilege { //give a role if code is correct.
			s.GuildMemberRoleAdd(global.Discord.GuildID, m.Author.ID, global.Discord.GraduateRoleID)
			s.ChannelMessageSend(global.Discord.WelcomeChannelID, correctmsg)

			return
		} else {
			s.ChannelMessageSend(Dmchannel.ID, "무슨 역할을 받고 싶으신건가요")
		}

	})

}

func Studentrole(s *discordgo.Session, m *discordgo.MessageCreate) { //암호 아무거나 입력해도 되는 graduaterole
	Dmchannel, err := s.UserChannelCreate(m.Author.ID)
	if err != nil {
		fmt.Println("실패", err)
		return
	}

	msg := " 1+1= ? "
	_, err = s.ChannelMessageSend(Dmchannel.ID, msg)
	if err != nil {
		fmt.Println("실패", err)
		return
	}

	s.AddHandlerOnce(func(s *discordgo.Session, m *discordgo.MessageCreate) {
		correctmsg := fmt.Sprintf("<@%s>님의 역할 설정 완료 되었습니다.", m.Author.ID)
		if Dmchannel.ID != m.ChannelID { //check the channel the message  received is  Direct Message.
			return
		} else if m.Author.ID == s.State.User.ID {
			return
		}
		s.GuildMemberRoleAdd(global.Discord.GuildID, m.Author.ID, global.Discord.GraduateRoleID)
		s.ChannelMessageSend(global.Discord.WelcomeChannelID, correctmsg)

	})
}
