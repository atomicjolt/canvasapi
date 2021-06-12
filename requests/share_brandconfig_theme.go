package requests

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/google/go-querystring/query"

	"github.com/atomicjolt/canvasapi"
	"github.com/atomicjolt/canvasapi/models"
)

// ShareBrandconfigTheme Create a SharedBrandConfig, which will give the given brand_config a name
// and make it available to other users of this account.
// https://canvas.instructure.com/doc/api/shared_brand_configs.html
//
// Path Parameters:
// # AccountID (Required) ID
//
// Form Parameters:
// # SharedBrandConfig (Required) Name to share this BrandConfig (theme) as.
// # SharedBrandConfig (Required) MD5 of brand_config to share
//
type ShareBrandconfigTheme struct {
	Path struct {
		AccountID string `json:"account_id"` //  (Required)
	} `json:"path"`

	Form struct {
		SharedBrandConfig struct {
			Name           string `json:"name"`             //  (Required)
			BrandConfigMd5 string `json:"brand_config_md5"` //  (Required)
		} `json:"shared_brand_config"`
	} `json:"form"`
}

func (t *ShareBrandconfigTheme) GetMethod() string {
	return "POST"
}

func (t *ShareBrandconfigTheme) GetURLPath() string {
	path := "accounts/{account_id}/shared_brand_configs"
	path = strings.ReplaceAll(path, "{account_id}", fmt.Sprintf("%v", t.Path.AccountID))
	return path
}

func (t *ShareBrandconfigTheme) GetQuery() (string, error) {
	return "", nil
}

func (t *ShareBrandconfigTheme) GetBody() (string, error) {
	v, err := query.Values(t.Form)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%v", v.Encode()), nil
}

func (t *ShareBrandconfigTheme) HasErrors() error {
	errs := []string{}
	if t.Path.AccountID == "" {
		errs = append(errs, "'AccountID' is required")
	}
	if t.Form.SharedBrandConfig.Name == "" {
		errs = append(errs, "'SharedBrandConfig' is required")
	}
	if t.Form.SharedBrandConfig.BrandConfigMd5 == "" {
		errs = append(errs, "'SharedBrandConfig' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *ShareBrandconfigTheme) Do(c *canvasapi.Canvas) (*models.SharedBrandConfig, error) {
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
