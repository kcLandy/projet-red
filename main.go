package main

import (
    "fmt"
    "math/rand"
    "time"
)

type Character struct {
    Name string
    Class string
    Level int
    MaxHP int
    HP int
    Inventory []string
    Gold int
    Skills []string
    Attack int
}

type RatEnemy struct {
	Level  int
	HP     int
	Attack int
}

type Equipment struct {
    Name string
    Recipe map[string]int
    Cost int            
    CraftMsg string         
}

var equipmentList = map[string]Equipment{
    "Hunter's Hat": {
        Name: "Hunter's Hat",
        Recipe: map[string]int{"Rat Tail": 1, "Rat Leather": 2},
        Cost: 10,
        CraftMsg: "You crafted an Hunter's Hat!",
    },
    "Hunter's Tunic": {
        Name: "Hunter's Tunic",
        Recipe: map[string]int{"Rat Fur": 5, "Rat Tail": 4, "Rat Leather": 3},
        Cost: 25,
        CraftMsg: "You crafted an Hunter's Tunic!",
    },
    "Hunter's Boots": {
        Name: "Hunter's Boots",
        Recipe: map[string]int{"Rat Tail": 2, "Rat Fur": 2},
        Cost: 15,
        CraftMsg: "You crafted Hunter's Boots!",
    },
}

func main() {
    fmt.Println("Start the game ?")
    fmt.Println("Yes")
    fmt.Println("No")

    var choice string
    fmt.Scan(&choice)

    switch choice {
    case "Yes", "yes":
        fmt.Println("You're a rats killer")
    
    rand.Seed(time.Now().UnixNano()) // sert à générer un nombre aléatoire

    c1 := charCreation()
    r1 := RatEnemy{}

    for {
        fmt.Println("Menu:")
        fmt.Println("1. Display Character Information")
        fmt.Println("2. Access Inventory")
        fmt.Println("3. Merchant")
        fmt.Println("4. Forge")
        fmt.Println("5. Use Rat poison")
        fmt.Println("6. Learn Hunter Skill")
        fmt.Println("7. Burn some rats")
        fmt.Println("8. Exit")

        var choice int
        fmt.Scan(&choice)

        switch choice {
        case 1:
            displayInfoChar(c1)
        case 2:
            accessInventory(&c1)
        case 3:
            Merchant(&c1)
        case 4:
            forgeMenu(&c1)
        case 5:
            if containsItem(&c1, "Rat poison") {
                removeInventory(&c1, "Rat poison")
                ratPoison(&r1)
            } else {
                fmt.Println("You don't have any Rat poison in your inventory.")
            }
        case 6:
            if containsItem(&c1, "Hunter Skill") {
                ratHunter(&c1, "Hunting")
            } else {
                fmt.Println("You need the Hunter Skill to learn hunting.")
            }
        case 7:
            r1 = ratEnemy()
            combat(c1,r1)
        case 8:
            fmt.Println("Exiting the program.")
            return
        default:
            fmt.Println("Invalid choice.")
        }
    }
case "No", "no":
    fmt.Println("You're not a real pizza maker")
default:
    fmt.Println("You don't know how to read maruta ?")
    return
}
}
