package service

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"os"
)

var db *sql.DB

type DbConfig struct {
	Type     string `json:"type"`
	Host     string `json:"host"`
	Port     int    `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	Database string `json:"database"`
	Protocol string `json:"protocol"`
}

func InitDatabase() error {
	var dc DbConfig
	file, err := os.Open("./config/db_config.json")

	if err != nil {
		fmt.Println(err)
		return err
	}

	defer file.Close()

	jsonParser := json.NewDecoder(file)
	jsonParser.Decode(&dc)

	auth := fmt.Sprintf("%s:%s@%s(%s:%d)/%s",
		dc.User, dc.Password, dc.Protocol, dc.Host, dc.Port, dc.Database)

	db, err = sql.Open(dc.Type, auth)
	if err != nil {
		return err
	}

	return nil
}

type userinfo struct {
	AccountID string
	RegistDay string
}

func LoadAccount(hashed_id string) (userinfo, error) {
	query := "SELECT user_id, regist_day FROM account WHERE acc_id = ?"
	var info userinfo

	err := db.QueryRow(query, hashed_id).Scan(&info.AccountID, &info.RegistDay)
	if err != nil {
		return userinfo{}, err
	}

	return info, nil

}

type attendance struct {
	attendacne int
	lastseen   string
}

func LoadAttendance(hashed_id string) (attendance, error) {
	query := "SELECT attend_count, last_seen FROM attendance WHERE attend_id = ?"
	var info attendance

	err := db.QueryRow(query, hashed_id).Scan(&info.attendacne, &info.lastseen)
	if err != nil {
		return attendance{}, err
	}

	return info, nil

}

type players struct {
	money int
	exp   int
}

func LoadPlayers(hashed_id string) (players, error) {
	query := "SELECT money,exp FROM players WHERE player_id = ?"
	var info players

	err := db.QueryRow(query, hashed_id).Scan(&info.money, &info.exp)
	if err != nil {
		return players{}, err
	}

	return info, nil

}

type notice struct { //공지는 따로 손을 봐야겠음
	id        int
	fromdate  string
	untildate string
	title     string
	content   string
}

func LoadNotice(hashed_id string) (notice, error) {
	query := "SELECT notice_id,fromdate,untildate,title,content FROM notice WHERE ?"
	var info notice

	err := db.QueryRow(query, hashed_id).Scan(&info.id, &info.fromdate, &info.untildate, &info.title, &info.content)
	if err != nil {
		return notice{}, err
	}

	return info, nil
}

func CheckAdmin(hashed_id string) (bool, error) {
	query := "SELECT COUNT(*) FROM root WHERE admin_id = ?"

	var count int

	err := db.QueryRow(query, hashed_id).Scan(&count)
	if err != nil {
		fmt.Println("관리자 인증 오류 발생")
		return false, err
	}

	if count != 0 {
		return true, nil
	} else {
		return false, nil
	}

}

func DBclose() {
	db.Close()
}
