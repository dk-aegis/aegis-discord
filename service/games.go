package service

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/bwmarrin/discordgo"
)

func Slotmachine(s *discordgo.Session, i *discordgo.InteractionCreate) {

	userID := i.Member.User.ID
	

	playermoney, err := LoadWallet(userID)

	if err != nil {
		fmt.Println(err)
		return
	}

	if playermoney.Money < 10 {
		SendInteractionMessage(s,i,"돈이 부족해요")
		return
	} else {
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseDeferredChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Flags:   discordgo.MessageFlagsEphemeral,
			},
		})
	}


	_, err = s.ChannelMessageSend(i.ChannelID, "**--SLOTS--**")
	if err != nil {
		fmt.Println("SendMessageError", err)
		return
	}

	var (
		slot1, slot2, slot3, slot4, slot5, slot6, slot7, slot8, slot9 string     = ":white_square_button:", ":white_square_button:", ":white_square_button:", ":white_square_button:", ":white_square_button:", ":white_square_button:", ":white_square_button:", ":white_square_button:", ":white_square_button:"
		slotlist                                                      [3][10]int //슬롯담을2차원배열
		slot_msg                                                      *discordgo.Message
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

	msg := fmt.Sprintf(`:white_small_square: :white_small_square: :white_small_square:
%s %s %s
%s %s %s :arrow_left:
%s %s %s
:white_small_square: :white_small_square: :white_small_square:`, slot1, slot2, slot3, slot4, slot5, slot6, slot7, slot8, slot9)
	slot_msg, err = s.ChannelMessageSend(i.ChannelID, msg)
	if err != nil {
		fmt.Println("SendMessageError", err)
		return
	}

	time.Sleep(time.Second * 2)

	for index := 0; index < 8; index++ {
		slot1, slot2, slot3 = Emojilist[slotlist[0][index+2]], Emojilist[slotlist[1][index+2]], Emojilist[slotlist[2][index+2]]
		slot4, slot5, slot6 = Emojilist[slotlist[0][index+1]], Emojilist[slotlist[1][index+1]], Emojilist[slotlist[2][index+1]]
		slot7, slot8, slot9 = Emojilist[slotlist[0][index]], Emojilist[slotlist[1][index]], Emojilist[slotlist[2][index]]
		msg := fmt.Sprintf(`:white_small_square: :white_small_square: :white_small_square:
%s %s %s
%s %s %s :arrow_left:
%s %s %s
:white_small_square: :white_small_square: :white_small_square:`, slot1, slot2, slot3, slot4, slot5, slot6, slot7, slot8, slot9)
		_, err = s.ChannelMessageEdit(slot_msg.ChannelID, slot_msg.ID, msg)
		time.Sleep(time.Millisecond * 300)
		if err != nil {
			fmt.Println(err)
			return
		}
	}

	//당첨 조건
	if slot4 == slot5 && slot5 == slot6 {
		s.FollowupMessageCreate(i.Interaction,true,&discordgo.WebhookParams{
			Content: "잭팟! (money += 5000)",
		})

		err = GiveMoneyExp(userID, 5000,1)
		if err != nil {
			fmt.Println(err)
			return
		}
	} else {
		s.FollowupMessageCreate(i.Interaction,true,&discordgo.WebhookParams{
			Content: "실패! (money -= 10)",
		})

		err = GiveMoneyExp(userID, -10,1)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}
