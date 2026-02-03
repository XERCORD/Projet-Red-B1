package main

import (
	"fmt"
	"math/rand"
	"projetred/combat"
	personnage "projetred/personnage"

	"time"
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
	Attack     int
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
	head     stuff
	plastron stuff
	feet     stuff
}

type stuff struct {
	name string
	pvb  int
}

type Monstre struct {
	Name      string
	HealthMax int
	Health    int
	Attack    int
}

func main() {
	fmt.Println("Bienvenue dans le Monde des Douze, les aventuriers cherchent les Dofus, des œufs de dragons aux pouvoirs immenses.")
	fmt.Println("Ces artefacts provoquent des conflits entre dieux, dragons et factions.")
	fmt.Println("La quête des Dofus primordiaux est centrale, influençant le destin du monde et entraînant des combats épiques.")
	fmt.Println(red + "LE MONDE COMPTE SUR TOI JEUNE AVENTURIER" + reset)
	fmt.Println("Appuyez sur Entrée pour continuer..." + reset)
	// définir un gobelin pour l'entrainement

	rand.Seed(time.Now().UnixNano())

	// Choisir une classe
	player := personnage.CreateCharacter()

	// Boucle principale du jeu
	for {

		fmt.Println("\n" + cyan + "================================================" + reset)
		fmt.Println("   Que voulez-vous faire ?")
		fmt.Println("================================================")
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
			personnage.DisplayInfo(player)
		case 2:
			personnage.GatherResources(&player)
		case 3:
			combat.Combat(&player)
		case 4:
			personnage.CraftItems(&player)
		case 5:
			personnage.AccessInventory(&player)
		case 6:
			personnage.MarchantMenu(&player)
		case 7:
			fmt.Println(green + "\nMerci d'avoir joué !" + reset)
			return
		default:
			fmt.Println(red + "Choix invalide." + reset)
		}
	}
}
