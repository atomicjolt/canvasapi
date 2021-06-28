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

// ListLicensesUsers A paginated list of licenses that can be applied
// https://canvas.instructure.com/doc/api/files.html
//
// Path Parameters:
// # UserID (Required) ID
//
type ListLicensesUsers struct {
	Path struct {
		UserID string `json:"user_id" url:"user_id,omitempty"` //  (Required)
	} `json:"path"`
}

func (t *ListLicensesUsers) GetMethod() string {
	return "GET"
}

func (t *ListLicensesUsers) GetURLPath() string {
	path := "users/{user_id}/content_licenses"
	path = strings.ReplaceAll(path, "{user_id}", fmt.Sprintf("%v", t.Path.UserID))
	return path
}

func (t *ListLicensesUsers) GetQuery() (string, error) {
	return "", nil
}

func (t *ListLicensesUsers) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *ListLicensesUsers) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *ListLicensesUsers) HasErrors() error {
	errs := []string{}
	if t.Path.UserID == "" {
		errs = append(errs, "'UserID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *ListLicensesUsers) Do(c *canvasapi.Canvas) ([]*models.License, error) {
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
