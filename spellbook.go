package main

import "fmt"

func ratHunter(character *Character, skill string) {
    for _, learnedSpell := range character.Skills {
        if learnedSpell == skill {
            fmt.Printf("You already know the spell: %s\n", skill)
            return
        }
    }
    character.Skills = append(character.Skills, skill)
    fmt.Printf("You learned a new spell: %s\n", skill)
}