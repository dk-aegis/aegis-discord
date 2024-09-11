package service

import (
	"fmt"
	"strconv"
	"time"

	"github.com/bwmarrin/discordgo"
)

func ShowUserInfo(s *discordgo.Session, i *discordgo.InteractionCreate) {

	At, err := LoadAttendance(i.Member.User.ID)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}
	Wl, err := LoadWallet(i.Member.User.ID)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}

	embed := &discordgo.MessageEmbed{
		Title: "User Infomation",
		Color: 0x00ff00,
		Fields: []*discordgo.MessageEmbedField{
			{
				Name:   "money",
				Value:  strconv.Itoa(Wl.Money),
				Inline: true,
			},
			{
				Name:   "exp",
				Value:  strconv.Itoa(Wl.Exp),
				Inline: true,
			},
			{
				Name:  "출석일수",
				Value: strconv.Itoa(At.Attend_count),
			},
			{
				Name:  "연속출석일수",
				Value: strconv.Itoa(At.Conseq_count),
			},
			{
				Name:  "마지막 출석 날짜",
				Value: At.Lastseen,
			},
			{
				Name: ".",
				Value: time.Now().Format("2006-01-02 15:04:05"),
			},
		},
	}

	err = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Embeds: []*discordgo.MessageEmbed{embed},
		},
	})

	if err != nil {
		fmt.Println("error response", err)
		return
	}

}

func GiveMoneyExp(userID string, money int, exp int) error { //돈주는함수

	query := `UPDATE wallet
	SET money = money + ?,
	exp = exp + ?
	WHERE id = ?`

	_, err := db.Exec(query, money, exp, userID)

	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
