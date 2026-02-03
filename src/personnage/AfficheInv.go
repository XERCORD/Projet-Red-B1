package personnage

import "fmt"

func totalItems(player *Player) int {
	// Comptabilise tous les objets dans l'inventaire
	total := player.Inventaire.Bow +
		player.Inventaire.MagicStaff +
		player.Inventaire.Sword +
		player.Inventaire.Potion_Poison +
		player.Inventaire.ChapAven +
		player.Inventaire.TunAven +
		player.Inventaire.Wood +
		player.Inventaire.Stone +
		player.Inventaire.Leaf +
		player.Inventaire.Fourrure +
		player.Inventaire.CuirSanglier +
		player.Inventaire.PlumeCorbeau +
		player.Inventaire.Peau_Troll +
		player.Inventaire.Potions +
		player.Inventaire.SpellBookCount
	return total
}

// Fonction pour afficher l'inventaire
func AccessInventory(player *Player) {
	fmt.Println(cyan + "\n|=================== Inventaire ==================|" + reset)

	// Information du joueur
	fmt.Printf("Pseudo : %s | Sexe : %s | Classe : %s Vie Max : %d | Vie Actuelle : %d | Niveau : %d | Pièces d'or : %d\n ", player.Pseudo, player.Sex, player.Class, player.HealthMax, player.Health, player.Level, player.Gold)

	// Objets
	fmt.Println(cyan + "\n[Equipement]" + reset)

	if player.Inventaire.Bow > 0 {
		fmt.Printf(green+"- Arc (%d)\n"+reset, player.Inventaire.Bow)
	}
	if player.Inventaire.MagicStaff > 0 {
		fmt.Printf(green+"- Bâton magique (%d)\n"+reset, player.Inventaire.MagicStaff)
	}

	fmt.Println(cyan + "\n[Objets]" + reset)
	if player.Inventaire.Sword > 0 {
		fmt.Printf(green+"- Épée (%d)\n"+reset, player.Inventaire.Sword)
	}
	if player.Inventaire.Potion_Poison > 0 {
		fmt.Printf(green+"- Potions de Poison (%d)\n"+reset, player.Inventaire.Potion_Poison)
	}
	if player.Inventaire.ChapAven > 0 {
		fmt.Printf(green+"- Chapeau de l'aventurier (%d)\n"+reset, player.Inventaire.ChapAven)
	}
	if player.Inventaire.TunAven > 0 {
		fmt.Printf(green+"- Tunique de l'aventurier (%d)\n"+reset, player.Inventaire.TunAven)
	}
	if totalItems(player) == 0 {
		fmt.Println(red + "Aucun objet." + reset)
	}

	fmt.Println(cyan + "\n[Compétences]" + reset)
	for _, skill := range player.Skills {
		fmt.Println(green + "- " + skill + reset)
	}

	// Ressources
	fmt.Println(cyan + "\n[Ressources]" + reset)
	fmt.Printf("- Bois : %d\n", player.Inventaire.Wood)
	fmt.Printf("- Pierre : %d\n", player.Inventaire.Stone)
	fmt.Printf("- Feuilles : %d\n", player.Inventaire.Leaf)
	fmt.Printf("- Fourrure de loup : %d\n", player.Inventaire.Fourrure)
	fmt.Printf("- Cuir de sanglier : %d\n", player.Inventaire.CuirSanglier)
	fmt.Printf("- Plume de corbeau : %d\n", player.Inventaire.PlumeCorbeau)
	fmt.Printf("- Peau de troll : %d\n", player.Inventaire.Peau_Troll)

	// Consommables
	fmt.Println(cyan + "\n[Consommables]" + reset)
	fmt.Printf("- Potions : %d\n", player.Inventaire.Potions)

	// Affichage de la limite d'inventaire
	fmt.Printf("\n"+cyan+"Capacité d'inventaire : %d/50\n"+reset, totalItems(player))

	fmt.Println("\n" + cyan + "================================================" + reset)
	fmt.Println("   Que voulez-vous faire ?")
	fmt.Println("================================================")
	fmt.Println(yellow + "1" + reset + " - Prendre une potion")
	fmt.Println(yellow + "2" + reset + " - Utiliser le Livre de compétence : Boule de Feu")
	fmt.Println(yellow + "0" + reset + " - Retour")

	var choice int
	fmt.Print("Choix : ")
	fmt.Scan(&choice)

	switch choice {
	case 1:
		takePot(player)
	case 2:
		if player.Inventaire.SpellBookCount > 0 {
			spellBook(player)
			player.Inventaire.SpellBookCount--
		} else {
			fmt.Println(red + "Vous n'avez pas de Livre de compétence : Boule de Feu." + reset)
		}
	case 0:
		return
	default:
		fmt.Println(red + "Choix invalide." + reset)
	}
}

func takePot(player *Player) {
	if player.Health == player.HealthMax {
		fmt.Println(yellow + "Votre vie est déjà au maximum !" + reset)
		return
	}

	if player.Inventaire.Potions > 0 {
		healAmount := 50
		actualHeal := player.HealthMax - player.Health
		if healAmount > actualHeal {
			healAmount = actualHeal
		}
		player.Health += healAmount
		player.Inventaire.Potions--
		fmt.Printf(green+"Vous avez utilisé une potion et regagnez %d points de vie.\n"+reset, healAmount)
		fmt.Printf("Votre vie actuelle est de %d / %d\n", player.Health, player.HealthMax)
	} else {
		fmt.Println(red + "Vous n'avez plus de potions !" + reset)
	}
}
