package 

import "fmt"

type Character struct {
    Name             string
    Class            string
    Level            int
    MaxHealthPoints     int
    CurrentHealthPoints int
    Inventory         []string

func main() {
    
    P1 := Character{
        Name:             "Raffy",
        Class:            "Warrior",
        Level:            1,
        MaxHealthPoints:     500,
        CurrentHealthPoints: 100,
        Inventory:        []string{"Sword", "Shield", "Healing Potion"},
    }

    fmt.Println("Name:", P1.Name)
    fmt.Println("Class:", P1.Class)
    fmt.Println("Level:", P1.Level)
    fmt.Println("Max Hit Points:", P1.MaxHealthPoints)
    fmt.Println("Current Hit Points:", P1.CurrentHealthPoints)
    fmt.Println("Inventory:", P1.Inventory)
}

