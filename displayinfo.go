package main

import "fmt"

func displayInfo(character Character) {
    fmt.Println("Character Information:")
    fmt.Println("Name:", character.Name)
    fmt.Println("Class:", character.Class)
    fmt.Println("Level:", character.Level)
    fmt.Println("HP:", character.CurrentHealthPoints, "/", character.MaxHealthPoints)
    fmt.Println("Inventory:", character.Inventory)
    fmt.Println("Gold:", character.Gold) // Afficher l'argent du joueur
    fmt.Println("Skills:", character.Skills) // Afficher les sorts
}
