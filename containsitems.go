package main

func containsItem(character *Character, item string) bool {
    for _, currentItem := range character.Inventory {
        if currentItem == item {
            return true
        }
    }
    return false
}