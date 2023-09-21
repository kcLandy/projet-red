package main

import "fmt"

func spellBook(character *Character, spell string) {
    for _, learnedSpell := range character.Skills {
        if learnedSpell == spell {
            fmt.Printf("You already know the spell: %s\n", spell)
            return
        }
    }
    character.Skills = append(character.Skills, spell)
    fmt.Printf("You learned a new spell: %s\n", spell)
}