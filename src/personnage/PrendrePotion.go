package personnage

import "fmt"

// Fonction pour utiliser une potion
func TakePot(player *Player) {
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
