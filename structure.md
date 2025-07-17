# Falloutdle Project structure

```
falloutdle/
├── server/
│       └── main.go                 # Main
├── internal/
│   ├── api/
│   │   ├── handlers/
│   │   │   ├── character.go        # characters handling
│   │   │   ├── game.go            # daily game handling
│   │   │   └── health.go          # Health check
│   │   ├── middleware/
│   │   │   ├── cors.go
│   │   │   ├── logging.go
│   │   │   └── ratelimit.go
│   │   └── routes/
│   │       └── routes.go          # Routes configuration
│   ├── domain/
│   │   ├── models/
│   │   │   ├── character.go       # Database
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
│   │   │   └── wiki.go     #  MediaWiki API Client
│   │   └── cache/
│   │       └── redis.go
│   └── config/
│       └── config.go              # Configuration
├── pkg/
│   ├── logger/
│       └── logger.go              # Custom Logger
├── data/
│   └── characters.json            # Données récupérées
├── docker/
│   ├── Dockerfile
│   └── docker-compose.yml
├── .env.example
├── .gitignore
├── go.mod
├── go.sum
└── README.md
```