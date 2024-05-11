package service

import (
	"github.com/bwmarrin/discordgo"
	_ "github.com/go-sql-driver/mysql"
)

func InitEvent() {
	eventList = make([]Event, 0)
	eventList = append(eventList, Event{
		FromDate:  "2024-03-18",
		UntilDate: "2024-04-01",
		Title:     "리틀 운영진 모집",
		Content:   "2학기를 이끌어갈 리틀 운영진을 뽑습니다",
	})
}

func ShowHomepage(s *discordgo.Session, m *discordgo.MessageCreate) {
	sendEmbedMessage(s, m.ChannelID, "Aegis 홈페이지", "https://dk-aegis.org", 0x00ff00)
}

func HelpMessage(s *discordgo.Session, m *discordgo.MessageCreate) {
	msg := `
		!도움말: 명령어 목록 도움말.

		!이벤트: 현재 진행중인 이벤트.
		
		!출석: 해당 날짜 출석.

		!홈페이지: Aegis 동아리 홈페이지.

		!슬롯: 슬롯머신을 돌립니다.

		!정보등록: 출석체크와 기타 기능을 사용하기 위해 정보를 등록합니다.

		!돈: 현재 보유중인 돈을 출력합니다.

		!경험치: 현재 보유중인 경험치를 출력합니다.

		!출석일수: 현재까지 출석한 일수를 출력합니다.
	`
	sendEmbedMessage(s, m.ChannelID, "명령어 도움말", msg, 0x00ff00)
}
