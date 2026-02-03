package combat

import (
	"fmt"
	"math/rand"
	"projetred/personnage"
)

const (
	reset  = "\033[0m"
	red    = "\033[31m"
	green  = "\033[32m"
	yellow = "\033[33m"
	blue   = "\033[34m"
	cyan   = "\033[36m"
)

// Fonction pour combattre des monstres
func Combat(player *personnage.Player) {
	c := 0
	monsterHealth := rand.Intn(50) + 50
	for monsterHealth > 0 || player.Health > 0 {
		// Attaque du joueur
		fmt.Printf(cyan + "\n================================================\n" + reset)
		fmt.Printf("   Vous allez se combattre avec %d de points de vie\n", player.Health)
		fmt.Println(cyan + "================================================" + reset)

		fmt.Println("Veux-tu prendre une potion avant d'aller au combat" + reset)
		fmt.Printf(yellow + "1" + reset + " - Prendre une potion\n")
		fmt.Printf(yellow + "2" + reset + " - Affronter un monstre\n")

		fmt.Println(yellow + "0" + reset + " - Retour")

		var choice int
		fmt.Print("Choix : ")
		fmt.Scan(&choice)
		fmt.Println(cyan + "================================================" + reset)
		switch choice {
		case 1:
			personnage.TakePot(player)
		case 2:
			if monsterHealth > 0 {
				c += 1
				fmt.Printf("Tour : %d\n", c)
				damage := player.Attack
				monsterHealth -= damage
				fmt.Printf(cyan + "\n===================================================\n" + reset)
				fmt.Printf("   Vous allez se combattre avec %d de points de vie\n", player.Health)
				fmt.Println(cyan + "===================================================" + reset)
				fmt.Printf(green+"Vous attaquez et infligez %d de dégâts. Vie du monstre restante : %d\n"+reset, damage, monsterHealth)

				// Attaque du monstre
				if monsterHealth > 0 {
					monsterDamage := rand.Intn(15) + 5
					player.Health -= monsterDamage
					fmt.Printf(red+"Le monstre vous attaque et inflige %d de dégâts.\n"+reset, monsterDamage)
				} else {
					fmt.Println(green + "Le monstre est vaincu !")
					fmt.Printf("Il vous reste %d de points de vie\n", player.Health)
				}
				if player.Health < 0 {
					fmt.Println(red + "Vous êtes mort...\n" + reset)
					player.Health = player.HealthMax / 2
					fmt.Printf(green+"Vous avez été ressuscité avec %d points de vie.\n"+reset, player.Health)
					// Le joueur peut trouver une potion après le combat
				}
				if rand.Float32() < 0.3 { // 30% de chance de trouver une potion
					player.Inventaire.Potions++
					fmt.Println(green + "Vous avez trouvé une potion !\n" + reset)

				}
			} else {
				fmt.Println(green + "Il n'y a plus de monstre.\n" + reset)
			}

		case 0:
			return
		default:
			fmt.Println(red + "Choix invalide." + reset)
		}
	}
}
