package main

import "fmt"

func addInventory(character *Character, item string) {
    if checkInventoryLimit(character) {
        fmt.Println("Inventory is full. You cannot add more items.")
        return
    }

    character.Inventory = append(character.Inventory, item)
    fmt.Printf("You acquired a %s.\n", item)
}
