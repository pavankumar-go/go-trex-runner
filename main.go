package main

import (
	"log"
	"os"

	"github.com/go-dino/game"
)

func main() {
	if err := game.Setup(); err != nil {
		log.Fatal("Setup creation failed : ", err)
		os.Exit(2)
	}
}
