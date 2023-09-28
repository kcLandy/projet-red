package main

import "fmt"

func checkDead(character *Character) {
    if character.HP <= 0 {
        fmt.Println("You died!")
        character.HP = character.MaxHP / 2
        fmt.Println("You have been revived with 50% of your max health.")
    }
}