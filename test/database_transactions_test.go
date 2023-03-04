package test

import (
	"context"
	"fmt"
	"goDatabase"
	"strconv"
	"testing"
)

func TestTransactions(t *testing.T) {
	db := goDatabase.GetConnection()
	defer db.Close()

	ctx := context.Background()
	tx, err := db.Begin()

	if err != nil {
		panic(err)
	}

	script := "INSERT INTO comments(email, comment) VALUES(?, ?)"

	// do transaction
	for i := 0; i < 10; i++ {
		email := "reedbulls91" + strconv.Itoa(i) + "@gmail.com"
		comment := "Komentar ke " + strconv.Itoa(i)

		result, err := tx.ExecContext(ctx, script, email, comment)
		if err != nil {
			panic(err)
		}

		id, err := result.LastInsertId()
		if err != nil {
			panic(err)
		}

		fmt.Println("Comment Id ", id)
	}

	err = tx.Rollback()
	if err != nil {
		panic(err)
	}
}
