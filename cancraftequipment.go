package main

func CanCraftEquipment(character *Character, equipment Equipment) bool {
    for ingredient, quantity := range equipment.Recipe {
        if countItem(character.Inventory, ingredient) < quantity {
            return false
        }
    }
    return character.Gold >= equipment.Cost && len(character.Inventory) < 10
}