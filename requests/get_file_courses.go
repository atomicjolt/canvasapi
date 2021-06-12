package requests

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/google/go-querystring/query"

	"github.com/atomicjolt/canvasapi"
	"github.com/atomicjolt/canvasapi/models"
	"github.com/atomicjolt/string_utils"
)

// GetFileCourses Returns the standard attachment json object
// https://canvas.instructure.com/doc/api/files.html
//
// Path Parameters:
// # CourseID (Required) ID
// # ID (Required) ID
//
// Query Parameters:
// # Include (Optional) . Must be one of userArray of additional information to include.
//
//    "user":: the user who uploaded the file or last edited its content
//    "usage_rights":: copyright and license information for the file (see UsageRights)
//
type GetFileCourses struct {
	Path struct {
		CourseID string `json:"course_id"` //  (Required)
		ID       string `json:"id"`        //  (Required)
	} `json:"path"`

	Query struct {
		Include []string `json:"include"` //  (Optional) . Must be one of user
	} `json:"query"`
}

func (t *GetFileCourses) GetMethod() string {
	return "GET"
}

func (t *GetFileCourses) GetURLPath() string {
	path := "courses/{course_id}/files/{id}"
	path = strings.ReplaceAll(path, "{course_id}", fmt.Sprintf("%v", t.Path.CourseID))
	path = strings.ReplaceAll(path, "{id}", fmt.Sprintf("%v", t.Path.ID))
	return path
}

func (t *GetFileCourses) GetQuery() (string, error) {
	v, err := query.Values(t.Query)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("?%v", v.Encode()), nil
}

func (t *GetFileCourses) GetBody() (string, error) {
	return "", nil
}

func (t *GetFileCourses) HasErrors() error {
	errs := []string{}
	if t.Path.CourseID == "" {
		errs = append(errs, "'CourseID' is required")
	}
	if t.Path.ID == "" {
		errs = append(errs, "'ID' is required")
	}
	for _, v := range t.Query.Include {
		if !string_utils.Include([]string{"user"}, v) {
			errs = append(errs, "Include must be one of user")
		}
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *GetFileCourses) Do(c *canvasapi.Canvas) (*models.File, error) {
	response, err := c.SendRequest(t)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		return nil, err
	}
	ret := models.File{}
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return nil, err
	}

	return &ret, nil
}
