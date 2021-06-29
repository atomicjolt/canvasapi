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

// ActivateRole Re-activates an inactive role (allowing it to be assigned to new users)
// https://canvas.instructure.com/doc/api/roles.html
//
// Path Parameters:
// # Path.AccountID (Required) ID
// # Path.ID (Required) ID
//
// Form Parameters:
// # Form.RoleID (Required) The unique identifier for the role
// # Form.Role (Optional) The name for the role
//
type ActivateRole struct {
	Path struct {
		AccountID string `json:"account_id" url:"account_id,omitempty"` //  (Required)
		ID        string `json:"id" url:"id,omitempty"`                 //  (Required)
	} `json:"path"`

	Form struct {
		RoleID int64  `json:"role_id" url:"role_id,omitempty"` //  (Required)
		Role   string `json:"role" url:"role,omitempty"`       //  (Optional)
	} `json:"form"`
}

func (t *ActivateRole) GetMethod() string {
	return "POST"
}

func (t *ActivateRole) GetURLPath() string {
	path := "accounts/{account_id}/roles/{id}/activate"
	path = strings.ReplaceAll(path, "{account_id}", fmt.Sprintf("%v", t.Path.AccountID))
	path = strings.ReplaceAll(path, "{id}", fmt.Sprintf("%v", t.Path.ID))
	return path
}

func (t *ActivateRole) GetQuery() (string, error) {
	return "", nil
}

func (t *ActivateRole) GetBody() (url.Values, error) {
	return query.Values(t.Form)
}

func (t *ActivateRole) GetJSON() ([]byte, error) {
	j, err := json.Marshal(t.Form)
	if err != nil {
		return nil, nil
	}
	return j, nil
}

func (t *ActivateRole) HasErrors() error {
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

func (t *ActivateRole) Do(c *canvasapi.Canvas) (*models.Role, error) {
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
