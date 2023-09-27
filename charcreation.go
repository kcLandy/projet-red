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
    fmt.Println("1.Gohano")
    fmt.Println("2. Gokloch")
    fmt.Println("3. Manjle")

    var classChoice int
    fmt.Print("Enter your choice: ")
    fmt.Scan(&classChoice)
    
    switch classChoice {
    case 1:
        class = "Gohano"
    case 2:
        class = "Gokloch"
    case 3:
        class = "Manjle"
    default:
        fmt.Println("Invalid choice. Defaulting to Human class.")
        class = "Gohano"
    }
    var maxHP int
    switch class {
    case "Gohano":
        maxHP = 100
    case "Gokloch":
        maxHP = 80
    case "Manjle":
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