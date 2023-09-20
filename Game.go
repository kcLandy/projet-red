package main

import "fmt"

type Character struct {
    Name              string
    Class             string
    Level             int
    MaxHealthPoints   int
    CurrentHealthPoints int
    Inventory          []string
}
func displayInfo(Character) {
    fmt.Println("Character Information:")
    fmt.Println("Name:", Character.Name)
    fmt.Println("Class:", Character.Class)
    fmt.Println("Level:", Character.Level)
    fmt.Println("Max Health Points:", Character.MaxHealthPoints)
    fmt.Println("Current Health Points:", Character.CurrentHealthPoints)
    fmt.Println("Inventory:", Character.Inventory)
    fmt.Println("Skills:", Character.Skills) 
}