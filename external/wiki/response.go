package wiki

// WikiResponse represents the MediaWiki API response structure
type WikiResponse struct {
	Query struct {
		Pages map[string]WikiPage `json:"pages"`
	} `json:"query"`
}

// WikiPage represents a single wiki page
type WikiPage struct {
	PageID    int            `json:"pageid"`
	Title     string         `json:"title"`
	Revisions []WikiRevision `json:"revisions"`
}

// WikiRevision represents a page revision
type WikiRevision struct {
	Slots struct {
		Main struct {
			ContentFormat string `json:"contentformat"`
			ContentModel  string `json:"contentmodel"`
			Content       string `json:"*"`
		} `json:"main"`
	} `json:"slots"`
}

// CategoryResponse represents the category members API response response when querying category members
type CategoryResponse struct {
	Query struct {
		CategoryMembers []CategoryMember `json:"categorymembers"`
	} `json:"query"`
}

// CategoryMember represents a member of a category
type CategoryMember struct {
	PageID int    `json:"pageid"`
	Title  string `json:"title"`
	NS     int    `json:"ns"`
}
