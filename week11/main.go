package main

import (
	"fmt"
	"log"

	"example.com/greeting/keyboard"
)

func main() {
	n, err := keyboard.GetFloat()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(n)
}
