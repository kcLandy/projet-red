package main

import (
	"fmt"
	"math/rand"
	"time"
)

func combat(fighter *Character, rats *Enemys) {
	fmt.Println("Heat the ovens !")

	for Character.CurrentHealthPoints > 0 && Enemys.HP > 0 {
		fighterAttack(fighter, rats)
		if rats.HP <= 0 {
			fmt.Println("Ready to bake !")
			break
		}
		 ratsAttacks(fighter, rats)
		 if fighter.CurrentHealthPoints <= 0 {
			fmt.Println("")
		 }

	}
}