package main

import (
	"os"

	"github.com/atakanozceviz/xtoj/actions"
)

func main() {
	port := os.Getenv("PORT")
	if err := actions.Start(port); err != nil {
		panic(err)
	}
}
