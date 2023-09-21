package main

func removeInventory(character *Character, item string) {
    for i, currentItem := range character.Inventory {
        if currentItem == item {
            character.Inventory = append(character.Inventory[:i], character.Inventory[i+1:]...)
            break
        }
    }
}