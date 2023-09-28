package main

import (
	"fmt"
	"math/rand"
	"time"
)

	func combat(character Character, rat RatEnemy) {
		fmt.Printf("A wild rat appears!\nLevel: %d\n", rat.Level)

		rand.Seed(time.Now().UnixNano())

		for character.HP > 0 && rat.HP > 0 {

			rat.HP -= character.Attack
	
			fmt.Printf("%s attacks the rat and deals %d damage!\n", character.Name, character.Attack)
	
			if rat.HP <= 0 {
				fmt.Println("The rat has been defeated!")
				break
			}
	
			character.HP -= rat.Attack
	
			fmt.Printf("The rat counterattacks and deals %d damage to %s!\n", rat.Attack, character.Name)
	
			if character.HP <= 0 {
				checkDead(&character)
				break
			}
	
			fmt.Printf("Current HP - %s: %d, Rat: %d\n", character.Name, character.HP, rat.HP)
		}
	}
	