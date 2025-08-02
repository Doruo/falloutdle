# Falloutdle Project Structure

```
falloutdle/
│
├── internal/                           # domains
│   │
│   ├── character/              
│   │   ├── model.go                    # character struct
│   │   ├── repository.go               # db interface + GORM
│   │   ├── gamecode.go                 # games code references
│   │   └── service.go                  # character logic interface
│   │
│   ├── game/               
│   │   ├── model.go                    # game struct
│   │   ├── repository.go               # db interface
│   │   └── service.go                  # game logic
│   │
│   └── database/
│       ├── model.go                    # GORM database struct
│       └── connection.go               # GORM database connection
│
├── external/
│   │
│   └── wiki/                           # external wiki api
│       ├── client.go                   # external wiki api client
│       └── response.go                 # external wiki api response
│
├── tests/                              # tests
│   ├── database_test.go                # database communication test
│   └── wiki_test.go                    # wiki api requests test
│   └── character_test.go               # characters test
│
├── cmd/                        
│   └── server/                         # entry point
│       ├── handlers/                   # api http requests handling
│       │    ├── routes/
│       │    │  └── routes/             # api routes
│       │    ├── game_handler.go        # game routes handling
│       │    └── character_handler.go   # characters routers handling
│       │               
│       └── main.go                     # main server
│
├── pkg/                    
│   └── libs/                           # public packages
│
├── .env.example                        # example attributs to use in env
├── .gitignore                  
├── go.mod                              # project dependencies
├── go.sum                              # project dependencies checksums
└── README.md
```