package test

import (
	"context"
	"fmt"
	"goDatabase"
	"testing"
)

func TestInsertSql(t *testing.T) {
	db := goDatabase.GetConnection()
	defer db.Close()

	ctx := context.Background()

	script := "insert into customer(id , name) values ('1' , 'reihan')"
	_, err := db.ExecContext(ctx, script)
	if err != nil {
		panic(err)
	}

	fmt.Println("Success insert new customer")

}
