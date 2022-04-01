package main

import (
	"context"
	"fmt"

	"github.com/goethesum/newREST/internal/comment"
	"github.com/goethesum/newREST/internal/comment/db"
)

// Run - is going to be responsible for
// the installation and startup of our
// go application
func Run() error {
	fmt.Println("starting the application")

	db, err := db.NewDatabse()
	if err != nil {
		fmt.Println("Failed to connect to the database")
		return err
	}

	if err := db.MigrateDB(); err != nil {
		fmt.Println("failed to migrate database")
		return err
	}

	fmt.Println("successfully conected and pinged databse")

	cmtService := comment.NewService(db)

	cmtService.PostComment(
		context.Background(),
		comment.Comment{
			ID:     "f33b420c-b0ed-11ec-b909-0242ac120002",
			Slug:   "manual-test",
			Author: "Author",
			Body:   "hello",
		},
	)

	fmt.Println(cmtService.GetComment(
		context.Background(),
		"f33b420c-b0ed-11ec-b909-0242ac120002",
	))

	return nil
}

func main() {
	fmt.Println("newREST")
	if err := Run(); err != nil {
		fmt.Println(err)
	}
}
