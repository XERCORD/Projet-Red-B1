package ennemie


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

type Monstre struct {
	Name      string
	HealthMax int
	Health    int
	Attack    int
}
