package service

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"os"

	//"golang.org/x/text/date"
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

//에러 처리코드는 나중에 작성 

type Attendance struct {
	Id string
	Attend_count int
	Lastseen	string
	Conseq_count int
}	

func LoadAttendance(userID string) (Attendance, error) {

	query := "SELECT attend_count, last_seen conseq_count FROM attendance WHERE attend_id = ?"
	info := Attendance{
		Id: userID,	
	}

	err := db.QueryRow(query, userID).Scan(&info.Attend_count, &info.Lastseen, &info.Conseq_count)
	if err != nil {
		return Attendance{}, err
	}

	return info, nil
}

type Players struct {
	Id 	string
	Money int
	Exp   int
}

func LoadPlayers(userID string) (Players, error) {
	query := "SELECT money,exp FROM players WHERE id = ?"
	info := Players{
		Id: userID,
	}

	err := db.QueryRow(query, userID).Scan(&info.Money, &info.Exp)
	if err != nil {
		return Players{}, err
	}
	return info, nil
}


func DBclose() {
	db.Close()
}
