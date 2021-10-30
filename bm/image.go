package bm

import "net/url"

// Image represents multimedia content for resources.
type Image struct {
	Description string   `json:"description,omitempty"`
	URL         *url.URL `json:"url,omitempty"`
	SortOrder   int64    `json:"sortOrder,omitempty"`
}
