package main

import (
	"fmt"
	"log"

	"github.com/haruhikoM/govent/atnd"
)

func main() {
	events, err := atnd.Get("golang")
	if err != nil {
		log.Fatal(err)
	}
	for _, item = range events {
		fmt.Println(event)
	}
}
