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

// ListLicensesGroups A paginated list of licenses that can be applied
// https://canvas.instructure.com/doc/api/files.html
//
// Path Parameters:
// # Path.GroupID (Required) ID
//
type ListLicensesGroups struct {
	Path struct {
		GroupID string `json:"group_id" url:"group_id,omitempty"` //  (Required)
	} `json:"path"`
}

func (t *ListLicensesGroups) GetMethod() string {
	return "GET"
}

func (t *ListLicensesGroups) GetURLPath() string {
	path := "groups/{group_id}/content_licenses"
	path = strings.ReplaceAll(path, "{group_id}", fmt.Sprintf("%v", t.Path.GroupID))
	return path
}

func (t *ListLicensesGroups) GetQuery() (string, error) {
	return "", nil
}

func (t *ListLicensesGroups) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *ListLicensesGroups) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *ListLicensesGroups) HasErrors() error {
	errs := []string{}
	if t.Path.GroupID == "" {
		errs = append(errs, "'Path.GroupID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *ListLicensesGroups) Do(c *canvasapi.Canvas) ([]*models.License, error) {
	response, err := c.SendRequest(t)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		return nil, err
	}
	ret := []*models.License{}
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return nil, err
	}

	return ret, nil
}
