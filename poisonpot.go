package main

import (
	"fmt"
	"time"
)

func poisonPot(character *Character) {
    fmt.Println("You are poisoned! Taking 10 points of poison damage per second for 3 seconds.")
    
    for i := 0; i < 3; i++ {
        character.CurrentHealthPoints -= 10
        if character.CurrentHealthPoints < 0 {
            character.CurrentHealthPoints = 0
        }
        fmt.Printf("Current Health Points: %d / %d\n", character.CurrentHealthPoints, character.MaxHealthPoints)
        time.Sleep(1 * time.Second)
    }
    
    fmt.Println("The poison has worn off.")
}