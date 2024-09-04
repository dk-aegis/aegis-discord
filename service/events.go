package service

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func sendEmbedMessage(s *discordgo.Session, chid, title, content string, color int) {
	embed := &discordgo.MessageEmbed{
		Title:       title,
		Description: content,
		Color:       color, // Green color
	}

	// Send the embedded message to the channel where the command was received.
	_, err := s.ChannelMessageSendEmbed(chid, embed)
	if err != nil {
		fmt.Println("Error sending embedded message:", err)
		return
	}
}
