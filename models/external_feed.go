package models

import (
	"fmt"
	"time"

	"github.com/atomicjolt/string_utils"
)

type ExternalFeed struct {
	ID          int64     `json:"id" url:"id,omitempty"`                     // The ID of the feed.Example: 5
	DisplayName string    `json:"display_name" url:"display_name,omitempty"` // The title of the feed, pulled from the feed itself. If the feed hasn't yet been pulled, a temporary name will be synthesized based on the URL.Example: My Blog
	Url         string    `json:"url" url:"url,omitempty"`                   // The HTTP/HTTPS URL to the feed.Example: http://example.com/myblog.rss
	HeaderMatch string    `json:"header_match" url:"header_match,omitempty"` // If not null, only feed entries whose title contains this string will trigger new posts in Canvas.Example: pattern
	CreatedAt   time.Time `json:"created_at" url:"created_at,omitempty"`     // When this external feed was added to Canvas.Example: 2012-06-01T00:00:00-06:00
	Verbosity   string    `json:"verbosity" url:"verbosity,omitempty"`       // The verbosity setting determines how much of the feed's content is imported into Canvas as part of the posting. 'link_only' means that only the title and a link to the item. 'truncate' means that a summary of the first portion of the item body will be used. 'full' means that the full item body will be used..Example: truncate
}

func (t *ExternalFeed) HasErrors() error {
	var s []string
	errs := []string{}
	s = []string{"link_only", "truncate", "full"}
	if t.Verbosity != "" && !string_utils.Include(s, t.Verbosity) {
		errs = append(errs, fmt.Sprintf("expected 'Verbosity' to be one of %v", s))
	}
	return nil
}
