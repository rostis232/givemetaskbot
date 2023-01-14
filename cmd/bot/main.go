package main

import (
	"github.com/rostis232/givemetaskbot/internal/pkg/app"
	"log"
)

func main() {
	a, err := app.NewApp()
	if err != nil {
		log.Fatal(err)
	}

	if err := a.Run(); err != nil {
		log.Fatalln(err)
	}

}
