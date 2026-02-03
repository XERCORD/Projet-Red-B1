package personnage

import "fmt"

func spellBook(player *Player) {
	for _, skill := range player.Skills {
		if skill == "Boule de Feu" {
			fmt.Println(red + "Vous avez déjà appris la compétence : Boule de Feu." + reset)
			return
		}
	}
	player.Skills = append(player.Skills, "Boule de Feu")
	fmt.Println(green + "Vous avez appris une nouvelle compétence : Boule de Feu." + reset)
}
