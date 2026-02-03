package personnage

import (
	"fmt"
	"math/rand"
)

// Fonction pour récolter des ressources
func GatherResources(player *Player) {
	for {
		fmt.Println(cyan + "\n================================================" + reset)
		fmt.Println("   Que voulez-vous récolter ?")
		fmt.Println(cyan + "================================================" + reset)
		fmt.Println(yellow + "1" + reset + " - Bois")
		fmt.Println(yellow + "2" + reset + " - Pierre")
		fmt.Println(yellow + "3" + reset + " - Feuilles")
		fmt.Println(yellow + "4" + reset + " - Fourrure de Loup")
		fmt.Println(yellow + "5" + reset + " - Peau de Troll")
		fmt.Println(yellow + "6" + reset + " - Cuir de Sanglier")
		fmt.Println(yellow + "7" + reset + " - Plume de Corbeau")
		fmt.Println(yellow + "0" + reset + " - Retour")
		fmt.Println(cyan + "================================================" + reset)

		var resourceChoice int
		fmt.Print("Choix : ")
		fmt.Scan(&resourceChoice)

		// Calculer le total d'objets dans l'inventaire
		totalItems := player.Inventaire.Wood + player.Inventaire.Stone + player.Inventaire.Leaf +
			player.Inventaire.Fourrure + player.Inventaire.Peau_Troll + player.Inventaire.CuirSanglier +
			player.Inventaire.PlumeCorbeau + player.Inventaire.Potions + player.Inventaire.Sword +
			player.Inventaire.Bow + player.Inventaire.MagicStaff

		// Si le total dépasse 50, on empêche de récolter plus de ressources
		if totalItems >= 50 {
			fmt.Println(red + "Inventaire plein. Vous ne pouvez plus récolter d'autres ressources." + reset)
			return
		}

		// Calculer l'espace restant
		spaceRemaining := 50 - totalItems

		switch resourceChoice {
		case 0:
			// Revenir au menu principal
			return
		case 1:
			wood := rand.Intn(10) + 1
			if wood > spaceRemaining {
				wood = spaceRemaining // Ajuster la quantité à récolter
			}
			player.Inventaire.Wood += wood
			fmt.Printf(green+"Vous avez récolté %d unités de bois. Total de bois : %d\n"+reset, wood, player.Inventaire.Wood)
		case 2:
			stone := rand.Intn(10) + 1
			if stone > spaceRemaining {
				stone = spaceRemaining // Ajuster la quantité à récolter
			}
			player.Inventaire.Stone += stone
			fmt.Printf(green+"Vous avez récolté %d unités de pierre. Total de pierre : %d\n"+reset, stone, player.Inventaire.Stone)
		case 3:
			leaf := rand.Intn(10) + 1
			if leaf > spaceRemaining {
				leaf = spaceRemaining // Ajuster la quantité à récolter
			}
			player.Inventaire.Leaf += leaf
			fmt.Printf(green+"Vous avez récolté %d feuilles. Total de feuilles : %d\n"+reset, leaf, player.Inventaire.Leaf)
		case 4:
			fourrure := rand.Intn(10) + 1
			if fourrure > spaceRemaining {
				fourrure = spaceRemaining // Ajuster la quantité à récolter
			}
			player.Inventaire.Fourrure += fourrure
			fmt.Printf(green+"Vous avez récolté %d unités de fourrure de loup. Total de fourrure de loup : %d\n"+reset, fourrure, player.Inventaire.Fourrure)
		case 5:
			peauTroll := rand.Intn(10) + 1
			if peauTroll > spaceRemaining {
				peauTroll = spaceRemaining // Ajuster la quantité à récolter
			}
			player.Inventaire.Peau_Troll += peauTroll
			fmt.Printf(green+"Vous avez récolté %d unités de peau de troll. Total de peau de troll : %d\n"+reset, peauTroll, player.Inventaire.Peau_Troll)
		case 6:
			cuirSanglier := rand.Intn(10) + 1
			if cuirSanglier > spaceRemaining {
				cuirSanglier = spaceRemaining // Ajuster la quantité à récolter
			}
			player.Inventaire.CuirSanglier += cuirSanglier
			fmt.Printf(green+"Vous avez récolté %d unités de cuir de sanglier. Total de cuir de sanglier : %d\n"+reset, cuirSanglier, player.Inventaire.CuirSanglier)
		case 7:
			plumeCorbeau := rand.Intn(10) + 1
			if plumeCorbeau > spaceRemaining {
				plumeCorbeau = spaceRemaining // Ajuster la quantité à récolter
			}
			player.Inventaire.PlumeCorbeau += plumeCorbeau
			fmt.Printf(green+"Vous avez récolté %d unités de plume de corbeau. Total de plume de corbeau : %d\n"+reset, plumeCorbeau, player.Inventaire.PlumeCorbeau)
		default:
			fmt.Println(red + "Choix invalide." + reset)
		}
	}
}
