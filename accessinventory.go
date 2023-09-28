package main

import "fmt"

func accessInventory(character *Character) {
    for {
        fmt.Println("Inventory Items:")
        if len(character.Inventory) > 0 {
            for i, item := range character.Inventory {
                fmt.Printf("%d. %s\n", i+1, item)
            }
        } else {
            fmt.Println("Inventory is empty.")
        }

        fmt.Println("Menu:")
        fmt.Println("1. Return to Previous Menu")
        fmt.Println("2. Use Potion")
        fmt.Println("3. Exit")

        var choice int
        fmt.Print("Enter your choice: ")
        fmt.Scan(&choice)

        switch choice {
        case 1:
            return
        case 2:
            if len(character.Inventory) > 0 {
                for i, item := range character.Inventory {
                    if item == "Potion" {
                        character.HP += 50
                        if character.HP > character.MaxHP {
                            character.HP = character.MaxHP
                        }
                        character.Inventory = append(character.Inventory[:i], character.Inventory[i+1:]...)
                        fmt.Println("You used a Potion. Current Health Points:", character.HP, "/", character.MaxHP)
                        break
                    }
                }
            } else {
                fmt.Println("No Potions in the inventory.")
            }
        case 3:
            fmt.Println("Exiting the inventory menu.")
            return
        default:
            fmt.Println("Invalid choice.")
        }
    }
}
