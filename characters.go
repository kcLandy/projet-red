package main

import "fmt"

type Classtype string

const (
	Warrior = "Warrior"
	Archer = "Archer"
	Mage = "Mage"
)
type character struct {
	name string
	class Classtype 
	level int
	maxlevel int
	maxHP int 
	currentHP int
}
func main() {
	player1 := character {
		name : "Raff",
		class : "Warrior",
		level : 1,
		maxlevel : 5,
		currentHP : 100,
		maxHP : 500,
	}
fmt.Println("character's name :", player1.name)
fmt.Println("character's class :", player1.class)
fmt.Println("character's level :", player1.level)
fmt.Println("character's max level :", player1.maxlevel)
fmt.Println("character's HP :", player1.currentHP)
fmt.Println("character's max HP :", player1.maxHP)
}
