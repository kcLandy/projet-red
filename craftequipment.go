package main

import "fmt"

func craftEquipment(character *Character, equipment Equipment) {
    for ingredient, quantity := range equipment.Recipe {
        removeMultipleItems(character, ingredient, quantity)
    }
    character.Gold -= equipment.Cost
    character.Inventory = append(character.Inventory, equipment.Name)
    fmt.Println(equipment.CraftMsg)
}