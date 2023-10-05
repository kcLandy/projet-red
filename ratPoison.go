package main

import (
	"fmt"
	"time"
)

func ratPoison(rat *RatEnemy) {
    fmt.Println("You poisoned the rat! He takes 10 points of poison damage per second for 3 seconds.")
    
    for i := 0; i < 3; i++ {
        rat.HP -= 10
        fmt.Printf("Current Health Points: %d \n", rat.HP)
        time.Sleep(1 * time.Second)
    }
    if rat.HP <= 0 {
        fmt.Println("You killed the rat")
    }
    
    fmt.Println("The poison has worn off.")
}