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
	"github.com/atomicjolt/string_utils"
)

// ShowUserDetails Shows details for user.
//
// Also includes an attribute "permissions", a non-comprehensive list of permissions for the user.
// Example:
//   !!!javascript
//   "permissions": {
//    "can_update_name": true, // Whether the user can update their name.
//    "can_update_avatar": false, // Whether the user can update their avatar.
//    "limit_parent_app_web_access": false // Whether the user can interact with Canvas web from the Canvas Parent app.
//   }
// https://canvas.instructure.com/doc/api/users.html
//
// Path Parameters:
// # Path.ID (Required) ID
//
// Query Parameters:
// # Query.Include (Optional) . Must be one of uuid, last_loginArray of additional information to include on the user record.
//    "locale", "avatar_url", "permissions", "email", and "effective_locale"
//    will always be returned
//
type ShowUserDetails struct {
	Path struct {
		ID string `json:"id" url:"id,omitempty"` //  (Required)
	} `json:"path"`

	Query struct {
		Include []string `json:"include" url:"include,omitempty"` //  (Optional) . Must be one of uuid, last_login
	} `json:"query"`
}

func (t *ShowUserDetails) GetMethod() string {
	return "GET"
}

func (t *ShowUserDetails) GetURLPath() string {
	path := "users/{id}"
	path = strings.ReplaceAll(path, "{id}", fmt.Sprintf("%v", t.Path.ID))
	return path
}

func (t *ShowUserDetails) GetQuery() (string, error) {
	v, err := query.Values(t.Query)
	if err != nil {
		return "", err
	}
	return v.Encode(), nil
}

func (t *ShowUserDetails) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *ShowUserDetails) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *ShowUserDetails) HasErrors() error {
	errs := []string{}
	if t.Path.ID == "" {
		errs = append(errs, "'Path.ID' is required")
	}
	for _, v := range t.Query.Include {
		if v != "" && !string_utils.Include([]string{"uuid", "last_login"}, v) {
			errs = append(errs, "Include must be one of uuid, last_login")
		}
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *ShowUserDetails) Do(c *canvasapi.Canvas) (*models.User, error) {
	response, err := c.SendRequest(t)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		return nil, err
	}
	ret := models.User{}
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return nil, err
	}

	return &ret, nil
}
