package main

import (
	"fmt"
	"strings"
)

func startTheGame() {
	var input string
	fmt.Println("Start the Game ?", "Yes","/","No")
	_, err := fmt.Scanln(&input)
	if err != nil {
		fmt.Println("WTF BROOO ???", err)
		return
	}

	input = strings.ToLower(input)
	if input == "Yes" {
		fmt.Println("Starting the Game...")
	} else if input == "No" {
		fmt.Println("You're a Pussy !!")
	}
}
