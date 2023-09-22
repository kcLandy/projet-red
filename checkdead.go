package main

import "fmt"

func checkDead(character *Character) {
    if character.CurrentHealthPoints <= 0 {
        fmt.Println("You died!")
        character.CurrentHealthPoints = character.MaxHealthPoints / 2
        fmt.Println("You have been revived with 50% of your max health.")
    }
}