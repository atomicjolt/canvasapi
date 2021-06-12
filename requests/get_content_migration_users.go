package requests

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/atomicjolt/canvasapi"
	"github.com/atomicjolt/canvasapi/models"
)

// GetContentMigrationUsers Returns data on an individual content migration
// https://canvas.instructure.com/doc/api/content_migrations.html
//
// Path Parameters:
// # UserID (Required) ID
// # ID (Required) ID
//
type GetContentMigrationUsers struct {
	Path struct {
		UserID string `json:"user_id"` //  (Required)
		ID     string `json:"id"`      //  (Required)
	} `json:"path"`
}

func (t *GetContentMigrationUsers) GetMethod() string {
	return "GET"
}

func (t *GetContentMigrationUsers) GetURLPath() string {
	path := "users/{user_id}/content_migrations/{id}"
	path = strings.ReplaceAll(path, "{user_id}", fmt.Sprintf("%v", t.Path.UserID))
	path = strings.ReplaceAll(path, "{id}", fmt.Sprintf("%v", t.Path.ID))
	return path
}

func (t *GetContentMigrationUsers) GetQuery() (string, error) {
	return "", nil
}

func (t *GetContentMigrationUsers) GetBody() (string, error) {
	return "", nil
}

func (t *GetContentMigrationUsers) HasErrors() error {
	errs := []string{}
	if t.Path.UserID == "" {
		errs = append(errs, "'UserID' is required")
	}
	if t.Path.ID == "" {
		errs = append(errs, "'ID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *GetContentMigrationUsers) Do(c *canvasapi.Canvas) (*models.ContentMigration, error) {
	response, err := c.SendRequest(t)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		return nil, err
	}
	ret := models.ContentMigration{}
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return nil, err
	}

	return &ret, nil
}
