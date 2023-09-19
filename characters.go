package main

import "fmt"

type character struct {
	Name string
	Class string 
	Level int
	Maxlevel int
	MaxHP int 
	CurrentHP int
	Inventory []int 
}
func main() {

	player1 := character {
		Name : "Raff",
		Class : "warrior",
		Level : 1,
		Maxlevel : 5,
		CurrentHP : 100,
		MaxHP : 500,
	}
fmt.Println("character name :", player1.name)
fmt.Println("character class :", player1.class)
fmt.Println("character level :", player1.level)
fmt.Println("character max level :", player1.maxlevel)
fmt.Println("character HP :", player1.currentHP)
fmt.Println("character max HP :", player1.maxHP)
}
type Class int
const(
	warrior Class = iota
	archery Class = iota
	healer Class = iota
	
)