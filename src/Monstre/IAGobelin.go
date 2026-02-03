package ennemie

import (
	"fmt"
	"projetred/personnage"
)

func GoblinPattern(goblin Monstre, joueur *personnage.Player, tour int) {
	var degats int
	if tour%3 == 0 {
		// Tous les 3 tours, le gobelin inflige 200% de ses dégâts
		degats = goblin.Attack * 2
		fmt.Printf("%s inflige à %s %d de dégâts\n", goblin.Name, joueur.Pseudo, degats)
	} else {
		// Sinon, le gobelin inflige 100% de ses dégâts
		degats = goblin.Attack
		fmt.Printf("%s inflige à %s %d de dégâts\n", goblin.Name, joueur.Pseudo, degats)
	}

	// Réduire les points de vie du joueur
	joueur.Health -= degats
	if joueur.Health < 0 {
		joueur.Health = 0 // Empêcher les points de vie négatifs
	}

	// Affichage des points de vie restants
	fmt.Printf("%s : %d/%d PV\n", joueur.Pseudo, joueur.Health, joueur.HealthMax)
}
