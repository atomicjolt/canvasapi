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
	"github.com/atomicjolt/string_utils"
)

// ListYourGroups Returns a paginated list of active groups for the current user.
// https://canvas.instructure.com/doc/api/groups.html
//
// Query Parameters:
// # Query.ContextType (Optional) . Must be one of Account, CourseOnly include groups that are in this type of context.
// # Query.Include (Optional) . Must be one of tabs- "tabs": Include the list of tabs configured for each group.  See the
//      {api:TabsController#index List available tabs API} for more information.
//
type ListYourGroups struct {
	Query struct {
		ContextType string   `json:"context_type" url:"context_type,omitempty"` //  (Optional) . Must be one of Account, Course
		Include     []string `json:"include" url:"include,omitempty"`           //  (Optional) . Must be one of tabs
	} `json:"query"`
}

func (t *ListYourGroups) GetMethod() string {
	return "GET"
}

func (t *ListYourGroups) GetURLPath() string {
	return ""
}

func (t *ListYourGroups) GetQuery() (string, error) {
	v, err := query.Values(t.Query)
	if err != nil {
		return "", err
	}
	return v.Encode(), nil
}

func (t *ListYourGroups) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *ListYourGroups) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *ListYourGroups) HasErrors() error {
	errs := []string{}
	if t.Query.ContextType != "" && !string_utils.Include([]string{"Account", "Course"}, t.Query.ContextType) {
		errs = append(errs, "ContextType must be one of Account, Course")
	}
	for _, v := range t.Query.Include {
		if v != "" && !string_utils.Include([]string{"tabs"}, v) {
			errs = append(errs, "Include must be one of tabs")
		}
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *ListYourGroups) Do(c *canvasapi.Canvas, next *url.URL) ([]*models.Group, *canvasapi.PagedResource, error) {
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
	ret := []*models.Group{}
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
