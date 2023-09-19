package main

import "fmt"

type character struct {
	name string
	class string 
	level int
	maxlevel int
	maxHP int 
	currentHP int
	inventory []int 
}
func main() {

	player1 := character {
		name : "Raff",
		class : "warrior",
		level : 1,
		maxlevel : 5,
		currentHP : 100,
		maxHP : 500,
	}
fmt.Println("character name :", player1.name)
fmt.Println("character class :", player1.class)
fmt.Println("character level :", player1.level)
fmt.Println("character HP :", player1.currentHP)
}
