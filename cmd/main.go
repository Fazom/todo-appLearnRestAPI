package main

import (
	"log"

	"github.com/Fazom/todo-appLearnRestAPI.git"
)

func main() {
	srv := new(todo.Server)
	if err := srv.Run("8080"); err != nil {
		log.Fatal(err)
	}
}
