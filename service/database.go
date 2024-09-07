package service

import (
	"discord/global"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"

	"fmt"

)

type DbConfig struct {
	Type     string
	Host     string
	Port     string
	User     string
	Password string
	Name     string
	Protocol string
}

var db *sql.DB

func InitDatabase() error {
	var dc global.DbConfig  = global.Discord.DB

	

	auth := fmt.Sprintf("%s:%s@%s(%s:%s)/%s",
		dc.User, dc.Password, dc.Protocol, dc.Host, dc.Port, dc.Name)

	var err error
	db, err = sql.Open(dc.Type, auth)
	if err != nil {
		return err
	}

	return nil
}

//에러 처리코드는 나중에 작성

type Attendance struct {
	Id           string
	Attend_count int
	Lastseen     string
	Conseq_count int
}

func LoadAttendance(userID string) (Attendance, error) {

	query := "SELECT attend_count, last_seen, conseq_count FROM attendance WHERE id = ?"
	info := Attendance{
		Id: userID,
	}

	err := db.QueryRow(query, userID).Scan(&info.Attend_count, &info.Lastseen, &info.Conseq_count)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return Attendance{}, err
	}

	return info, nil
}

type Wallet struct {
	Id    string
	Money int
	Exp   int
}

func LoadWallet(userID string) (Wallet, error) {
	query := "SELECT money, exp FROM wallet WHERE id = ?"
	info := Wallet{
		Id: userID,
	}

	err := db.QueryRow(query, userID).Scan(&info.Money, &info.Exp)
	if err != nil {
		return Wallet{}, err
	}
	return info, nil
}

func DBclose() {
	db.Close()
}
