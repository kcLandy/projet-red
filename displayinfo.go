package main

import (
    "fmt"
)

type Character struct {
    Name             string
    Class            string
    Level            int
    MaxHealthPoints     int
    CurrentHealthPoints int
    Inventory         []string
}

func displayInfo(character Character) {
    fmt.Println("Character Information:")
    fmt.Println("Name:", character.Name)
    fmt.Println("Class:", character.Class)
    fmt.Println("Level:", character.Level)
    fmt.Println("Max Hit Points:", character.MaxHealthPoints)
    fmt.Println("Current Hit Points:", character.CurrentHealthPoints)
    fmt.Println("Inventory:", character.Inventory)
}

func main() {
    c1 := Character{
        Name: "John",
        Class: "Warrior",
        Level: 1,
        MaxHealthPoints: 500,
        CurrentHealthPoints: 100,
        Inventory: []string{"Sword", "Shield", "Healing Potion"},
    }

    fmt.Println("Menu:")
    fmt.Println("1. Display Character Information")
    fmt.Println("2. Exit")

    var choice int
    fmt.Print("Enter your choice: ")
    fmt.Scan(&choice)

    switch choice {
    case 1:
        displayInfo(c1)
    case 2:
        fmt.Println("Exiting the program.")
    case 3:
        fmt.Println("Setting.")
    default:
        fmt.Println("Invalid choice.")
    }
}
