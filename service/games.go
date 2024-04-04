package service

import (
	"fmt"
	"time"

	"github.com/bwmarrin/discordgo"
)

func Slotmachine(s *discordgo.Session, m *discordgo.MessageCreate) {
	var (
		slot1, slot2, slot3, slot4, slot5, slot6, slot7, slot8, slot9 string = ":white_square_button:", ":white_square_button:", ":white_square_button:", ":white_square_button:", ":white_square_button:", ":white_square_button:", ":white_square_button:", ":white_square_button:", ":white_square_button:"
	)
	msg := fmt.Sprintf("     **SLOTS**\n:white_small_square: :white_small_square: :white_small_square:\n%s %s %s\n%s %s %s\n%s %s %s\n:white_small_square: :white_small_square: :white_small_square:", slot1, slot2, slot3, slot4, slot5, slot6, slot7, slot8, slot9)
	message_, err := s.ChannelMessageSend(m.ChannelID, msg)
	if err != nil {
		fmt.Println("SendMessageError", err)
		return
	}
	time.Sleep(time.Second * 2)

	result := "hi man"

	_, err = s.ChannelMessageEdit(m.ChannelID, message_.ID, result)
	if err != nil {
		fmt.Println("SendMessageError", err)
		return
	}

}
