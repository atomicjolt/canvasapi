package requests

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/google/go-querystring/query"

	"github.com/atomicjolt/canvasapi"
	"github.com/atomicjolt/canvasapi/models"
)

// SearchForContentShareUsers Returns a paginated list of users you can share content with.  Requires the content share
// feature and the user must have the manage content permission for the course.
// https://canvas.instructure.com/doc/api/courses.html
//
// Path Parameters:
// # Path.CourseID (Required) ID
//
// Query Parameters:
// # Query.SearchTerm (Required) Term used to find users.  Will search available share users with the search term in their name.
//
type SearchForContentShareUsers struct {
	Path struct {
		CourseID string `json:"course_id" url:"course_id,omitempty"` //  (Required)
	} `json:"path"`

	Query struct {
		SearchTerm string `json:"search_term" url:"search_term,omitempty"` //  (Required)
	} `json:"query"`
}

func (t *SearchForContentShareUsers) GetMethod() string {
	return "GET"
}

func (t *SearchForContentShareUsers) GetURLPath() string {
	path := "courses/{course_id}/content_share_users"
	path = strings.ReplaceAll(path, "{course_id}", fmt.Sprintf("%v", t.Path.CourseID))
	return path
}

func (t *SearchForContentShareUsers) GetQuery() (string, error) {
	v, err := query.Values(t.Query)
	if err != nil {
		return "", err
	}
	return v.Encode(), nil
}

func (t *SearchForContentShareUsers) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *SearchForContentShareUsers) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *SearchForContentShareUsers) HasErrors() error {
	errs := []string{}
	if t.Path.CourseID == "" {
		errs = append(errs, "'Path.CourseID' is required")
	}
	if t.Query.SearchTerm == "" {
		errs = append(errs, "'Query.SearchTerm' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *SearchForContentShareUsers) Do(c *canvasapi.Canvas, next *url.URL) ([]*models.User, *canvasapi.PagedResource, error) {
	var err error
	var response *http.Response
	if next != nil {
		response, err = c.Send(next, t.GetMethod(), nil)
	} else {
		response, err = c.SendRequest(t)
	}

	if err != nil {
		return nil, nil, err
	}
	if err != nil {
		return nil, nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		return nil, nil, err
	}
	ret := []*models.User{}
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return nil, nil, err
	}

	pagedResource, err := canvasapi.ExtractPagedResource(response.Header)
	if err != nil {
		return nil, nil, err
	}

	return ret, pagedResource, nil
}
