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
    fmt.Println("1. Gohano")
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
        fmt.Println("Invalid choice. You must choose one !")
    }
    var (
        HP int
        maxHP int
        attack int
    )

    switch class {
    case "Gohano":
        HP = 100
        maxHP = 100
        attack = 25
    case "Gokloch":
        HP = 80
        maxHP = 80
        attack = 30
    case "Manjle":
        HP = 120
        maxHP = 120
        attack = 20
    }
    character := Character{
        Name: name,
        Class: class,
        Level: 1,
        HP: HP,
        MaxHP: maxHP,
        Inventory: []string{"Punch"},
        Gold: 100,
        Attack: attack,
    }
    return character
}
