package main

import "fmt"

type Character struct {
	Name string
	Class string
	Level int
	MaxHP int 
	CurrentHP int 
	Inventory []string
}

func Init(name string, class string, level int, maxHP int, currentHP int, inventory []string) Character {
player := Character{
	Name: name,
	Class: class,
	Level: level,
	MaxHP: maxHP,
	CurrentHP: currentHP,
	Inventory: inventory,
}
return player
}

func main() {
p1 := Init("Raff", "Elf", 1, 100, 40, []string{"Potion", "Potion", "Potion"})

fmt.Println("Player 1's name:", p1.Name)
fmt.Println("Player 1's class:", p1.Class)
fmt.Println("Player 1's level:", p1.Level)
fmt.Println("Player 1's maxHP:", p1.MaxHP)
fmt.Println("Player 1's currentHP:", p1.CurrentHP)
fmt.Println("Player 1's inventory:", p1.Inventory)
}
