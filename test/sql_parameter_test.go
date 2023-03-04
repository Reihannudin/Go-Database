package test

import (
	"context"
	"fmt"
	"goDatabase"
	"testing"
)

func TestExecSqlParameter(t *testing.T) {
	db := goDatabase.GetConnection()
	defer db.Close()

	ctx := context.Background()

	username := "reihan'; DROP TABLE user; #"
	password := "123"

	script := "INSERT INTO user(username, password) VALUES(?, ?)"
	_, err := db.ExecContext(ctx, script, username, password)
	if err != nil {
		panic(err)
	}

	fmt.Println("Success insert new user")
}
