package personnage

// Structure du personnage
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

// Structure des Objet dans l'Inventaire
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

// Strucutre de l'equipement
type Equipment struct {
	Head     stuff
	Plastron stuff
	Feet     stuff
}

type stuff struct {
	name string
	pvb  int
}

// Structure du Monstre
type Monstre struct {
	Name      string
	HealthMax int
	Health    int
	Attack    int
}

const (
	minPseudoLength = 3  // Longueur minimale du pseudo
	maxPseudoLength = 12 // Longueur maximale du pseudo
)
