package test

import (
	"context"
	"fmt"
	"goDatabase"
	"goDatabase/model"
	"goDatabase/repository/impl"
	"testing"
)

func TestCommentInsert(t *testing.T) {
	commentRepository := impl.NewCommentRepository(goDatabase.GetConnection())

	ctx := context.Background()
	comment := model.Comment{
		Email:   "repository@test.com",
		Comment: "Test Repository",
	}

	result, err := commentRepository.Insert(ctx, comment)
	if err != nil {
		panic(err)
	}

	fmt.Println(result)
}

func TestFindById(t *testing.T) {
	commentRepository := impl.NewCommentRepository(goDatabase.GetConnection())

	comment, err := commentRepository.FindById(context.Background(), 12)
	if err != nil {
		panic(err)
	}

	fmt.Println(comment)
}

func TestFindAll(t *testing.T) {
	commentRepository := impl.NewCommentRepository(goDatabase.GetConnection())

	comments, err := commentRepository.FindAll(context.Background())
	if err != nil {
		panic(err)
	}

	for _, comment := range comments {
		fmt.Println(comment)
	}
}
