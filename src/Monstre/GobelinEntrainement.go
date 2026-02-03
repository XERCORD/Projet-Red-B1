package ennemie

// Initialiser un monstre d'entrainement
func InitGoblin() Monstre {
	return Monstre{
		Name:      "Gobelin d'entraînement",
		HealthMax: 40,
		Health:    40,
		Attack:    5,
	}
}
