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

type Equipment struct {
    Name string
    Recipe map[string]int
    Cost int            
    CraftMsg string         
}

var equipmentList = map[string]Equipment{
    "Adventurer's Hat": {
        Name: "Adventurer's Hat",
        Recipe: map[string]int{"Crow Feather": 1, "Boar Leather": 1},
        Cost: 5,
        CraftMsg: "You crafted an Adventurer's Hat!",
    },
    "Adventurer's Tunic": {
        Name: "Adventurer's Tunic",
        Recipe: map[string]int{"Wolf Fur": 2, "Troll Skin": 1},
        Cost: 10,
        CraftMsg: "You crafted an Adventurer's Tunic!",
    },
    "Adventurer's Boots": {
        Name: "Adventurer's Boots",
        Recipe: map[string]int{"Wolf Fur": 1, "Boar Leather": 1},
        Cost: 7,
        CraftMsg: "You crafted Adventurer's Boots!",
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

    for {
        fmt.Println("Menu:")
        fmt.Println("1. Display Character Information")
        fmt.Println("2. Access Inventory")
        fmt.Println("3. Merchant")
        fmt.Println("4. Forge")
        fmt.Println("5. Check for Death")
        fmt.Println("6. Use Rat poison")
        fmt.Println("7. Learn Hunter Skill")
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
            checkDead(&c1)
        case 6:
            if containsItem(&c1, "Rat poison") {
                removeInventory(&c1, "Rat poison")
                ratPoison(&c1)
            } else {
                fmt.Println("You don't have any Rat poison in your inventory.")
            }
        case 7:
            if containsItem(&c1, "Hunter Skill") {
                spellBook(&c1, "Hunting")
            } else {
                fmt.Println("You need the Hunter Skill to learn hunting.")
            }
        case 8:
            fmt.Println("Exiting the program.")
            return
        default:
            fmt.Println("Invalid choice.")
        }
    }
case "No", "no":
    fmt.Println("You're a little pussy")
default:
    fmt.Println("You don't know how to read maruta ?")
    return
}
}
