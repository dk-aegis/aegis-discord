package global

import (
	"encoding/json"
	"os"

	"github.com/joho/godotenv"
)

type DiscordConfig struct {
	GuildID           string `json:"guild_id"`
	WelcomeChannelID  string `json:"welcome_channel_id"`
	ModeratorRoleID   string `json:"moderator_role_id"`
	StudyRoleID       string `json:"study_role_id"`
	GraduateRoleID    string `json:"graduate_role_id"`
	StudentRoleID     string `json:"student_role_id"`
	GeneralRoleID     string `json:"general_role_id"`
}

var Config DiscordConfig

func InitDiscordConfig() error {
	err := godotenv.Load()
	if err != nil {
		return err
	}


	return nil
}
