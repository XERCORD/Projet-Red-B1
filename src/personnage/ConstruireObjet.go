package personnage

import "fmt"

// Fonction pour construire des objets
func CraftItems(player *Player) {
	for {
		fmt.Println(cyan + "\n================================================" + reset)
		fmt.Println("   Que voulez-vous fabriquer ?")
		fmt.Println(cyan + "================================================" + reset)
		fmt.Println(yellow + "1" + reset + " - Épée (5 bois, 5 pierre)")
		fmt.Println(yellow + "2" + reset + " - Arc (5 bois, 5 feuilles)")
		fmt.Println(yellow + "3" + reset + " - Bâton magique (5 bois, 5 feuilles, 5 pierre)")
		fmt.Println(yellow + "4" + reset + " - Chapeau de l'aventurier (1 Plume de Corbeau, 1 Cuir de Sanglier)")
		fmt.Println(yellow + "5" + reset + " - Tunique de l'aventurier (2 Fourrures de Loup, 1 Peau de troll)")
		fmt.Println(yellow + "6" + reset + " - Bottes de l'aventurier (1 Fourrure de Loup, 1 Cuir de Sanglier)")
		fmt.Println(yellow + "0" + reset + " - Retour")
		fmt.Println(cyan + "================================================" + reset)

		var craftChoice int
		fmt.Print("Choix : ")
		fmt.Scan(&craftChoice)

		switch craftChoice {
		case 0:
			// Revenir au menu principal
			return
		case 1:
			if player.Inventaire.Wood >= 5 && player.Inventaire.Stone >= 5 && player.Gold >= 5 {
				player.Gold -= 5
				player.Inventaire.Wood -= 5
				player.Inventaire.Stone -= 5
				player.Inventaire.Sword++
				fmt.Println(green + "Vous avez fabriqué une épée." + reset)
			} else {
				fmt.Println(red + "Vous n'avez pas assez de ressources." + reset)
			}
		case 2:
			if player.Inventaire.Wood >= 5 && player.Inventaire.Leaf >= 5 && player.Gold >= 5 {
				player.Gold -= 5
				player.Inventaire.Wood -= 5
				player.Inventaire.Leaf -= 5
				player.Inventaire.Bow++
				fmt.Println(green + "Vous avez fabriqué un arc." + reset)
			} else {
				fmt.Println(red + "Vous n'avez pas assez de ressources." + reset)
			}
		case 3:
			if player.Inventaire.Wood >= 5 && player.Inventaire.Leaf >= 5 && player.Inventaire.Stone >= 5 && player.Gold >= 5 {
				player.Gold -= 5
				player.Inventaire.Wood -= 5
				player.Inventaire.Leaf -= 5
				player.Inventaire.Stone -= 5
				player.Inventaire.MagicStaff++
				fmt.Println(green + "Vous avez fabriqué un bâton magique." + reset)
			} else {
				fmt.Println(red + "Vous n'avez pas assez de ressources." + reset)
			}
		case 4:
			if player.Gold >= 5 && player.Inventaire.PlumeCorbeau >= 1 && player.Inventaire.CuirSanglier >= 1 {
				player.Gold -= 5
				player.Inventaire.PlumeCorbeau -= 1
				player.Inventaire.CuirSanglier -= 1
				player.equip.Head = stuff{"Coiffe Bouftou", 10}
				player.HealthMax += player.equip.Head.pvb
				fmt.Println(green + "Vous avez fabriqué une Coiffe Bouftou." + reset)
			} else {
				fmt.Println(red + "Vous n'avez pas assez de ressources." + reset)
			}
		case 5:
			if player.Gold >= 5 && player.Inventaire.Fourrure >= 2 && player.Inventaire.Peau_Troll >= 1 {
				player.Gold -= 5
				player.Inventaire.Fourrure -= 2
				player.Inventaire.Peau_Troll -= 1
				player.equip.Head = stuff{"Cape Bouftou", 10}
				player.HealthMax += player.equip.Plastron.pvb
				fmt.Println(green + "Vous avez fabriqué une Cape Bouftou." + reset)
			} else {
				fmt.Println(red + "Vous n'avez pas assez de ressources." + reset)
			}
		case 6:
			if player.Gold >= 5 && player.Inventaire.Fourrure >= 1 && player.Inventaire.CuirSanglier >= 1 {
				player.Gold -= 5
				player.Inventaire.Fourrure -= 1
				player.Inventaire.CuirSanglier -= 1
				player.equip.Head = stuff{"Botte de kheir", 10}
				player.HealthMax += player.equip.Feet.pvb
				fmt.Println(green + "Vous avez fabriqué une paire de Boufbotte." + reset)
			} else {
				fmt.Println(red + "Vous n'avez pas assez de ressources." + reset)
			}
		default:
			fmt.Println(red + "Choix invalide." + reset)
		}
	}
}
