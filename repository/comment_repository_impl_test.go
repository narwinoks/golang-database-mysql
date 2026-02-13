package repository

import (
	"context"
	"fmt"
	golangdatabase "golang-database"
	"golang-database/entity"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestCommentInsert(t *testing.T) {
	commentRepository := NewCommentRepository(golangdatabase.GetConnection())
	ctx := context.Background()
	comment := entity.Comment{
		Email:   "repository@test.com",
		Comment: "test comment for repository",
	}
	result, err := commentRepository.Insert(ctx, comment)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}
func TestCommentFindById(t *testing.T) {
	commentRepository := NewCommentRepository(golangdatabase.GetConnection())
	comment, error := commentRepository.FindById(context.Background(), 38)
	if error != nil {
		panic(error)
	}
	fmt.Println(comment)
}

func TestCommentFindAll(t *testing.T) {
	commentRepository := NewCommentRepository(golangdatabase.GetConnection())
	comments, error := commentRepository.FindAll(context.Background())
	if error != nil {
		panic(error)
	}
	for _, comment := range comments {
		fmt.Println(comment)
	}
}
