package main

import (
	"fmt"
	"math/rand"
	"time"
	"unicode"
)

// ANSI Color Codes
const (
	reset  = "\033[0m"
	red    = "\033[31m"
	green  = "\033[32m"
	yellow = "\033[33m"
	blue   = "\033[34m"
	cyan   = "\033[36m"
)

type Player struct {
	Pseudo     string
	Sex        string
	Class      string
	Level      int
	HealthMax  int
	Health     int
	Inventaire Objet
	Skills     []string
	Gold       int
	equip      Equipment
}

type Objet struct {
	Wood           int
	Stone          int
	Leaf           int
	Potions        int
	Potion_Poison  int
	Sword          int
	Bow            int
	MagicStaff     int
	SpellBookCount int
	Fourrure       int
	Peau_Troll     int
	CuirSanglier   int
	PlumeCorbeau   int
	ChapAven       int
	TunAven        int
}

type Equipment struct {
	Head     stuff
	Plastron stuff
	Feet     stuff
}

type stuff struct {
	name string
	pvb  int
}

const (
	minPseudoLength = 3  // Longueur minimale du pseudo
	maxPseudoLength = 12 // Longueur maximale du pseudo
)

func main() {
	fmt.Println("Bienvenue dans le Monde des Douze, les aventuriers cherchent les Dofus, des œufs de dragons aux pouvoirs immenses.")
	fmt.Println("Ces artefacts provoquent des conflits entre dieux, dragons et factions.")
	fmt.Println("La quête des Dofus primordiaux est centrale, influençant le destin du monde et entraînant des combats épiques.")
	fmt.Println(red + "LE MONDE COMPTE SUR TOI JEUNE AVENTURIER" + reset)
	fmt.Println("Appuyez sur Entrée pour continuer..." + reset)
	fmt.Scanln()

	rand.Seed(time.Now().UnixNano())

	// Choisir une classe
	player := createCharacter()
	monstre := InitGoblin()
	// Boucle principale du jeu
	for {

		fmt.Println("\n" + blue + "=================================================" + reset)
		fmt.Println(yellow + "   	     Que voulez-vous faire ?" + reset)
		fmt.Println("\n" + cyan + "|================================================|")
		fmt.Println(yellow + "1" + reset + " - Afficher l'Information du personnage")
		fmt.Println(yellow + "2" + reset + " - Récolter des ressources")
		fmt.Println(yellow + "3" + reset + " - Combattre des monstres")
		fmt.Println(yellow + "4" + reset + " - Forgeron")
		fmt.Println(yellow + "5" + reset + " - Consulter votre inventaire")
		fmt.Println(yellow + "6" + reset + " - Accéder au Marchand")
		fmt.Println(yellow + "7" + reset + " - Quitter le jeu")
		fmt.Println(cyan + "================================================" + reset)

		var choice int
		fmt.Print("Choix : ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			DisplayInfo(player)
		case 2:
			gatherResources(&player)
		case 3:
			combat(&player)
		case 4:
			craftItems(&player)
		case 5:
			accessInventory(&player)
		case 6:
			marchantMenu(&player)
		case 7:
			fmt.Println(green + "\nMerci d'avoir joué !" + reset)

			return
		case 8:
			entrainementGobelin(monstre, &player)
		default:
			fmt.Println(red + "Choix invalide." + reset)
		}
	}
}

func DisplayInfo(player Player) {
	fmt.Printf("Pseudo : %s\nSexe : %s\nClasse : %s\nNiveau : %d\nVie actuelle : %d\nVie Max : %d\nPièce d'or : %d\n Equipement :\n-Casque : %s\n-Plastron : %s\n-Bottes : %s\n", player.Pseudo, player.Sex, player.Class, player.Level, player.Health, player.HealthMax, player.Gold, player.equip.Head.name, player.equip.Plastron.name, player.equip.Feet.name)
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

// Fonction pour créer le personnage (pseudo, sexe, classe)
func createCharacter() Player {
	var player Player
	var pseudo string

	// Boucle pour obtenir un pseudo valide
	for {
		fmt.Println(cyan + "|================================================|" + reset)
		fmt.Println("   Entrez votre pseudo (lettres uniquement, entre", minPseudoLength, "et", maxPseudoLength, "caractères) :")
		fmt.Println(cyan + "|================================================|" + reset)
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
	fmt.Println(cyan + "\n|================================================|" + reset)
	fmt.Println("   Choisissez votre sexe :")
	fmt.Println(cyan + "|================================================|" + reset)
	fmt.Println(yellow + "1" + reset + " - Homme")
	fmt.Println(yellow + "2" + reset + " - Femme")
	fmt.Println(cyan + "|================================================|" + reset)

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
				player.Health = 100
				player.Inventaire.Potions = 3
				player.Level = 1
				player.Gold = 100
				return player // Quitte la boucle après un choix valide
			case "2":
				player.Class = "Crâ"
				fmt.Println(green + "Vous avez choisi : Crâ" + reset)
				player.HealthMax = 100
				player.Health = 50
				player.Inventaire.Potions = 3
				player.Level = 1
				player.Gold = 100
				return player
			case "3":
				player.Class = "Osamodas"
				fmt.Println(green + "Vous avez choisi : Osamodas" + reset)
				player.HealthMax = 250
				player.Health = 125
				player.Inventaire.Potions = 3
				player.Level = 1
				player.Gold = 100
				return player
			case "4":
				player.Class = "Eniripsa"
				fmt.Println(green + "Vous avez choisi : Eniripsa" + reset)
				player.HealthMax = 100
				player.Health = 50
				player.Inventaire.Potions = 3
				player.Level = 1
				player.Gold = 100
				return player
			default:
				fmt.Println(red + "Choix invalide, veuillez entrer un chiffre entre 1 et 4." + reset)
			}
		} else {
			fmt.Println(red + "Choix invalide, veuillez entrer un seul chiffre entre 1 et 4." + reset)
		}
	}

	player.Skills = append(player.Skills, "Coup de poing")
	fmt.Println(green + "Vous avez appris une nouvelle compétence : Coup de poing" + reset)

	player.equip.Head = stuff{name: "Chapeau de l'aventurier", pvb: 0}
	player.equip.Plastron = stuff{name: "Plastron de l'aventurier", pvb: 0}
	player.equip.Feet = stuff{name: "Botte de l'aventurier", pvb: 0}
	return player
}
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

