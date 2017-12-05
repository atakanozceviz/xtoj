package main

import (
	"log"
	"os"

	"github.com/atakanozceviz/xtoj/actions"
)

func main() {
	port := os.Getenv("PORT")
	if err := actions.Start(port); err != nil {
		log.Println(err)
	}
}
