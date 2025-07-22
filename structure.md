# Falloutdle Project structure

```
falloutdle/
│
├── internal/
│   │
│   ├── character/           # domain
│   │   ├── model.go         # Character struct
│   │   ├── repository.go    # database interface + GORM
│   │   ├── gamecode.go      # fallout games code references
│   │   └── service.go       # character logic interface
│   │
│   └── database/
│       └── connection.go    # GORM database connection
│
├── external/
│   │
│   ├── wiki/                # fandom wiki api
│   │   ├── client.go        # wiki api client
│   │   └── response.go      # wiki api response
│   │
│   └── game/               
│       └── game.go          # game logic
│
├── tests/
│   ├── database_test.go     # database communication test
│   └── wiki_test.go         # wiki api requests test
│
├── server/
│   └── main.go              # main server
│
├── pkg/                    
│   └── utils/               # public code
│
├── .env.example             # example attributs to use in env
├── .gitignore
├── go.mod                   # depedancies
├── go.sum                   # dependancies checksums
└── README.md                # Docs
```