// Fonction pour récolter des ressources
func gatherResources(player *Player) {
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

// Fonction pour combattre des monstres
func combat(player *Player) {
	monsterHealth := rand.Intn(50) + 50
	fmt.Printf(cyan + "\n================================================\n" + reset)
	fmt.Printf("   Vous combattez un monstre avec %d points de vie !\n", monsterHealth)
	fmt.Println(cyan + "================================================" + reset)

	for monsterHealth > 0 && player.Health > 0 {
		// Attaque du joueur
		damage := rand.Intn(20) + 10
		monsterHealth -= damage
		fmt.Printf(green+"Vous attaquez et infligez %d de dégâts. Vie du monstre restante : %d\n"+reset, damage, monsterHealth)

		if monsterHealth <= 0 {
			fmt.Println(green + "Vous avez vaincu le monstre !" + reset)
			break
		}

		// Attaque du monstre
		monsterDamage := rand.Intn(15) + 5
		player.Health -= monsterDamage
		fmt.Printf(red+"Le monstre vous attaque et inflige %d de dégâts. Votre vie restante : %d\n"+reset, monsterDamage, player.Health)

		if player.Health <= 0 {
			fmt.Println(red + "Vous êtes mort..." + reset)
			player.Health = player.HealthMax / 2
			fmt.Printf(green+"Vous avez été ressuscité avec %d points de vie.\n"+reset, player.Health)
		}
	}

	// Le joueur peut trouver une potion après le combat
	if rand.Float32() < 0.3 { // 30% de chance de trouver une potion
		player.Inventaire.Potions++
		fmt.Println(green + "Vous avez trouvé une potion !" + reset)
	}
}

// Fonction pour construire des objets
func craftItems(player *Player) {
	for {
		fmt.Println(cyan + "\n================================================" + reset)
		fmt.Println("   Que voulez-vous fabriquer ?")
		fmt.Println(cyan + "================================================" + reset)
		fmt.Println(yellow + "1" + reset + " - Épée (5 bois, 5 pierre)")
		fmt.Println(yellow + "2" + reset + " - Arc (5 bois, 5 feuilles)")
		fmt.Println(yellow + "3" + reset + " - Bâton magique (5 bois, 5 feuilles, 5 pierre)")
		fmt.Println(yellow + "4" + reset + " - Coiffe Bouftou (5 cuire de sanglier, 3 plume de corbeau)")
		fmt.Println(yellow + "5" + reset + " - Cape bouftou (4 fourrures de loup, 1 peau de Troll)")
		fmt.Println(yellow + "6" + reset + " - Boufbotte (3 fourrures de loup, 4 cuir de sanglier)")
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
				player.Inventaire.PlumeCorbeau -= 3
				player.Inventaire.CuirSanglier -= 5
				player.equip.Head = stuff{"Coiffe Bouftou", 10}
				player.HealthMax += player.equip.Head.pvb
				fmt.Println(green + "Vous avez fabriqué une Coiffe Bouftou." + reset)
			} else {
				fmt.Println(red + "Vous n'avez pas assez de ressources." + reset)
			}
		case 5:
			if player.Gold >= 5 && player.Inventaire.Fourrure >= 2 && player.Inventaire.Peau_Troll >= 1 {
				player.Gold -= 5
				player.Inventaire.Fourrure -= 4
				player.Inventaire.Peau_Troll -= 1
				player.equip.Head = stuff{"Cape Bouftou", 15}
				player.HealthMax += player.equip.Plastron.pvb
				fmt.Println(green + "Vous avez fabriqué une Cape Bouftou." + reset)
			} else {
				fmt.Println(red + "Vous n'avez pas assez de ressources." + reset)
			}
		case 6:
			if player.Gold >= 5 && player.Inventaire.Fourrure >= 1 && player.Inventaire.CuirSanglier >= 1 {
				player.Gold -= 5
				player.Inventaire.Fourrure -= 3
				player.Inventaire.CuirSanglier -= 4
				player.equip.Feet = stuff{"Boubfbotte", 5}
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

// Tâche 6
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

// Fonction pour afficher l'inventaire
func accessInventory(player *Player) {
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

// Tâche 7
func marchantMenu(player *Player) {
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

// Tâche 10
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

// Tâche 16
func InitGoblin() Monstre {
	return Monstre{
		Name:      "Gobelin d'entraînement",
		HealthMax: 40,
		Health:    40,
		Attack:    5,
	}
}

func goblinPattern(goblin Monstre, joueur *Player, tour int) {
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
