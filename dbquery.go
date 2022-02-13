package main

import (
	"fmt"

	"github.com/mateors/msql"
)

func GetUser(userEmail, userPassword string) ([]map[string]interface{}, error) {
	qs := fmt.Sprintf("SELECT Email, Password FROM user WHERE Email='%s' AND Password='%s';", userEmail, userPassword)

	row, err := msql.GetAllRowsByQuery(qs, db)
	FetchError(err)
	return row, nil
}
