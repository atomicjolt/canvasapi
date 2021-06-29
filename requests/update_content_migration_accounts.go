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

// UpdateContentMigrationAccounts Update a content migration. Takes same arguments as {api:ContentMigrationsController#create create} except that you
// can't change the migration type. However, changing most settings after the
// migration process has started will not do anything. Generally updating the
// content migration will be used when there is a file upload problem, or when
// importing content selectively. If the first upload has a problem you can
// supply new _pre_attachment_ values to start the process again.
// https://canvas.instructure.com/doc/api/content_migrations.html
//
// Path Parameters:
// # Path.AccountID (Required) ID
// # Path.ID (Required) ID
//
type UpdateContentMigrationAccounts struct {
	Path struct {
		AccountID string `json:"account_id" url:"account_id,omitempty"` //  (Required)
		ID        string `json:"id" url:"id,omitempty"`                 //  (Required)
	} `json:"path"`
}

func (t *UpdateContentMigrationAccounts) GetMethod() string {
	return "PUT"
}

func (t *UpdateContentMigrationAccounts) GetURLPath() string {
	path := "accounts/{account_id}/content_migrations/{id}"
	path = strings.ReplaceAll(path, "{account_id}", fmt.Sprintf("%v", t.Path.AccountID))
	path = strings.ReplaceAll(path, "{id}", fmt.Sprintf("%v", t.Path.ID))
	return path
}

func (t *UpdateContentMigrationAccounts) GetQuery() (string, error) {
	return "", nil
}

func (t *UpdateContentMigrationAccounts) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *UpdateContentMigrationAccounts) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *UpdateContentMigrationAccounts) HasErrors() error {
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

func (t *UpdateContentMigrationAccounts) Do(c *canvasapi.Canvas) (*models.ContentMigration, error) {
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
