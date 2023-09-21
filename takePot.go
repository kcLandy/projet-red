func takePot(player *Character) {
	
	if len(player.Inventory) == 0 {
		fmt.Println("The inventory is empty. You can't use any potions.")
		return
	}
	for i, item := range player.Inventory {
		if item == "Potion" {
			fmt.Println("You used one potion !")

			player.CurrentHP +=50

			if player.CurrentHP > player.MaxHP {
			} else {
				player.CurrentHP = player.MaxHP
			}
			player.Inventory = append(player.Inventory[:i], player.Inventory[i+1]...)

			fmt.Println("CurrentHP : %d / %d\n", player.CurrentHP, player.MaxHP)
			return
			}
		}
		fmt.Println("You don't have any potion in inventory.")
	}