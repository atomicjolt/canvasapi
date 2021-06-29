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

// UnShareBrandconfigTheme Delete a SharedBrandConfig, which will unshare it so you nor anyone else in
// your account will see it as an option to pick from.
// https://canvas.instructure.com/doc/api/shared_brand_configs.html
//
// Path Parameters:
// # Path.ID (Required) ID
//
type UnShareBrandconfigTheme struct {
	Path struct {
		ID string `json:"id" url:"id,omitempty"` //  (Required)
	} `json:"path"`
}

func (t *UnShareBrandconfigTheme) GetMethod() string {
	return "DELETE"
}

func (t *UnShareBrandconfigTheme) GetURLPath() string {
	path := "shared_brand_configs/{id}"
	path = strings.ReplaceAll(path, "{id}", fmt.Sprintf("%v", t.Path.ID))
	return path
}

func (t *UnShareBrandconfigTheme) GetQuery() (string, error) {
	return "", nil
}

func (t *UnShareBrandconfigTheme) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *UnShareBrandconfigTheme) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *UnShareBrandconfigTheme) HasErrors() error {
	errs := []string{}
	if t.Path.ID == "" {
		errs = append(errs, "'Path.ID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *UnShareBrandconfigTheme) Do(c *canvasapi.Canvas) (*models.SharedBrandConfig, error) {
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
