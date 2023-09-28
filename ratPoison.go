package main

import (
	"fmt"
	"time"
)

func ratPoison(rat *RatEnemy) {
    fmt.Println("You poisoned the rat! He takes 10 points of poison damage per second for 3 seconds.")
    
    for i := 0; i < 3; i++ {
        rat.HP -= 10
        if rat.HP < 0 {
            rat.HP = 0
        }
        fmt.Printf("Current Health Points: %d \n", rat.HP)
        time.Sleep(1 * time.Second)
    }
    
    fmt.Println("The poison has worn off.")
}