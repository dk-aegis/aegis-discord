package service

import (
	"fmt"


)

func GiveMoneyExp(userID string, money int, exp int) error { //돈주는함수

	query := `UPDATE wallet
	SET money = money + ?,
	exp = exp + ?
	WHERE id = ?`

	_ , err := db.Exec(query, money, exp, userID)

	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

