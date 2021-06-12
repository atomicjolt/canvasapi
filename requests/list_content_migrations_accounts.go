package requests

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/atomicjolt/canvasapi"
	"github.com/atomicjolt/canvasapi/models"
)

// ListContentMigrationsAccounts Returns paginated content migrations
// https://canvas.instructure.com/doc/api/content_migrations.html
//
// Path Parameters:
// # AccountID (Required) ID
//
type ListContentMigrationsAccounts struct {
	Path struct {
		AccountID string `json:"account_id"` //  (Required)
	} `json:"path"`
}

func (t *ListContentMigrationsAccounts) GetMethod() string {
	return "GET"
}

func (t *ListContentMigrationsAccounts) GetURLPath() string {
	path := "accounts/{account_id}/content_migrations"
	path = strings.ReplaceAll(path, "{account_id}", fmt.Sprintf("%v", t.Path.AccountID))
	return path
}

func (t *ListContentMigrationsAccounts) GetQuery() (string, error) {
	return "", nil
}

func (t *ListContentMigrationsAccounts) GetBody() (string, error) {
	return "", nil
}

func (t *ListContentMigrationsAccounts) HasErrors() error {
	errs := []string{}
	if t.Path.AccountID == "" {
		errs = append(errs, "'AccountID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *ListContentMigrationsAccounts) Do(c *canvasapi.Canvas) ([]*models.ContentMigration, error) {
	response, err := c.SendRequest(t)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		return nil, err
	}
	ret := []*models.ContentMigration{}
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return nil, err
	}

	return ret, nil
}
