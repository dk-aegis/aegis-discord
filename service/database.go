package service

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"os"
	"time"
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

type UserInfo struct {
	AccountID int
	RegistDay time.Time
}

func LoadAccount(hashed_id string) (UserInfo, error) { //아직 미완성
	query := fmt.Sprintf("SELECT account_id, regist_day FROM account WHERE user_id = %s", hashed_id)
	var userinfo UserInfo

	err := db.QueryRow(query).Scan(userinfo.AccountID, userinfo.RegistDay)

	if err != nil {
		return userinfo, err
	}
	return userinfo, err
}

func DBclose() {
	db.Close()
}
