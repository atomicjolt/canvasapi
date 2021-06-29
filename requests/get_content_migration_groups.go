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

// GetContentMigrationGroups Returns data on an individual content migration
// https://canvas.instructure.com/doc/api/content_migrations.html
//
// Path Parameters:
// # Path.GroupID (Required) ID
// # Path.ID (Required) ID
//
type GetContentMigrationGroups struct {
	Path struct {
		GroupID string `json:"group_id" url:"group_id,omitempty"` //  (Required)
		ID      string `json:"id" url:"id,omitempty"`             //  (Required)
	} `json:"path"`
}

func (t *GetContentMigrationGroups) GetMethod() string {
	return "GET"
}

func (t *GetContentMigrationGroups) GetURLPath() string {
	path := "groups/{group_id}/content_migrations/{id}"
	path = strings.ReplaceAll(path, "{group_id}", fmt.Sprintf("%v", t.Path.GroupID))
	path = strings.ReplaceAll(path, "{id}", fmt.Sprintf("%v", t.Path.ID))
	return path
}

func (t *GetContentMigrationGroups) GetQuery() (string, error) {
	return "", nil
}

func (t *GetContentMigrationGroups) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *GetContentMigrationGroups) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *GetContentMigrationGroups) HasErrors() error {
	errs := []string{}
	if t.Path.GroupID == "" {
		errs = append(errs, "'Path.GroupID' is required")
	}
	if t.Path.ID == "" {
		errs = append(errs, "'Path.ID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *GetContentMigrationGroups) Do(c *canvasapi.Canvas) (*models.ContentMigration, error) {
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
