package personnage

import "fmt"

// Fonction pour afficher l'information du personnage
func DisplayInfo(player Player) {
	fmt.Printf("Pseudo : %s\nSexe : %s\nClasse : %s\nNiveau : %d\nVie actuelle : %d\nVie Max : %d\nPièce d'or : %d\n", player.Pseudo, player.Sex, player.Class, player.Level, player.Health, player.HealthMax, player.Gold)
}
