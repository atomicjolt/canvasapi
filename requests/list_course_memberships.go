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

// ListCourseMemberships Return active NamesAndRoleMemberships in the given course.
// https://canvas.instructure.com/doc/api/names_and_role.html
//
// Path Parameters:
// # CourseID (Required) ID
//
// Query Parameters:
// # Rlid (Optional) If specified only NamesAndRoleMemberships with access to the LTI link references by this `rlid` will be included.
//    Also causes the member array to be included for each returned NamesAndRoleMembership.
//    If the `role` parameter is also present, it will be 'and-ed' together with this parameter
// # Role (Optional) If specified only NamesAndRoleMemberships having this role in the given Course will be included.
//    Value must be a fully-qualified LTI/LIS role URN.
//    If the `rlid` parameter is also present, it will be 'and-ed' together with this parameter
// # Limit (Optional) May be used to limit the number of NamesAndRoleMemberships returned in a page. Defaults to 50.
//
type ListCourseMemberships struct {
	Path struct {
		CourseID string `json:"course_id" url:"course_id,omitempty"` //  (Required)
	} `json:"path"`

	Query struct {
		Rlid  string `json:"rlid" url:"rlid,omitempty"`   //  (Optional)
		Role  string `json:"role" url:"role,omitempty"`   //  (Optional)
		Limit string `json:"limit" url:"limit,omitempty"` //  (Optional)
	} `json:"query"`
}

func (t *ListCourseMemberships) GetMethod() string {
	return "GET"
}

func (t *ListCourseMemberships) GetURLPath() string {
	path := "/lti/courses/{course_id}/names_and_roles"
	path = strings.ReplaceAll(path, "{course_id}", fmt.Sprintf("%v", t.Path.CourseID))
	return path
}

func (t *ListCourseMemberships) GetQuery() (string, error) {
	v, err := query.Values(t.Query)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("?%v", v.Encode()), nil
}

func (t *ListCourseMemberships) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *ListCourseMemberships) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *ListCourseMemberships) HasErrors() error {
	errs := []string{}
	if t.Path.CourseID == "" {
		errs = append(errs, "'CourseID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *ListCourseMemberships) Do(c *canvasapi.Canvas) (*models.NamesAndRoleMemberships, error) {
	response, err := c.SendRequest(t)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		return nil, err
	}
	ret := models.NamesAndRoleMemberships{}
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return nil, err
	}

	return &ret, nil
}
