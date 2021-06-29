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

// UpdateSharedTheme Update the specified shared_brand_config with a new name or to point to a new brand_config.
// Uses same parameters as create.
// https://canvas.instructure.com/doc/api/shared_brand_configs.html
//
// Path Parameters:
// # Path.AccountID (Required) ID
// # Path.ID (Required) ID
//
type UpdateSharedTheme struct {
	Path struct {
		AccountID string `json:"account_id" url:"account_id,omitempty"` //  (Required)
		ID        string `json:"id" url:"id,omitempty"`                 //  (Required)
	} `json:"path"`
}

func (t *UpdateSharedTheme) GetMethod() string {
	return "PUT"
}

func (t *UpdateSharedTheme) GetURLPath() string {
	path := "accounts/{account_id}/shared_brand_configs/{id}"
	path = strings.ReplaceAll(path, "{account_id}", fmt.Sprintf("%v", t.Path.AccountID))
	path = strings.ReplaceAll(path, "{id}", fmt.Sprintf("%v", t.Path.ID))
	return path
}

func (t *UpdateSharedTheme) GetQuery() (string, error) {
	return "", nil
}

func (t *UpdateSharedTheme) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *UpdateSharedTheme) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *UpdateSharedTheme) HasErrors() error {
	errs := []string{}
	if t.Path.AccountID == "" {
		errs = append(errs, "'Path.AccountID' is required")
	}
	if t.Path.ID == "" {
		errs = append(errs, "'Path.ID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *UpdateSharedTheme) Do(c *canvasapi.Canvas) (*models.SharedBrandConfig, error) {
	response, err := c.SendRequest(t)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		return nil, err
	}
	ret := models.SharedBrandConfig{}
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return nil, err
	}

	return &ret, nil
}
