package 

import "fmt"

type Character struct {
	Name string
	Class string
	Level int
	MaxHP int 
	CurrentHP int 
	Inventory []string
}

func Init(name string, class string, level int, maxHP int, currentHP int, inventory []string) Character {
player := Character{
	Name: name,
	Class: class,
	Level: level,
	MaxHP: maxHP,
	CurrentHP: currentHP,
	Inventory: inventory,
}
return player
}

func displayInfo(player Character) {
	fmt.Println("\n show off player's informartions :")
	fmt.Println("\n player's name:", player.Name)
	fmt.Println("\n player's class:", player.Class)
	fmt.Println("\n player's level:", player.Level)
	fmt.Println("\n player's maxHP:", player.MaxHP)
	fmt.Println("\n player's currentHP:", player.CurrentHP)
	fmt.Println("\n player's inventory:", player.Inventory)
}

func main() {
	p1 := Init("Raff", "Elf", 1, 100, 40, []string{"Potion", "Potion", "Potion"})

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
			disploInfo(p1)
		case 2:
			fmt.Println("\n Player's inventory:", p1.Inventory)
		case 3:
			fmt.Println("\n See you soon !")
			return
		default:
			fmt.Println("\n Option unavailable, please choose a valid option.")
		}
	}
}
