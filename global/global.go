package global

import (
	"log"

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
	Bot              BotConfig
	DB               DbConfig
	Role             RoleConfig
	GuildID          string
	WelcomeChannelID string
}

var Discord Config

func InitDiscordConfig() error {
	var env map[string]string

	env, err := godotenv.Read()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	Discord = Config{
		Bot: BotConfig{
			Token: env["BOT_TOKEN"],
			AppID: env["BOT_ID"],
		},
		DB: DbConfig{
			Type:     env["DB_TYPE"],
			User:     env["DB_USER"],
			Password: env["DB_PSWD"],
			Protocol: env["DB_PROTOCOL"],
			Port:     env["DB_PORT"],
			Host:     env["DB_HOST"],
			Name:     env["DB_NAME"],
		},
		Role: RoleConfig{
			ModRoleID:      env["MODROLE_ID"],
			StudyRoleID:    env["STUDYROLE_ID"],
			GraduateRoleID: env["GRADUROLE_ID"],
			StudentRoleID:  env["STUDENTROLE_ID"],
		},
		GuildID:          env["GUILD_ID"],
		WelcomeChannelID: env["WELCOME_CHAN_ID"],
	}

	return nil
}
