package main

import (
	"discord/service"

	"github.com/bwmarrin/discordgo"
)

// This function will be called (due to AddHandler above) every time a new
// message is created on any channel that the authenticated bot has access to.
func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

	// Ignore all messages created by the bot itself
	// This isn't required in this specific example but it's a good practice.
	if m.Author.Bot {
		return
	}

	if m.Content == "!이벤트" {
		go service.ShowEvent(s, m)
	} else if m.Content == "!홈페이지" {
		go service.ShowHomepage(s, m)
	} else if m.Content == "!도움말" {
		go service.HelpMessage(s, m)
	} else if m.Content == "!출석" {
		go service.Attendance(s, m)
	} else if m.Content == "!출석체크" {
		go service.Attendance(s, m)
	} else if m.Content == "!슬롯" {
		go service.Slotmachine(s, m)
	} else if m.Content == "!정보등록" {
		go service.Regist_user(s, m)
	}

}
