package main

func checkInventoryLimit(character *Character) bool {
    return len(character.Inventory) >= 10
}