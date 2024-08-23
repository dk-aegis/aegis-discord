package service

import (
	"github.com/bwmarrin/discordgo"
	_ "github.com/go-sql-driver/mysql"
)

func ShowHomepage(s *discordgo.Session, m *discordgo.MessageCreate) {
	sendEmbedMessage(s, m.ChannelID, "Aegis 홈페이지", "https://dk-aegis.org", 0x00ff00)
}

func HelpMessage(s *discordgo.Session, m *discordgo.MessageCreate) {
	msg := `
		!도움말: 명령어 목록 도움말을 확인합니다.

		!이벤트: 현재 진행중인 이벤트를 확인합니다.
		
		!출석: 출석체크를 합니다.

		!홈페이지: Aegis 동아리 홈페이지.

		!슬롯: 슬롯머신을 돌립니다.

		!정보등록: 출석체크와 기타 기능을 사용하기 위해 정보를 등록합니다.

		!돈: 현재 보유중인 돈을 확인합니다.

		!경험치: 현재 보유중인 경험치를 확인합니다.

		!출석일수: 현재까지 출석한 일수를 확인합니다.

		!이벤트 목록: 이벤트 목록을 출력합니다.

		(권한필요)!이벤트 {등록} {제목} {내용} {공지글URL} {시작날짜} {종료날짜}: 이벤트를 등록합니다

		(권한필요)!이벤트 삭제 {id}: 이벤트를 삭제합니다.
	`
	sendEmbedMessage(s, m.ChannelID, "명령어 도움말", msg, 0x00ff00)
}
