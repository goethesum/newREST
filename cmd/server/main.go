package main

import (
	"context"
	"fmt"

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
	if err := db.Ping(context.Background()); err != nil {
		return nil
	}
	fmt.Println("successfully conected and pinged databse")

	return nil
}

func main() {
	fmt.Println("newREST")
	if err := Run(); err != nil {
		fmt.Println(err)
	}
}
