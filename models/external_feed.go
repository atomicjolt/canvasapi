package models

import (
	"fmt"
	"time"

	"github.com/atomicjolt/string_utils"
)

type ExternalFeed struct {
	ID          int64     `json:"id"`           // The ID of the feed.Example: 5
	DisplayName string    `json:"display_name"` // The title of the feed, pulled from the feed itself. If the feed hasn't yet been pulled, a temporary name will be synthesized based on the URL.Example: My Blog
	Url         string    `json:"url"`          // The HTTP/HTTPS URL to the feed.Example: http://example.com/myblog.rss
	HeaderMatch string    `json:"header_match"` // If not null, only feed entries whose title contains this string will trigger new posts in Canvas.Example: pattern
	CreatedAt   time.Time `json:"created_at"`   // When this external feed was added to Canvas.Example: 2012-06-01T00:00:00-06:00
	Verbosity   string    `json:"verbosity"`    // The verbosity setting determines how much of the feed's content is imported into Canvas as part of the posting. 'link_only' means that only the title and a link to the item. 'truncate' means that a summary of the first portion of the item body will be used. 'full' means that the full item body will be used..Example: truncate
}

func (t *ExternalFeed) HasError() error {
	var s []string
	s = []string{"link_only", "truncate", "full"}
	if !string_utils.Include(s, t.Verbosity) {
		return fmt.Errorf("expected 'verbosity' to be one of %v", s)
	}
	return nil
}
