package main

import (
	"fmt"
	"log"

	"github.com/7minutech/gator_go/internal/config"
)

func main() {
	myConfig, err := config.Read()
	if err != nil {
		log.Fatal(err)
	}
	myConfig.SetUser("7minutech")
	myConfig, err = config.Read()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Username: %s, dbUrl: %s", myConfig.CurrentUserName, myConfig.DbUrl)
}
