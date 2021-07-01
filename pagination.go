package canvasapi

import (
	"fmt"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
	"strings"
)

var resourceRegex = regexp.MustCompile(`<(.*?)>; rel="(.*?)"`)

type PagedResource struct {
	Current, First, Last, Next *PagedLink
}

type PagedLink struct {
	URL  *url.URL
	Page int
}

func ExtractPagedResource(header http.Header) (*PagedResource, error) {
	errs := []string{}
	pagedResource := &PagedResource{}
	links := header.Get("Link")
	parts := resourceRegex.FindAllStringSubmatch(links, -1)
	m := map[string]*PagedLink{}

	var err error
	for _, part := range parts {
		m[part[2]], err = newPagedLink(part[1])
		if err != nil {
			return pagedResource, err
		}
	}
	var ok bool
	if pagedResource.Current, ok = m["current"]; !ok {
		errs = append(errs, "could not find current link")
	}
	if pagedResource.First, ok = m["first"]; !ok {
		errs = append(errs, "could not find first link")
	}
	if pagedResource.Last, ok = m["last"]; !ok {
		errs = append(errs, "could not find last link")
	}
	pagedResource.Next, _ = m["next"]
	if len(errs) > 0 {
		return nil, fmt.Errorf(strings.Join(errs, ", "))
	}
	return pagedResource, nil
}

func newPagedLink(urlstr string) (*PagedLink, error) {
	u, err := url.Parse(urlstr)
	if err != nil {
		return nil, err
	}
	page, err := strconv.ParseInt(u.Query().Get("page"), 10, 32)
	if err != nil {
		return nil, fmt.Errorf("could not parse page num: %w", err)
	}
	return &PagedLink{
		URL:  u,
		Page: int(page),
	}, nil
}
