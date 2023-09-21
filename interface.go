package main

import "fmt"

func main () {
	player := "Raff"
	inventory := []string{"weapon", "shield", "potions"}

	for {
		fmt.Println("\n Menu :")
		fmt.Println("1. Show off player's informations")
		fmt.Println("2. Access to inventory")
		fmt.Println("3. Exit")

		var choice int 
		fmt.Print("Choose an option")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Println("\n player's informations :")
			fmt.Println(player)
		case 2:
			fmt.Println("\n Inventory :")
			for i, item := range inventory {
				fmt.Printf("%d. %s\n", i+1, item)
			}
			var subchoice int 
			fmt.Print("\n Back to main menu")

			switch subchoice {
			case 1:
			case 2:
				fmt.Println("\n See you soon !")
				return
			default:
				fmt.Println("\n Option unavailable, please choose a valid option.")
			}
		case 3:
			fmt.Println("\n See you soon !")
			return
		default:
			fmt.Println("\n Option unavailable, please choose a valid option.")
		}
	}
}