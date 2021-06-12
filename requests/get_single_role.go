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

// GetSingleRole Retrieve information about a single role
// https://canvas.instructure.com/doc/api/roles.html
//
// Path Parameters:
// # ID (Required) ID
// # AccountID (Required) The id of the account containing the role
//
// Query Parameters:
// # RoleID (Required) The unique identifier for the role
// # Role (Optional) The name for the role
//
type GetSingleRole struct {
	Path struct {
		ID        string `json:"id"`         //  (Required)
		AccountID string `json:"account_id"` //  (Required)
	} `json:"path"`

	Query struct {
		RoleID int64  `json:"role_id"` //  (Required)
		Role   string `json:"role"`    //  (Optional)
	} `json:"query"`
}

func (t *GetSingleRole) GetMethod() string {
	return "GET"
}

func (t *GetSingleRole) GetURLPath() string {
	path := "accounts/{account_id}/roles/{id}"
	path = strings.ReplaceAll(path, "{id}", fmt.Sprintf("%v", t.Path.ID))
	path = strings.ReplaceAll(path, "{account_id}", fmt.Sprintf("%v", t.Path.AccountID))
	return path
}

func (t *GetSingleRole) GetQuery() (string, error) {
	v, err := query.Values(t.Query)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("?%v", v.Encode()), nil
}

func (t *GetSingleRole) GetBody() (string, error) {
	return "", nil
}

func (t *GetSingleRole) HasErrors() error {
	errs := []string{}
	if t.Path.ID == "" {
		errs = append(errs, "'ID' is required")
	}
	if t.Path.AccountID == "" {
		errs = append(errs, "'AccountID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *GetSingleRole) Do(c *canvasapi.Canvas) (*models.Role, error) {
	response, err := c.SendRequest(t)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		return nil, err
	}
	ret := models.Role{}
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return nil, err
	}

	return &ret, nil
}
