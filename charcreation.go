package main

import (
	"fmt"
	"strings"
)

func charCreation() Character {
    var name string
    var class string

    fmt.Print("Enter your character's name: ")
    fmt.Scanln(&name)
    name = strings.ToLower(name)
    fmt.Println("Choose your character's class:")
    fmt.Println("1. Human")
    fmt.Println("2. Elf")
    fmt.Println("3. Dwarf")

    var classChoice int
    fmt.Print("Enter your choice: ")
    fmt.Scan(&classChoice)
    
    switch classChoice {
    case 1:
        class = "Human"
    case 2:
        class = "Elf"
    case 3:
        class = "Dwarf"
    default:
        fmt.Println("Invalid choice. Defaulting to Human class.")
        class = "Human"
    }
    var maxHP int
    switch class {
    case "Human":
        maxHP = 100
    case "Elf":
        maxHP = 80
    case "Dwarf":
        maxHP = 120
    }
    currentHP := maxHP / 2
    character := Character{
        Name:             name,
        Class:            class,
        Level:            1,
        MaxHealthPoints:  maxHP,
        CurrentHealthPoints: currentHP,
        Inventory:         []string{"Punch"},
        Gold:              100,
    }
    return character
}