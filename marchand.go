func addInventory(player *Character, item string) {
	player.Inventory = append(player.Inventory, item)
}

func removeInventory(player *Character, item string) {
	for i, currentItem := range player.Inventory {
		if currentItem == item {
			player.Inventory = append(player.Inventory[:i], player.Inventory[i+1]...)
			fmt.Printf("%s has been add to your inventory.\n", item)
			return
		}
	}
	var merchant = []string{"free health potion"}

	func buyFromMerchant(player *Character, item string) {
		for i, merchantItem := range merchant {
			if merchantItem == item {
				addInventory(player, item)

				merchant = append(merchant[:i], merchant[i+1]...)

				fmt.Printf("You bought the %s from the Merchant.\n", item)
				return
			}
		}
		fmt.Printf("%s in no more avalaible at the merchant store.\n", item)
	}

	func main() {
		or {
			fmt.Println("\n Menu :")
			fmt.Println("1. Show off player's informations")
			fmt.Println("2. Access to inventory")
			fmt.Println("3. Merchant")
			fmt.Println("4. Quitter")
	
			var choice int 
			fmt.Print("Choose an option")
			fmt.Scan(&choice)
	
			switch choice {
			case 1:
				displayInfo(p1)
			case 2:
				accessInventory(p1)
			case 3:
				fmt.Println("\n Itemps available in your Merchant store:")
				for i, item := range merchant {
					fmt.Printf("%d. %s\n", i+1, item)
				}
				fmt.Println("0. Back to the main menu")

				var merchantChoice int 
				fmt.Print("Please chose an item: ")
				fmt.Scan(&merchantChoice)

				if merchantChoice == 0 {
					break
				}
				if merchantchoice > 0 && merchantChoice <= len(merchant) {
					buyFromMerchant(&p1, merchant[marchandChoice-1])
				} else {
					fmt.Printl("\n Option unavailable, please chose a valid option.")
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
			case 4:
				fmt.Println("\n See you soon !")
				return
			default:
				fmt.Println("\n Option unavailable, please choose a valid option.")
			}
		}
	}
}
