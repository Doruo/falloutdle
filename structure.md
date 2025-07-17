# Architecture du projet Falloutdle

## Arborescence du projet

```
falloutdle/
├── cmd/
│   └── server/
│       └── main.go                 # Point d'entrée
├── internal/
│   ├── api/
│   │   ├── handlers/
│   │   │   ├── character.go        # Handlers pour personnages
│   │   │   ├── game.go            # Handlers pour jeu quotidien
│   │   │   └── health.go          # Health check
│   │   ├── middleware/
│   │   │   ├── cors.go
│   │   │   ├── logging.go
│   │   │   └── ratelimit.go
│   │   └── routes/
│   │       └── routes.go          # Configuration des routes
│   ├── domain/
│   │   ├── models/
│   │   │   ├── character.go       # Structures de données
│   │   │   ├── game.go
│   │   │   └── guess.go
│   │   └── services/
│   │       ├── character_service.go
│   │       ├── game_service.go
│   │       └── scraper_service.go
│   ├── infrastructure/
│   │   ├── database/
│   │   │   ├── migrations/
│   │   │   ├── connection.go
│   │   │   └── queries.go
│   │   ├── external/
│   │   │   └── wiki_client.go     # Client pour API MediaWiki
│   │   └── cache/
│   │       └── redis.go
│   └── config/
│       └── config.go              # Configuration
├── pkg/
│   ├── logger/
│   │   └── logger.go              # Logger personnalisé
│   └── utils/
│       ├── response.go            # Réponses HTTP standardisées
│       └── validator.go           # Validation des données
├── scripts/
│   └── scraper/
│       └── main.go                # Script de récupération des données
├── data/
│   ├── characters.json            # Données récupérées
│   └── migrations/
├── docker/
│   ├── Dockerfile
│   └── docker-compose.yml
├── .env.example
├── .gitignore
├── go.mod
├── go.sum
├── Makefile
└── README.md
```

## Vue d'ensemble

Falloutdle est une application web de type "Wordle" basée sur l'univers Fallout, développée en Go avec une architecture hexagonale (Clean Architecture) pour assurer la maintenabilité et la testabilité.

## Structure du projet

### 📁 `cmd/`
Point d'entrée de l'application
- **`server/main.go`** : Point d'entrée principal du serveur HTTP

### 📁 `internal/`
Code métier de l'application (non exposé publiquement)

#### 🌐 `api/`
Couche présentation - Interface HTTP

- **`handlers/`** : Gestionnaires des requêtes HTTP
  - `character.go` : Endpoints pour la gestion des personnages
  - `game.go` : Endpoints pour le jeu quotidien (soumission de tentatives, récupération du jeu du jour)
  - `health.go` : Endpoint de santé pour le monitoring

- **`middleware/`** : Middlewares HTTP
  - `cors.go` : Configuration CORS pour les requêtes cross-origin
  - `logging.go` : Logging des requêtes HTTP
  - `ratelimit.go` : Limitation du taux de requêtes

- **`routes/`** : Configuration du routage
  - `routes.go` : Définition et configuration de toutes les routes

#### 🏗️ `domain/`
Couche métier - Logique business

- **`models/`** : Structures de données métier
  - `character.go` : Modèle des personnages Fallout
  - `game.go` : Modèle du jeu quotidien
  - `guess.go` : Modèle des tentatives de devinettes

- **`services/`** : Services métier
  - `character_service.go` : Logique métier des personnages
  - `game_service.go` : Logique du jeu (génération quotidienne, validation des tentatives)
  - `scraper_service.go` : Service de récupération de données depuis les sources externes

#### 🔧 `infrastructure/`
Couche infrastructure - Accès aux données et services externes

- **`database/`** : Accès aux données
  - `migrations/` : Scripts de migration de base de données
  - `connection.go` : Configuration et connexion à la base de données
  - `queries.go` : Requêtes SQL

- **`external/`** : Clients pour services externes
  - `wiki_client.go` : Client pour l'API MediaWiki (récupération des données Fallout)

- **`cache/`** : Système de cache
  - `redis.go` : Configuration et utilisation de Redis

#### ⚙️ `config/`
Configuration de l'application
- **`config.go`** : Gestion centralisée de la configuration

### 📦 `pkg/`
Utilitaires réutilisables (exposés publiquement)

- **`logger/`** : Système de logging
  - `logger.go` : Logger personnalisé avec niveaux et formatage

- **`utils/`** : Utilitaires généraux
  - `response.go` : Structures de réponses HTTP standardisées
  - `validator.go` : Validation des données d'entrée

### 🔄 `scripts/`
Scripts d'administration et de maintenance

- **`scraper/`** : Scripts de récupération de données
  - `main.go` : Script principal pour récupérer les données des personnages depuis les wikis

### 📊 `data/`
Données de l'application

- **`characters.json`** : Données des personnages récupérées
- **`migrations/`** : Scripts de migration de base de données

### 🐳 `docker/`
Configuration Docker

- **`Dockerfile`** : Image Docker de l'application
- **`docker-compose.yml`** : Orchestration des services (app + Redis + PostgreSQL)

### 📋 Fichiers de configuration racine

- **`.env.example`** : Modèle de variables d'environnement
- **`.gitignore`** : Fichiers ignorés par Git
- **`go.mod`** : Module Go et dépendances
- **`go.sum`** : Checksums des dépendances
- **`Makefile`** : Commandes de build et déploiement
- **`README.md`** : Documentation du projet

## Flux de données

```
Requête HTTP → Middleware → Handler → Service → Repository → Database
                                   ↓
                                Cache (Redis)
```

## Principes architecturaux

### Clean Architecture
- **Séparation des responsabilités** : Chaque couche a une responsabilité spécifique
- **Inversion des dépendances** : Les couches internes ne dépendent pas des couches externes
- **Testabilité** : Chaque couche peut être testée indépendamment

### Avantages de cette structure

1. **Maintenabilité** : Code organisé et facile à maintenir
2. **Testabilité** : Chaque composant peut être testé unitairement
3. **Évolutivité** : Ajout de nouvelles fonctionnalités facilité
4. **Séparation des préoccupations** : Chaque package a une responsabilité claire

## Technologies utilisées

- **Go** : Langage principal
- **PostgreSQL** : Base de données principale
- **Redis** : Cache et sessions
- **Docker** : Containerisation
- **MediaWiki API** : Source de données des personnages

## Commandes utiles

```bash
# Build de l'application
make build

# Lancement en développement
make dev

# Tests
make test

# Lancement du scraper
make scrape

# Déploiement Docker
make docker-up
```