package service

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/bwmarrin/discordgo"
	_ "github.com/go-sql-driver/mysql"
)

func DoAttendance(s *discordgo.Session, i *discordgo.InteractionCreate) {
	userID := i.Member.User.ID
	var current_date string
	today := time.Now().Format("2006-01-02") //오늘 날짜

	query := `
	SELECT last_seen FROM attendance WHERE id = ?
	`
	db.QueryRow(query,userID).Scan(&current_date)

	if current_date == today { //출석함??
		msg := "이미 출석을 하셨습니다."
		SendInteractionMessage(s, i, msg)
		return
	}

	query = `
	UPDATE attendance
	SET attend_count = attend_count + 1, 
		conseq_count = CASE 
			WHEN last_seen = DATE_SUB(CURRENT_DATE, INTERVAL 1 DAY) THEN conseq_count + 1
			ELSE 1
		END,
		last_seen = CURRENT_DATE
	WHERE id = ? AND last_seen != CURRENT_DATE;`	

	_ , err := db.Exec(query, userID)
	if err != nil {
		fmt.Println("Attend | sql 실행하는데 문제가 생겼어요 : ",err)
	
		return
	}


	err = GiveMoneyExp(userID, 1000, 25)
	if err != nil {
		fmt.Println("지급에 문제가 생겼어요" , err , userID)
		return
	}

	At, err := LoadAttendance(userID)
	if err != nil {
		fmt.Println("출석 목록 불러오는데 문제가 생겼어요", err , userID)
		return
	}

	msglist := map[int]string{
		1: "지금까지 %d일 출석하셨고, 연속 출석 기록은 %d일입니다. 계속 이어가세요!",
		2: "출석 %d일째! 연속 %d일 출석 중입니다. 계속해서 열심히 참여해 주세요!",
		3: "%d일 동안 꾸준히 출석 중입니다! 연속으로 %d일째 출석 중이에요!",
		4: "오늘로 %d일째 출석! 연속으로 %d일간 빠지지 않고 출석하셨습니다!",
		5: "멋집니다! %d일째 출석하며, 연속 출석 기록은 %d일입니다!",
		6: "%d일째 출석 성공! 연속 출석 %d일을 달성했습니다!",
		7: "현재 %d일째 출석 중이고, 연속 출석은 %d일째입니다! 앞으로도 계속 출석하세요!",
	}
	
	randmsg := msglist[rand.Intn(7)+1] //범위 잘못 정함. 

	msg := fmt.Sprintf(randmsg, At.Attend_count, At.Conseq_count)
	SendInteractionMessage(s, i, msg)
}
