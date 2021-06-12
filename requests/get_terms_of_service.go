package requests

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/atomicjolt/canvasapi"
	"github.com/atomicjolt/canvasapi/models"
)

// GetTermsOfService Returns the terms of service for that account
// https://canvas.instructure.com/doc/api/accounts.html
//
// Path Parameters:
// # AccountID (Required) ID
//
type GetTermsOfService struct {
	Path struct {
		AccountID string `json:"account_id"` //  (Required)
	} `json:"path"`
}

func (t *GetTermsOfService) GetMethod() string {
	return "GET"
}

func (t *GetTermsOfService) GetURLPath() string {
	path := "accounts/{account_id}/terms_of_service"
	path = strings.ReplaceAll(path, "{account_id}", fmt.Sprintf("%v", t.Path.AccountID))
	return path
}

func (t *GetTermsOfService) GetQuery() (string, error) {
	return "", nil
}

func (t *GetTermsOfService) GetBody() (string, error) {
	return "", nil
}

func (t *GetTermsOfService) HasErrors() error {
	errs := []string{}
	if t.Path.AccountID == "" {
		errs = append(errs, "'AccountID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *GetTermsOfService) Do(c *canvasapi.Canvas) (*models.TermsOfService, error) {
	response, err := c.SendRequest(t)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		return nil, err
	}
	ret := models.TermsOfService{}
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return nil, err
	}

	return &ret, nil
}
