package main

func removeMultipleItems(character *Character, item string, quantity int) {
    remaining := quantity
    for i := len(character.Inventory) - 1; i >= 0; i-- {
        if character.Inventory[i] == item && remaining > 0 {
            character.Inventory = append(character.Inventory[:i], character.Inventory[i+1:]...)
            remaining--
        }
    }
}
