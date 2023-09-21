package main

import (
    "fmt"
    "math/rand"
    "time"
)

type Character struct {
    Name              string
    Class             string
    Level             int
    MaxHealthPoints   int
    CurrentHealthPoints int
    Inventory          []string
    Gold               int
    Skills             []string
}

type Equipment struct {
    Name     string
    Recipe   map[string]int // Ingrédients nécessaires et leurs quantités
    Cost     int            // Coût en pièces d'or pour la fabrication
    CraftMsg string         // Message de réussite de la fabrication
}

var equipmentList = map[string]Equipment{
    "Adventurer's Hat": {
        Name:   "Adventurer's Hat",
        Recipe: map[string]int{"Crow Feather": 1, "Boar Leather": 1},
        Cost:   5,
        CraftMsg: "You crafted an Adventurer's Hat!",
    },
    "Adventurer's Tunic": {
        Name:   "Adventurer's Tunic",
        Recipe: map[string]int{"Wolf Fur": 2, "Troll Skin": 1},
        Cost:   10,
        CraftMsg: "You crafted an Adventurer's Tunic!",
    },
    "Adventurer's Boots": {
        Name:   "Adventurer's Boots",
        Recipe: map[string]int{"Wolf Fur": 1, "Boar Leather": 1},
        Cost:   7,
        CraftMsg: "You crafted Adventurer's Boots!",
    },
}

func main() {
    rand.Seed(time.Now().UnixNano())

    c1 := charCreation()

    for {
        fmt.Println("Menu:")
        fmt.Println("1. Display Character Information")
        fmt.Println("2. Access Inventory")
        fmt.Println("3. Merchant")
        fmt.Println("4. Blacksmith")
        fmt.Println("5. Check for Death")
        fmt.Println("6. Use Poison Potion")
        fmt.Println("7. Learn Spell: Fireball")
        fmt.Println("8. Exit")

        var choice int
        fmt.Print("Enter your choice: ")
        fmt.Scan(&choice)

        switch choice {
        case 1:
            displayInfo(c1)
        case 2:
            accessInventory(&c1)
        case 3:
            fmt.Println("Merchant Menu:")
            fmt.Println("1. Buy Potion (3 Gold)")
            fmt.Println("2. Buy Poison Potion (6 Gold)")
            fmt.Println("3. Buy Spell Book: Fireball (25 Gold)")
            fmt.Println("4. Buy Wolf Fur (4 Gold)")
            fmt.Println("5. Buy Troll Skin (7 Gold)")
            fmt.Println("6. Buy Boar Leather (3 Gold)")
            fmt.Println("7. Buy Crow Feather (1 Gold)")
            fmt.Println("8. Return to Previous Menu")

            var merchantChoice int
            fmt.Print("Enter your choice: ")
            fmt.Scan(&merchantChoice)

            switch merchantChoice {
            case 1:
                if c1.Gold >= 3 {
                    addInventory(&c1, "Potion")
                    c1.Gold -= 3 // Coût de la potion
                } else {
                    fmt.Println("Not enough gold to buy a Potion.")
                }
            case 2:
                if c1.Gold >= 6 {
                    addInventory(&c1, "Poison Potion")
                    c1.Gold -= 6 // Coût de la potion empoisonnée
                } else {
                    fmt.Println("Not enough gold to buy a Poison Potion.")
                }
            case 3:
                if c1.Gold >= 25 {
                    addInventory(&c1, "Spell Book: Fireball")
                    spellBook(&c1, "Fireball")
                    c1.Gold -= 25 // Coût du livre de sort
                } else {
                    fmt.Println("Not enough gold to buy a Spell Book: Fireball.")
                }
            case 4:
                if c1.Gold >= 4 {
                    addInventory(&c1, "Wolf Fur")
                    c1.Gold -= 4 // Coût de la Fourrure de Loup
                } else {
                    fmt.Println("Not enough gold to buy Wolf Fur.")
                }
            case 5:
                if c1.Gold >= 7 {
                    addInventory(&c1, "Troll Skin")
                    c1.Gold -= 7 // Coût de la Peau de Troll
                } else {
                    fmt.Println("Not enough gold to buy Troll Skin.")
                }
            case 6:
                if c1.Gold >= 3 {
                    addInventory(&c1, "Boar Leather")
                    c1.Gold -= 3 // Coût du Cuir de Sanglier
                } else {
                    fmt.Println("Not enough gold to buy Boar Leather.")
                }
            case 7:
                if c1.Gold >= 1 {
                    addInventory(&c1, "Crow Feather")
                    c1.Gold -= 1 // Coût de la Plume de Corbeau
                } else {
                    fmt.Println("Not enough gold to buy Crow Feather.")
                }
            case 8:
                // Retourner au menu précédent
            default:
                fmt.Println("Invalid choice.")
            }
        case 4:
            blacksmithMenu(&c1)
        case 5:
            checkDead(&c1)
        case 6:
            if containsItem(&c1, "Poison Potion") {
                removeInventory(&c1, "Poison Potion")
                poisonPot(&c1)
            } else {
                fmt.Println("You don't have any Poison Potions in your inventory.")
            }
        case 7:
            if containsItem(&c1, "Spell Book: Fireball") {
                spellBook(&c1, "Fireball")
            } else {
                fmt.Println("You need the Spell Book: Fireball to learn this spell.")
            }
        case 8:
            fmt.Println("Exiting the program.")
            return
        default:
            fmt.Println("Invalid choice.")
        }
    }
}