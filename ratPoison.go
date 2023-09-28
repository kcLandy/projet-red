package main

import (
	"fmt"
	"time"
)

func ratPoison(character *Character) {
    fmt.Println("You are poisoned! Taking 10 points of poison damage per second for 3 seconds.")
    
    for i := 0; i < 3; i++ {
        character.HP -= 10
        if character.HP < 0 {
            character.MaxHP = 0
        }
        fmt.Printf("Current Health Points: %d / %d\n", character.HP, character.MaxHP)
        time.Sleep(1 * time.Second)
    }
    
    fmt.Println("The poison has worn off.")
}