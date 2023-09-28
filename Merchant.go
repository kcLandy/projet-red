package main

import "fmt"

func Merchant(c1 *Character) {
	charac1 := c1

	fmt.Println("Merchant Menu:")
            fmt.Println("1. Buy Health Potion (3 Gold)")
            fmt.Println("2. Buy Rat Poison (6 Gold)")
            fmt.Println("3. Buy Rat Hunter Skill (25 Gold)")
            fmt.Println("4. Buy Rat Fur (4 Gold)")
            fmt.Println("5. Buy Rat Tail (7 Gold)")
            fmt.Println("6. Buy Rat Leather (3 Gold)")
            fmt.Println("8. Return to Previous Menu")

            var merchantChoice int
            fmt.Print("Enter your choice: ")
            fmt.Scan(&merchantChoice)

            switch merchantChoice {
            case 1:
                if c1.Gold >= 3 {
                    addInventory(charac1, "Potion")
                    c1.Gold -= 3
                } else {
                    fmt.Println("Not enough gold to buy a Health Potion.")
                }
            case 2:
                if c1.Gold >= 6 {
                    addInventory(charac1, "Rat poison")
                    c1.Gold -= 6
                } else {
                    fmt.Println("Not enough gold to buy a Rat poison.")
                }
            case 3:
                if c1.Gold >= 25 {
                    addInventory(charac1, "Rat Hunter Skil")
                    ratHunter(charac1, "Rat Hunting")
                    c1.Gold -= 25
                } else {
                    fmt.Println("Not enough gold to buy a Rat Hunter Skill.")
                }
            case 4:
                if c1.Gold >= 4 {
                    addInventory(charac1, "Rat Fur")
                    c1.Gold -= 4
                } else {
                    fmt.Println("Not enough gold to buy Rat Fur.")
                }
            case 5:
                if c1.Gold >= 7 {
                    addInventory(charac1, "Rat Tail")
                    c1.Gold -= 7
                } else {
                    fmt.Println("Not enough gold to buy Rat Tail.")
                }
            case 6:
                if c1.Gold >= 3 {
                    addInventory(charac1, "Rat Leather")
                    c1.Gold -= 3
                } else {
                    fmt.Println("Not enough gold to buy Rat Leather.")
                }
            case 7:
            default:
                fmt.Println("Invalid choice.")
            }
}
