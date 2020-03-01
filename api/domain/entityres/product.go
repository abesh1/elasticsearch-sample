package entityres

type Suggestion struct {
	Author  interface{} `json:"author"`
	Product interface{} `json:"product"`
	Query   string      `json:"query"`
}

type SuggestionAuthor struct {
	ID   uint64 `json:"id"`
	Name string `json:"name"`
}

type SuggestionAuthorList []SuggestionAuthor

type SuggestionProduct struct {
	ID   uint64 `json:"id"`
	Name string `json:"name"`
}

type SuggestionProductList []SuggestionProduct
