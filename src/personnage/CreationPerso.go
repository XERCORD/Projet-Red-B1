package personnage

import (
	"fmt"
	"unicode"
)

// Fonction pour créer le personnage (pseudo, sexe, classe)
func CreateCharacter() Player {
	var player Player
	var pseudo string

	// Choisir le pseudo
	for {
		fmt.Println(cyan + "================================================" + reset)
		fmt.Println("   Entrez votre pseudo (lettres uniquement, entre", minPseudoLength, "et", maxPseudoLength, "caractères) :")
		fmt.Println(cyan + "================================================" + reset)
		fmt.Print("Pseudo : ")
		fmt.Scan(&pseudo)

		// Vérification de la longueur du pseudo
		if len(pseudo) < minPseudoLength || len(pseudo) > maxPseudoLength {
			fmt.Println(red+"Erreur : Le pseudo doit contenir entre", minPseudoLength, "et", maxPseudoLength, "caractères."+reset)
			continue
		}
		// Vérification des caractères (lettres uniquement)
		if isValidPseudo(pseudo) {
			player.Pseudo = pseudo
			break
		} else {
			fmt.Println(red + "Erreur : Le pseudo ne doit contenir que des lettres." + reset)
		}
	}
	// Choisir le sexe
	fmt.Println(cyan + "\n================================================" + reset)
	fmt.Println("   Choisissez votre sexe :")
	fmt.Println(cyan + "================================================" + reset)
	fmt.Println(yellow + "1" + reset + " - Homme")
	fmt.Println(yellow + "2" + reset + " - Femme")
	fmt.Println(cyan + "================================================" + reset)

	var sexChoice int
	for {
		fmt.Print("Choix : ")
		fmt.Scan(&sexChoice)
		if sexChoice == 1 || sexChoice == 2 {
			break
		} else {
			fmt.Println(red + "Choix invalide. Veuillez entrer 1 ou 2." + reset)
		}
	}

	switch sexChoice {
	case 1:
		player.Sex = "Homme"
	case 2:
		player.Sex = "Femme"
	}

	// Choisir la classe
	for {
		fmt.Println(cyan + "\n================================================" + reset)
		fmt.Println("   Choisissez votre classe :")
		fmt.Println(cyan + "================================================" + reset)
		fmt.Println(yellow + "1" + reset + " - Iop")
		fmt.Println(yellow + "2" + reset + " - Crâ")
		fmt.Println(yellow + "3" + reset + " - Osamodas")
		fmt.Println(yellow + "4" + reset + " - Eniripsa")
		fmt.Println(cyan + "================================================" + reset)

		var classChoice string
		fmt.Print("Choix : ")
		fmt.Scan(&classChoice)
		// Vérification que l'entrée contient exactement un caractère, qui doit être un chiffre entre '1' et '4'
		if len(classChoice) == 1 && unicode.IsDigit(rune(classChoice[0])) {
			switch classChoice {
			case "1":
				player.Class = "Iop"
				fmt.Println(green + "Vous avez choisi : Iop" + reset)
				player.HealthMax = 200
				player.Health = 200
				player.Inventaire.Potions = 3
				player.Level = 1
				player.Gold = 100
				player.Attack = 10
				player.Skills = append(player.Skills, "Coup de poing")
				fmt.Println(green + "Vous avez appris une nouvelle compétence : Coup de poing" + reset)
				player.equip.Head = stuff{name: "Chapeau de l'aventurier", pvb: 0}
				player.equip.Plastron = stuff{name: "Plastron de l'aventurier", pvb: 0}
				player.equip.Feet = stuff{name: "Botte de l'aventurier", pvb: 0}
				return player // Quitte la boucle après un choix valide
			case "2":
				player.Class = "Crâ"
				fmt.Println(green + "Vous avez choisi : Crâ" + reset)
				player.HealthMax = 100
				player.Health = 100
				player.Inventaire.Potions = 3
				player.Level = 1
				player.Gold = 100
				player.Attack = 30
				player.Skills = append(player.Skills, "Coup de poing")
				fmt.Println(green + "Vous avez appris une nouvelle compétence : Coup de poing" + reset)
				player.equip.Head = stuff{name: "Chapeau de l'aventurier", pvb: 0}
				player.equip.Plastron = stuff{name: "Plastron de l'aventurier", pvb: 0}
				player.equip.Feet = stuff{name: "Botte de l'aventurier", pvb: 0}
				return player // Quitte la boucle après un choix valide

			case "3":
				player.Class = "Osamodas"
				fmt.Println(green + "Vous avez choisi : Osamodas" + reset)
				player.HealthMax = 150
				player.Health = 250
				player.Inventaire.Potions = 3
				player.Level = 1
				player.Gold = 100
				player.Attack = 20
				player.Skills = append(player.Skills, "Coup de poing")
				fmt.Println(green + "Vous avez appris une nouvelle compétence : Coup de poing" + reset)
				player.equip.Head = stuff{name: "Chapeau de l'aventurier", pvb: 0}
				player.equip.Plastron = stuff{name: "Plastron de l'aventurier", pvb: 0}
				player.equip.Feet = stuff{name: "Botte de l'aventurier", pvb: 0}
				return player
			case "4":
				player.Class = "Eniripsa"
				fmt.Println(green + "Vous avez choisi : Eniripsa" + reset)
				player.HealthMax = 100
				player.Health = 40
				player.Inventaire.Potions = 3
				player.Level = 1
				player.Gold = 100
				player.Attack = 10
				player.Skills = append(player.Skills, "Coup de poing")
				fmt.Println(green + "Vous avez appris une nouvelle compétence : Coup de poing" + reset)
				player.equip.Head = stuff{name: "Chapeau de l'aventurier", pvb: 0}
				player.equip.Plastron = stuff{name: "Plastron de l'aventurier", pvb: 0}
				player.equip.Feet = stuff{name: "Botte de l'aventurier", pvb: 0}
				return player
			default:
				fmt.Println(red + "Choix invalide, veuillez entrer un chiffre entre 1 et 4." + reset)
			}
		} else {
			fmt.Println(red + "Choix invalide, veuillez entrer un seul chiffre entre 1 et 4." + reset)
		}
	}
}

// Fonction pour vérifier que le pseudo ne contient que des lettres
func isValidPseudo(pseudo string) bool {
	for _, char := range pseudo {
		if !unicode.IsLetter(char) {
			return false
		}
	}
	return true
}
