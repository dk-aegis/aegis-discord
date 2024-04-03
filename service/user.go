package service

import (
	"discord/global"

	"fmt"

	"github.com/bwmarrin/discordgo"
)

func MemberJoin(s *discordgo.Session, m *discordgo.GuildMemberAdd) {
	msg := fmt.Sprintf("Welcome To Aegis Server <@%s>!\n\n'재학생 이시면 !재학생, 졸업생 이시면 !졸업생 을 입력해주세요.", m.User.ID)
	s.ChannelMessageSend(global.Discord.WelcomeChannelID, msg)
	s.GuildMemberRoleAdd(global.Discord.GuildID, m.User.ID, global.Discord.GeneralRoleID)

}

func Level(s *discordgo.Session, m *discordgo.MessageCreate) {
	s.ChannelMessageSend(m.ChannelID, "현재 해당 기능은 구현되지 않았습니다.")
}

func Rolerole(s *discordgo.Session, m *discordgo.MessageCreate) {
	s.ChannelMessageSend(m.ChannelID, "!재학생 또는 !졸업생을 입력해주세요.")

}

func Graduaterole(s *discordgo.Session, m *discordgo.MessageCreate) {

	Dmchannel, err := s.UserChannelCreate(m.Author.ID)
	if err != nil {
		fmt.Println("error", err)
		return
	} else if m.Member.Roles[0] != global.Discord.GeneralRoleID {
		return
	}

	msg := " 1+1= ? " //암호 물어보는 메세지
	_, err = s.ChannelMessageSend(Dmchannel.ID, msg)
	if err != nil {
		fmt.Println("error", err)
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
			s.GuildMemberRoleRemove(global.Discord.GuildID, m.Author.ID, global.Discord.GeneralRoleID)
			s.ChannelMessageSend(global.Discord.WelcomeChannelID, correctmsg)

			return
		} else {
			s.ChannelMessageSend(Dmchannel.ID, "재학생이시죠?")
		}

	})

}

func Studentrole(s *discordgo.Session, m *discordgo.MessageCreate) { //암호 아무거나 입력해도 되는 graduaterole
	Dmchannel, err := s.UserChannelCreate(m.Author.ID)
	if err != nil {
		fmt.Println("error", err)
		return
	} else if m.Member.Roles[0] != global.Discord.GeneralRoleID {
		return
	}

	msg := " 1+1= ? " //암호 물어보는 메세지
	_, err = s.ChannelMessageSend(Dmchannel.ID, msg)
	if err != nil {
		fmt.Println("error", err)
		return
	}

	s.AddHandlerOnce(func(s *discordgo.Session, m *discordgo.MessageCreate) {
		correctmsg := fmt.Sprintf("<@%s>님의 역할 설정 완료 되었습니다.", m.Author.ID)
		if Dmchannel.ID != m.ChannelID { //check the channel the message  received is  Direct Message.
			return
		} else if m.Author.ID == s.State.User.ID {
			return
		}
		s.GuildMemberRoleAdd(global.Discord.GuildID, m.Author.ID, global.Discord.StudentRoleID)
		s.GuildMemberRoleRemove(global.Discord.GuildID, m.Author.ID, global.Discord.GeneralRoleID)
		s.ChannelMessageSend(global.Discord.WelcomeChannelID, correctmsg)

	})
}
