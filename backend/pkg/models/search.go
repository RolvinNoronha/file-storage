package models

// SearchResult holds the complete ES search response
type SearchResult struct {
	Hits struct {
		Total struct {
			Value    int    `json:"value"`
			Relation string `json:"relation"`
		} `json:"total"`
		Hits []*Hit `json:"hits"`
	} `json:"hits"`
}

// Hit represents a single search result hit
type Hit struct {
	Source FileDocument `json:"_source"` // This will hold our FileDocument
}
