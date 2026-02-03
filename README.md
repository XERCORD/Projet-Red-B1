# Monde des Douze — Jeu de combat en terminal

Jeu de rôle textuel en Go, inspiré de l'univers Dofus. Vous incarnez un aventurier dans le Monde des Douze : créez votre personnage, récoltez des ressources, combattez des monstres et gérez votre inventaire depuis le terminal.

---

## Élèves participants

- **Xerly JI** — [@XERCORD](https://github.com/XERCORD)
- **Eddy AMIR** — [@Kottah02](https://github.com/Kottah02)
- **Etienne LE BERRE** — [@etiennelb95](https://github.com/etiennelb95)

---

## Prérequis

- **Go 1.23** ou supérieur ([télécharger Go](https://go.dev/dl/))

Vérifier l’installation :

```bash
go version
```

---

## Démarrer le jeu

Le point d’entrée du jeu est dans le dossier `src/`. À la racine du projet :

**Windows (PowerShell / CMD) :**
```bash
cd src
go run .
```

**Linux / macOS :**
```bash
cd src && go run .
```

**Alternative (depuis n’importe quel dossier) :**
```bash
go run ./src/...
```
*(Si votre `GOPATH` ou modules sont configurés pour le dépôt.)*

Au lancement, le jeu affiche l’intro puis vous demande de **créer votre personnage** (pseudo, sexe, classe). Ensuite le **menu principal** s’affiche : choisissez une action en tapant le numéro puis Entrée.

---

## Menu principal

| Choix | Action |
|-------|--------|
| **1** | Afficher les infos du personnage |
| **2** | Récolter des ressources (bois, pierre, feuilles, fourrures, etc.) |
| **3** | Combattre des monstres (avec option potion avant le combat) |
| **4** | Forgeron (fabriquer armes et équipement) |
| **5** | Consulter l’inventaire (et utiliser potions / livre Boule de Feu) |
| **6** | Marchand (acheter potions, livres, ressources, armes) |
| **7** | Quitter le jeu |

---

## Création du personnage

- **Pseudo** : 3 à 12 caractères, lettres uniquement.
- **Sexe** : Homme (1) ou Femme (2).
- **Classe** :
  - **Iop** — guerrier
  - **Crâ** — combattant
  - **Osamodas** — invocateur
  - **Eniripsa** — soigneur

Chaque classe a des points de vie et des caractéristiques différents.

---

## Structure du projet

```
Projet-Red-B1/
├── README.md           # Ce fichier
├── go.mod              # Module racine (challenge)
├── go.sum
├── src/                # Jeu principal (module projetred)
│   ├── go.mod
│   ├── main.go         # Point d’entrée + menu principal
│   ├── combat/
│   │   └── Combat.go   # Logique de combat
│   ├── personnage/
│   │   ├── structure.go
│   │   ├── CreationPerso.go
│   │   ├── AfficheInfo.go
│   │   ├── AfficheInv.go
│   │   ├── ConstruireObjet.go
│   │   ├── RecolteRessources.go
│   │   ├── PrendrePotion.go
│   │   ├── MenuMarchand.go
│   │   ├── LivreBouledeFeu.go
│   │   └── Couleur.go
│   └── Monstre/
│       ├── structure.go
│       ├── GobelinEntrainement.go
│       └── IAGobelin.go
└── test.go             # Ancienne version / tests (optionnel)
```

---

## Résumé des fonctionnalités

- **Personnage** : pseudo, sexe, classe, PV, or, équipement, compétences.
- **Combat** : prendre une potion avant d’affronter un monstre, dégâts tour par tour, mort puis résurrection avec la moitié des PV max.
- **Ressources** : bois, pierre, feuilles, fourrure, peau de troll, cuir de sanglier, plume de corbeau (inventaire limité à 50 emplacements).
- **Forgeron** : épée, arc, bâton magique, coiffe/cape/bottes Bouftou (coûts en ressources et or).
- **Marchand** : potions, livre Boule de Feu, ressources, armes.
- **Inventaire** : consulter objets/équipement/ressources, utiliser potion ou livre de compétence Boule de Feu.

---

## Licence

Projet à but pédagogique (Ynov B1). Libre d’utilisation et de modification selon les règles de votre formation.

---

*Bonne aventure dans le Monde des Douze.*
