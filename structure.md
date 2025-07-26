# Falloutdle Project structure

```
falloutdle/
│
├── internal/
│   │
│   ├── character/              # domain
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
│   └── wiki/                   # wiki api
│       ├── client.go           # wiki api client
│       └── response.go         # wiki api response
│
├── tests/
│   ├── database_test.go        # database communication test
│   └── wiki_test.go            # wiki api requests test
│
├── cmd/                        # entry point
│   └── server/              
│       ├── handlers/           # HTTP handle
│       │    ├── handler.go     # wiki api requests test
│       │    └── routes.go      # main server
│       └── main.go             # main server
│
├── pkg/                    
│   └── libs/                   # public packages
│
├── .env.example                # example attributs to use in env
├── .gitignore
├── go.mod                      # projetct depedancies
├── go.sum                      # project dependancies checksums
└── README.md                   # docs
```