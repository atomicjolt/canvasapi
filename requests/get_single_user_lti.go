package requests

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/url"
	"strings"

	"github.com/atomicjolt/canvasapi"
	"github.com/atomicjolt/canvasapi/models"
)

// GetSingleUserLti Get a single Canvas user by Canvas id or LTI id. Tool providers may only access
// users that have been assigned an assignment associated with their tool.
// https://canvas.instructure.com/doc/api/plagiarism_detection_platform_users.html
//
// Path Parameters:
// # ID (Required) ID
//
type GetSingleUserLti struct {
	Path struct {
		ID string `json:"id" url:"id,omitempty"` //  (Required)
	} `json:"path"`
}

func (t *GetSingleUserLti) GetMethod() string {
	return "GET"
}

func (t *GetSingleUserLti) GetURLPath() string {
	path := "/lti/users/{id}"
	path = strings.ReplaceAll(path, "{id}", fmt.Sprintf("%v", t.Path.ID))
	return path
}

func (t *GetSingleUserLti) GetQuery() (string, error) {
	return "", nil
}

func (t *GetSingleUserLti) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *GetSingleUserLti) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *GetSingleUserLti) HasErrors() error {
	errs := []string{}
	if t.Path.ID == "" {
		errs = append(errs, "'ID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *GetSingleUserLti) Do(c *canvasapi.Canvas) (*models.User, error) {
	response, err := c.SendRequest(t)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		return nil, err
	}
	ret := models.User{}
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return nil, err
	}

	return &ret, nil
}
