# Falloutdle Project structure

```
falloutdle/
│
├── internal/                   # domains
│   │
│   ├── character/              
│   │   ├── model.go            # character struct
│   │   ├── repository.go       # database interface + GORM
│   │   ├── gamecode.go         # games code references
│   │   └── service.go          # character logic interface
│   │
│   ├── game/               
│   │   ├── model.go            # game structure
│   │   └── service.go          # game logic
│   │
│   └── database/
│       └── connection.go       # GORM database connection
│
├── external/
│   │
│   └── wiki/                   # external wiki api
│       ├── client.go           # external wiki api client
│       └── response.go         # external wiki api response
│
├── tests/                      # tests
│   ├── database_test.go        # database communication test
│   └── wiki_test.go            # wiki api requests test
│
├── cmd/                        
│   └── server/                 # entry point
│       ├── handlers/           # api http requests handling
│       │    ├── handler.go     # api handling
│       │    └── routes.go      # api routes
│       └── main.go             # main server
│
├── pkg/                    
│   └── libs/                   # public packages
│
├── .env.example                # example attributs to use in env
├── .gitignore                  
├── go.mod                      # project dependencies
├── go.sum                      # project dependencies checksums
└── README.md
```