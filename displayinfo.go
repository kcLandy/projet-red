package main

import "fmt"

func displayInfoChar(character Character) {
    fmt.Println("Character Information:")
    fmt.Println("Name:", character.Name)
    fmt.Println("Class:", character.Class)
    fmt.Println("Level:", character.Level)
    fmt.Println("HP:", character.HP, "/", character.MaxHP)
    fmt.Println("Inventory:", character.Inventory)
    fmt.Println("Gold:", character.Gold)
    fmt.Println("Skills:", character.Skills)
}
