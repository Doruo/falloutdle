# Falloutdle Project structure

```
falloutdle/
│
├── internal/
│   │
│   ├── character/           # domain
│   │   ├── model.go         # character struct
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
│       ├── model.go         # game structure
│       └── service.go       # game logic
│
├── tests/
│   ├── database_test.go     # database communication test
│   └── wiki_test.go         # wiki api requests test
│
├── cmd/                     # entry point
│   └── server/              
│       └── main.go          # main server
│
├── pkg/                    
│   └── libs/                # public packages
│
├── .env.example             # example attributs to use in env
├── .gitignore
├── go.mod                   # projetct depedancies
├── go.sum                   # project dependancies checksums
└── README.md                # docs
```