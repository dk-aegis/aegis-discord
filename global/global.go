package global

import (
	"encoding/json"
	"os"
)

type DiscordConfig struct {
	GuildID           string `json:"guild_id"`
	WelcomeChannelID  string `json:"welcome_channel_id"`
	ModeratorRoleID   string `json:"moderator_role_id"`
	StudyRoleID       string `json:"study_role_id"`
	GraduateRoleID    string `json:"graduate_role_id"`
	StudentRoleID     string `json:"student_role_id"`
	GeneralRoleID     string `json:"general_role_id"`
	ExcutivePrivilege string `json:"executive_privilege"`
}

var Discord DiscordConfig

func InitDiscordConfig() error {
	file, err := os.Open("./config/discord.json")
	if err != nil {
		return err
	}

	defer file.Close()

	jsonParser := json.NewDecoder(file)
	jsonParser.Decode(&Discord)

	return nil
}
