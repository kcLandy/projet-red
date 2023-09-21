package main

import "fmt"

const (
	Warrior = "Warrior"
	Archer = "Archer"
	Mage = "Mage"
    Elf = "Elf"
)

type Character struct {
    name string
    class string
    level int
    maxHP  int
    hP int
    inventory []string
}
func displayInfo(Player Character) {
    fmt.Println("Player's Info:")
    fmt.Println("Name:", Player.name)
    fmt.Println("Class:", Player.class)
    fmt.Println("Level:", Player.level)
    fmt.Println("HP:", Player.hP, "/", Player.maxHP)
    fmt.Println("Inventory:", Player.inventory)
}
func accessInventory(Player Character) {
    fmt.Println("Menu:")
    fmt.Println("1.Inventory")
    fmt.Println("2.Return")
    fmt.Println("3.Exit")

    var choice int
        fmt.Print("Select an option: ")
        fmt.Scan(&choice)
}
func main(){
    Player := Character{
        name: "Captain",
        class: Elf,
        level: 1,
        maxHP: 100,
        hP: 40,
        inventory: []string {"3 Potion"},
    }
    displayInfo(Player)
    accessInventory((Player))
}
