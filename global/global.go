package global

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type BotConfig struct {
	Token string 
	AppID string 
}

type DbConfig struct {
	Type     string
	User     string
	Password string
	Protocol string
	Port     string
	Host     string
	Name     string
}

type RoleConfig struct {
	ModRoleID      string
	StudyRoleID    string
	GraduateRoleID string
	StudentRoleID  string
}

type Config struct {
	Bot BotConfig
	DB DbConfig
	Role RoleConfig
	GuildID string
	WelcomeChannelID string
}

var Discord Config

func InitDiscordConfig() error {
	err := godotenv.Load()
	//조금 더 알아보고 환경변수 할당하는 코드 짜기
	if err != nil {
		return err
	}


	return nil
}
