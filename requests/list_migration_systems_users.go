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

// ListMigrationSystemsUsers Lists the currently available migration types. These values may change.
// https://canvas.instructure.com/doc/api/content_migrations.html
//
// Path Parameters:
// # Path.UserID (Required) ID
//
type ListMigrationSystemsUsers struct {
	Path struct {
		UserID string `json:"user_id" url:"user_id,omitempty"` //  (Required)
	} `json:"path"`
}

func (t *ListMigrationSystemsUsers) GetMethod() string {
	return "GET"
}

func (t *ListMigrationSystemsUsers) GetURLPath() string {
	path := "users/{user_id}/content_migrations/migrators"
	path = strings.ReplaceAll(path, "{user_id}", fmt.Sprintf("%v", t.Path.UserID))
	return path
}

func (t *ListMigrationSystemsUsers) GetQuery() (string, error) {
	return "", nil
}

func (t *ListMigrationSystemsUsers) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *ListMigrationSystemsUsers) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *ListMigrationSystemsUsers) HasErrors() error {
	errs := []string{}
	if t.Path.UserID == "" {
		errs = append(errs, "'Path.UserID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *ListMigrationSystemsUsers) Do(c *canvasapi.Canvas) ([]*models.Migrator, error) {
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
