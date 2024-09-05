package service

import (
	"fmt"


)


func GiveMoney(userID string, money int) error { //돈주는함수

	query := `UPDATE players
	SET money = money + ?
	WHERE player_id = ?`

	_, err := db.Exec(query, money, userID)

	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

func GiveExp(hashString string, exp int) error {

	query := `UPDATE players
	SET exp = exp + ?
	WHERE player_id = ?`

	_, err := db.Exec(query, exp, hashString)

	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}
