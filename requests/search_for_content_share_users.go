package requests

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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
// # CourseID (Required) ID
//
// Query Parameters:
// # SearchTerm (Required) Term used to find users.  Will search available share users with the search term in their name.
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
		errs = append(errs, "'CourseID' is required")
	}
	if t.Query.SearchTerm == "" {
		errs = append(errs, "'SearchTerm' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *SearchForContentShareUsers) Do(c *canvasapi.Canvas) ([]*models.User, error) {
	response, err := c.SendRequest(t)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		return nil, err
	}
	ret := []*models.User{}
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return nil, err
	}

	return ret, nil
}
