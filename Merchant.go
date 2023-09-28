package main

import "fmt"

func Merchant(c1 *Character) {
	charac1 := c1

	fmt.Println("Merchant Menu:")
            fmt.Println("1. Buy Potion (3 Gold)")
            fmt.Println("2. Buy Rat poison (6 Gold)")
            fmt.Println("3. Buy Hunter Skill (25 Gold)")
            fmt.Println("4. Buy Wolf Fur (4 Gold)")
            fmt.Println("5. Buy Troll Skin (7 Gold)")
            fmt.Println("6. Buy Boar Leather (3 Gold)")
            fmt.Println("7. Buy Crow Feather (1 Gold)")
            fmt.Println("8. Return to Previous Menu")

            var merchantChoice int // changer var 
            fmt.Print("Enter your choice: ")
            fmt.Scan(&merchantChoice)

            switch merchantChoice {
            case 1:
                if c1.Gold >= 3 {
                    addInventory(charac1, "Potion")
                    c1.Gold -= 3
                } else {
                    fmt.Println("Not enough gold to buy a Potion.")
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
                    addInventory(charac1, "Hunter Skil")
                    spellBook(charac1, "Hunting")
                    c1.Gold -= 25
                } else {
                    fmt.Println("Not enough gold to buy a Hunter Skill.")
                }
            case 4:
                if c1.Gold >= 4 {
                    addInventory(charac1, "Wolf Fur")
                    c1.Gold -= 4
                } else {
                    fmt.Println("Not enough gold to buy Wolf Fur.")
                }
            case 5:
                if c1.Gold >= 7 {
                    addInventory(charac1, "Troll Skin")
                    c1.Gold -= 7
                } else {
                    fmt.Println("Not enough gold to buy Troll Skin.")
                }
            case 6:
                if c1.Gold >= 3 {
                    addInventory(charac1, "Boar Leather")
                    c1.Gold -= 3
                } else {
                    fmt.Println("Not enough gold to buy Boar Leather.")
                }
            case 7:
                if c1.Gold >= 1 {
                    addInventory(charac1, "Crow Feather")
                    c1.Gold -= 1
                } else {
                    fmt.Println("Not enough gold to buy Crow Feather.")
                }
            case 8:
            default:
                fmt.Println("Invalid choice.")
            }
}