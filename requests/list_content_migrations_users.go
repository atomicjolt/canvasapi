package requests

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/atomicjolt/canvasapi"
	"github.com/atomicjolt/canvasapi/models"
)

// ListContentMigrationsUsers Returns paginated content migrations
// https://canvas.instructure.com/doc/api/content_migrations.html
//
// Path Parameters:
// # UserID (Required) ID
//
type ListContentMigrationsUsers struct {
	Path struct {
		UserID string `json:"user_id"` //  (Required)
	} `json:"path"`
}

func (t *ListContentMigrationsUsers) GetMethod() string {
	return "GET"
}

func (t *ListContentMigrationsUsers) GetURLPath() string {
	path := "users/{user_id}/content_migrations"
	path = strings.ReplaceAll(path, "{user_id}", fmt.Sprintf("%v", t.Path.UserID))
	return path
}

func (t *ListContentMigrationsUsers) GetQuery() (string, error) {
	return "", nil
}

func (t *ListContentMigrationsUsers) GetBody() (string, error) {
	return "", nil
}

func (t *ListContentMigrationsUsers) HasErrors() error {
	errs := []string{}
	if t.Path.UserID == "" {
		errs = append(errs, "'UserID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *ListContentMigrationsUsers) Do(c *canvasapi.Canvas) ([]*models.ContentMigration, error) {
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
