package main

import (
	"fmt"
	"strings"
)

func charCreation() Character {
    var name string
    var class string

    // Demander à l'utilisateur de saisir son nom
    fmt.Print("Enter your character's name: ")
    fmt.Scanln(&name)

    // Mettre en forme le nom (initiale en majuscule, le reste en minuscule)
    name = strings.ToLower(name)

    // Demander à l'utilisateur de choisir sa classe
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

    // Attribuer les points de vie maximum en fonction de la classe
    var maxHP int
    switch class {
    case "Human":
        maxHP = 100
    case "Elf":
        maxHP = 80
    case "Dwarf":
        maxHP = 120
    }

    // Initialiser les points de vie actuels à 50% des points de vie maximum
    currentHP := maxHP / 2

    // Créer le personnage avec les paramètres définis
    character := Character{
        Name:             name,
        Class:            class,
        Level:            1,
        MaxHealthPoints:  maxHP,
        CurrentHealthPoints: currentHP,
        Inventory:         []string{"Punch"}, // Initialiser avec le "Coup de Poing"
        Gold:              100, // Initialiser avec 100 pièces d'or
    }

    return character
}