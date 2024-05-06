package service

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/bwmarrin/discordgo"
)

func Slotmachine(s *discordgo.Session, m *discordgo.MessageCreate) {
	_, err := s.ChannelMessageSend(m.ChannelID, "**--SLOTS--**")
	if err != nil {
		fmt.Println("SendMessageError", err)
		return
	}

	var (
		slot1, slot2, slot3, slot4, slot5, slot6, slot7, slot8, slot9 string     = ":white_square_button:", ":white_square_button:", ":white_square_button:", ":white_square_button:", ":white_square_button:", ":white_square_button:", ":white_square_button:", ":white_square_button:", ":white_square_button:"
		slotlist                                                      [3][10]int //슬롯담을2차원배열
		message_                                                      *discordgo.Message
	)
	Emojilist := map[int]string{ //슬롯에 나올 이모지들 1번이 당첨.
		1: ":smile:",
		2: ":cry:",
		3: ":cold_face:",
		4: ":zany_face:",
		5: ":detective:",
	}

	for i := 0; i < 3; i++ {
		for j := 0; j < 10; j++ {
			slotlist[i][j] = rand.Intn(5) + 1
		}
	}

	msg := fmt.Sprintf(":white_small_square: :white_small_square: :white_small_square:\n%s %s %s\n%s %s %s :arrow_left:\n%s %s %s\n:white_small_square: :white_small_square: :white_small_square:", slot1, slot2, slot3, slot4, slot5, slot6, slot7, slot8, slot9)
	message_, err = s.ChannelMessageSend(m.ChannelID, msg)
	if err != nil {
		fmt.Println("SendMessageError", err)
		return
	}

	time.Sleep(time.Second * 2)

	for i := 0; i < 8; i++ {
		slot1, slot2, slot3 = Emojilist[slotlist[0][i+2]], Emojilist[slotlist[1][i+2]], Emojilist[slotlist[2][i+2]]
		slot4, slot5, slot6 = Emojilist[slotlist[0][i+1]], Emojilist[slotlist[1][i+1]], Emojilist[slotlist[2][i+1]]
		slot7, slot8, slot9 = Emojilist[slotlist[0][i]], Emojilist[slotlist[1][i]], Emojilist[slotlist[2][i]]
		msg := fmt.Sprintf(":white_small_square: :white_small_square: :white_small_square:\n%s %s %s\n%s %s %s :arrow_left:\n%s %s %s\n:white_small_square: :white_small_square: :white_small_square:", slot1, slot2, slot3, slot4, slot5, slot6, slot7, slot8, slot9)
		_, err = s.ChannelMessageEdit(m.ChannelID, message_.ID, msg)
		time.Sleep(time.Millisecond * 300)
		if err != nil {
			fmt.Println("SendMessageError", err)
			return
		}
	}

	if slot4 == Emojilist[1] && slot4 == slot5 && slot5 == slot6 {
		_, err = s.ChannelMessageSend(m.ChannelID, "잭팟!")
		if err != nil {
			fmt.Println("SendMessageError", err)
			return
		}
	} else {
		_, err = s.ChannelMessageSend(m.ChannelID, "실패!")
		if err != nil {
			fmt.Println("SendMessageError", err)
			return
		}

	}
}
