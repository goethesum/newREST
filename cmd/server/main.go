package main

import "fmt"

// Run - is going to be responsible for
// the installation and startup of our
// go application
func Run() error {
	fmt.Println("starting the application")
	return nil
}

func main() {
	fmt.Println("newREST")
	if err := Run(); err != nil {
		fmt.Println(err)
	}
}
