package requests

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/atomicjolt/canvasapi"
	"github.com/atomicjolt/canvasapi/models"
)

// ListMigrationSystemsAccounts Lists the currently available migration types. These values may change.
// https://canvas.instructure.com/doc/api/content_migrations.html
//
// Path Parameters:
// # AccountID (Required) ID
//
type ListMigrationSystemsAccounts struct {
	Path struct {
		AccountID string `json:"account_id"` //  (Required)
	} `json:"path"`
}

func (t *ListMigrationSystemsAccounts) GetMethod() string {
	return "GET"
}

func (t *ListMigrationSystemsAccounts) GetURLPath() string {
	path := "accounts/{account_id}/content_migrations/migrators"
	path = strings.ReplaceAll(path, "{account_id}", fmt.Sprintf("%v", t.Path.AccountID))
	return path
}

func (t *ListMigrationSystemsAccounts) GetQuery() (string, error) {
	return "", nil
}

func (t *ListMigrationSystemsAccounts) GetBody() (string, error) {
	return "", nil
}

func (t *ListMigrationSystemsAccounts) HasErrors() error {
	errs := []string{}
	if t.Path.AccountID == "" {
		errs = append(errs, "'AccountID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *ListMigrationSystemsAccounts) Do(c *canvasapi.Canvas) ([]*models.Migrator, error) {
	response, err := c.SendRequest(t)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		return nil, err
	}
	ret := []*models.Migrator{}
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return nil, err
	}

	return ret, nil
}
