package personnage

import "fmt"

func MarchantMenu(player *Player) {
	for {
		fmt.Println(cyan + "\n=================== Marchand ==================" + reset)
		fmt.Println(yellow + "1" + reset + " - Acheter une Potion de vie (3 pièces d'or)")
		fmt.Println(yellow + "2" + reset + " - Acheter une Potion de poison (6 pièces d'or)")
		fmt.Println(yellow + "3" + reset + " - Acheter un livre de Compétence (25 pièces d'or)")
		fmt.Println(yellow + "4" + reset + " - Acheter une Fourrure de Loup (4 pièces d'or)")
		fmt.Println(yellow + "5" + reset + " - Acheter une Peau de Troll (7 pièces d'or)")
		fmt.Println(yellow + "6" + reset + " - Acheter un Cuir de Sanglier (3 pièces d'or)")
		fmt.Println(yellow + "7" + reset + " - Acheter une épée (100 pièces d'or)")
		fmt.Println(yellow + "8" + reset + " - Acheter une Arc (100 pièces d'or)")
		fmt.Println(yellow + "9" + reset + " - Acheter une Baguette Magique (100 pièces d'or)")
		fmt.Println(yellow + "0" + reset + " - Retour")
		fmt.Printf("\n Vous Avez %d Or\n", player.Gold)
		fmt.Println(cyan + "================================================" + reset)

		var choice int
		fmt.Print("Choix : ")
		fmt.Scan(&choice)

		// Calculer le total d'objets dans l'inventaire
		totalItems := player.Inventaire.Wood + player.Inventaire.Stone + player.Inventaire.Leaf +
			player.Inventaire.Fourrure + player.Inventaire.Peau_Troll + player.Inventaire.CuirSanglier +
			player.Inventaire.PlumeCorbeau + player.Inventaire.Potions + player.Inventaire.Sword +
			player.Inventaire.Bow + player.Inventaire.MagicStaff

		switch choice {
		case 1:
			if player.Gold >= 3 {
				if totalItems < 50 {
					addInventory(player, "potion", 1)
					fmt.Println(green + "Vous avez acheté une Potion de vie." + reset)
					player.Gold -= 3
				} else {
					fmt.Println(red + "Inventaire plein. Vous ne pouvez pas acheter cet objet." + reset)
				}
			} else {
				fmt.Println(red + "Vous n'avez pas assez de pièces d'or." + reset)
			}
		case 2:
			if player.Gold >= 6 {
				if totalItems < 50 {
					addInventory(player, "potionPoison", 1)
					fmt.Println(green + "Vous avez acheté une Potion de poison." + reset)
					player.Gold -= 6
				} else {
					fmt.Println(red + "Inventaire plein. Vous ne pouvez pas acheter cet objet." + reset)
				}
			} else {
				fmt.Println(red + "Vous n'avez pas assez de pièces d'or." + reset)
			}
		case 3:
			if player.Gold >= 25 {
				if totalItems < 50 {
					player.Inventaire.SpellBookCount++
					fmt.Println(green + "Vous avez acheté un Livre de compétence : Boule de Feu." + reset)
					player.Gold -= 25
				} else {
					fmt.Println(red + "Inventaire plein. Vous ne pouvez pas acheter cet objet." + reset)
				}
			} else {
				fmt.Println(red + "Vous n'avez pas assez de pièces d'or." + reset)
			}
		case 4:
			if player.Gold >= 4 {
				if totalItems < 50 {
					player.Inventaire.Fourrure++
					fmt.Println(green + "Vous avez acheté une Fourrure de Loup." + reset)
					player.Gold -= 4
				} else {
					fmt.Println(red + "Inventaire plein. Vous ne pouvez pas acheter cet objet." + reset)
				}
			} else {
				fmt.Println(red + "Vous n'avez pas assez de pièces d'or." + reset)
			}
		case 5:
			if player.Gold >= 7 {
				if totalItems < 50 {
					player.Inventaire.Peau_Troll++
					fmt.Println(green + "Vous avez acheté une Peau de Troll." + reset)
					player.Gold -= 7
				} else {
					fmt.Println(red + "Inventaire plein. Vous ne pouvez pas acheter cet objet." + reset)
				}
			} else {
				fmt.Println(red + "Vous n'avez pas assez de pièces d'or." + reset)
			}
		case 6:
			if player.Gold >= 3 {
				if totalItems < 50 {
					player.Inventaire.CuirSanglier++
					fmt.Println(green + "Vous avez acheté un Cuir de Sanglier." + reset)
					player.Gold -= 3
				} else {
					fmt.Println(red + "Inventaire plein. Vous ne pouvez pas acheter cet objet." + reset)
				}
			} else {
				fmt.Println(red + "Vous n'avez pas assez de pièces d'or." + reset)
			}
		case 7:
			if player.Gold >= 100 {
				if totalItems < 50 {
					player.Inventaire.Sword++
					fmt.Println(green + "Vous avez acheté une épée." + reset)
					player.Gold -= 100
				} else {
					fmt.Println(red + "Inventaire plein. Vous ne pouvez pas acheter cet objet." + reset)
				}
			} else {
				fmt.Println(red + "Vous n'avez pas assez de pièces d'or." + reset)
			}
		case 8:
			if player.Gold >= 100 {
				if totalItems < 50 {
					player.Inventaire.Bow++
					fmt.Println(green + "Vous avez acheté un arc." + reset)
					player.Gold -= 100
				} else {
					fmt.Println(red + "Inventaire plein. Vous ne pouvez pas acheter cet objet." + reset)
				}
			} else {
				fmt.Println(red + "Vous n'avez pas assez de pièces d'or." + reset)
			}
		case 9:
			if player.Gold >= 100 {
				if totalItems < 50 {
					player.Inventaire.MagicStaff++
					fmt.Println(green + "Vous avez acheté une baguette magique." + reset)
					player.Gold -= 100
				} else {
					fmt.Println(red + "Inventaire plein. Vous ne pouvez pas acheter cet objet." + reset)
				}
			} else {
				fmt.Println(red + "Vous n'avez pas assez de pièces d'or." + reset)
			}
		case 0:
			// Retourne au menu précédent
			return
		default:
			fmt.Println(red + "Choix invalide." + reset)
		}
	}
}

func addInventory(player *Player, item string, quantity int) {
	switch item {
	case "potion":
		player.Inventaire.Potions += quantity
	case "potionPoison":
		player.Inventaire.Potion_Poison += quantity
	case "sword":
		player.Inventaire.Sword += quantity
	case "bow":
		player.Inventaire.Bow += quantity
	case "magicstaff":
		player.Inventaire.MagicStaff += quantity
	}
}
