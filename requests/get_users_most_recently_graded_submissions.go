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
	"github.com/atomicjolt/string_utils"
)

// GetUsersMostRecentlyGradedSubmissions
// https://canvas.instructure.com/doc/api/users.html
//
// Path Parameters:
// # ID (Required) ID
//
// Query Parameters:
// # Include (Optional) . Must be one of assignmentAssociations to include with the group
// # OnlyCurrentEnrollments (Optional) Returns submissions for only currently active enrollments
// # OnlyPublishedAssignments (Optional) Returns submissions for only published assignments
//
type GetUsersMostRecentlyGradedSubmissions struct {
	Path struct {
		ID string `json:"id" url:"id,omitempty"` //  (Required)
	} `json:"path"`

	Query struct {
		Include                  []string `json:"include" url:"include,omitempty"`                                       //  (Optional) . Must be one of assignment
		OnlyCurrentEnrollments   bool     `json:"only_current_enrollments" url:"only_current_enrollments,omitempty"`     //  (Optional)
		OnlyPublishedAssignments bool     `json:"only_published_assignments" url:"only_published_assignments,omitempty"` //  (Optional)
	} `json:"query"`
}

func (t *GetUsersMostRecentlyGradedSubmissions) GetMethod() string {
	return "GET"
}

func (t *GetUsersMostRecentlyGradedSubmissions) GetURLPath() string {
	path := "users/{id}/graded_submissions"
	path = strings.ReplaceAll(path, "{id}", fmt.Sprintf("%v", t.Path.ID))
	return path
}

func (t *GetUsersMostRecentlyGradedSubmissions) GetQuery() (string, error) {
	v, err := query.Values(t.Query)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("?%v", v.Encode()), nil
}

func (t *GetUsersMostRecentlyGradedSubmissions) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *GetUsersMostRecentlyGradedSubmissions) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *GetUsersMostRecentlyGradedSubmissions) HasErrors() error {
	errs := []string{}
	if t.Path.ID == "" {
		errs = append(errs, "'ID' is required")
	}
	for _, v := range t.Query.Include {
		if v != "" && !string_utils.Include([]string{"assignment"}, v) {
			errs = append(errs, "Include must be one of assignment")
		}
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *GetUsersMostRecentlyGradedSubmissions) Do(c *canvasapi.Canvas) ([]*models.Submission, error) {
	response, err := c.SendRequest(t)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		return nil, err
	}
	ret := []*models.Submission{}
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return nil, err
	}

	return ret, nil
}
