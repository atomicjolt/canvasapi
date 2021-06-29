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

// EditUser Modify an existing user. To modify a user's login, see the documentation for logins.
// https://canvas.instructure.com/doc/api/users.html
//
// Path Parameters:
// # Path.ID (Required) ID
//
// Form Parameters:
// # Form.User.Name (Optional) The full name of the user. This name will be used by teacher for grading.
// # Form.User.ShortName (Optional) User's name as it will be displayed in discussions, messages, and comments.
// # Form.User.SortableName (Optional) User's name as used to sort alphabetically in lists.
// # Form.User.TimeZone (Optional) The time zone for the user. Allowed time zones are
//    {http://www.iana.org/time-zones IANA time zones} or friendlier
//    {http://api.rubyonrails.org/classes/ActiveSupport/TimeZone.html Ruby on Rails time zones}.
// # Form.User.Email (Optional) The default email address of the user.
// # Form.User.Locale (Optional) The user's preferred language, from the list of languages Canvas supports.
//    This is in RFC-5646 format.
// # Form.User.Avatar.Token (Optional) A unique representation of the avatar record to assign as the user's
//    current avatar. This token can be obtained from the user avatars endpoint.
//    This supersedes the user [avatar] [url] argument, and if both are included
//    the url will be ignored. Note: this is an internal representation and is
//    subject to change without notice. It should be consumed with this api
//    endpoint and used in the user update endpoint, and should not be
//    constructed by the client.
// # Form.User.Avatar.Url (Optional) To set the user's avatar to point to an external url, do not include a
//    token and instead pass the url here. Warning: For maximum compatibility,
//    please use 128 px square images.
// # Form.User.Title (Optional) Sets a title on the user profile. (See {api:ProfileController#settings Get user profile}.)
//    Profiles must be enabled on the root account.
// # Form.User.Bio (Optional) Sets a bio on the user profile. (See {api:ProfileController#settings Get user profile}.)
//    Profiles must be enabled on the root account.
// # Form.User.Pronouns (Optional) Sets pronouns on the user profile.
//    Passing an empty string will empty the user's pronouns
//    Only Available Pronouns set on the root account are allowed
//    Adding and changing pronouns must be enabled on the root account.
//
type EditUser struct {
	Path struct {
		ID string `json:"id" url:"id,omitempty"` //  (Required)
	} `json:"path"`

	Form struct {
		User struct {
			Name         string `json:"name" url:"name,omitempty"`                   //  (Optional)
			ShortName    string `json:"short_name" url:"short_name,omitempty"`       //  (Optional)
			SortableName string `json:"sortable_name" url:"sortable_name,omitempty"` //  (Optional)
			TimeZone     string `json:"time_zone" url:"time_zone,omitempty"`         //  (Optional)
			Email        string `json:"email" url:"email,omitempty"`                 //  (Optional)
			Locale       string `json:"locale" url:"locale,omitempty"`               //  (Optional)
			Avatar       struct {
				Token string `json:"token" url:"token,omitempty"` //  (Optional)
				Url   string `json:"url" url:"url,omitempty"`     //  (Optional)
			} `json:"avatar" url:"avatar,omitempty"`

			Title    string `json:"title" url:"title,omitempty"`       //  (Optional)
			Bio      string `json:"bio" url:"bio,omitempty"`           //  (Optional)
			Pronouns string `json:"pronouns" url:"pronouns,omitempty"` //  (Optional)
		} `json:"user" url:"user,omitempty"`
	} `json:"form"`
}

func (t *EditUser) GetMethod() string {
	return "PUT"
}

func (t *EditUser) GetURLPath() string {
	path := "users/{id}"
	path = strings.ReplaceAll(path, "{id}", fmt.Sprintf("%v", t.Path.ID))
	return path
}

func (t *EditUser) GetQuery() (string, error) {
	return "", nil
}

func (t *EditUser) GetBody() (url.Values, error) {
	return query.Values(t.Form)
}

func (t *EditUser) GetJSON() ([]byte, error) {
	j, err := json.Marshal(t.Form)
	if err != nil {
		return nil, nil
	}
	return j, nil
}

func (t *EditUser) HasErrors() error {
	errs := []string{}
	if t.Path.ID == "" {
		errs = append(errs, "'Path.ID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *EditUser) Do(c *canvasapi.Canvas) (*models.User, error) {
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
