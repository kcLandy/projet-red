package main

import "fmt"

func forgeMenu(character *Character) {
    fmt.Println("Blacksmith Menu:")
    fmt.Println("1. Craft Adventurer's Hat")
    fmt.Println("2. Craft Adventurer's Tunic")
    fmt.Println("3. Craft Adventurer's Boots")
    fmt.Println("4. Return to Previous Menu")

    var forgeChoice int
    fmt.Print("Enter your choice: ")
    fmt.Scan(&forgeChoice)

    switch forgeChoice {
    case 1, 2, 3:
        equipmentName := ""
        switch forgeChoice {
        case 1:
            equipmentName = "Adventurer's Hat"
        case 2:
            equipmentName = "Adventurer's Tunic"
        case 3:
            equipmentName = "Adventurer's Boots"
        }

        equipment, found := equipmentList[equipmentName]
        if !found {
            fmt.Println("Invalid equipment choice.")
            return
        }

        if CanCraftEquipment(character, equipment) {
            craftEquipment(character, equipment)
        } else {
            fmt.Println("You don't have the necessary resources or gold to craft", equipmentName)
        }
    case 4:
    default:
        fmt.Println("Invalid choice.")
    }
}