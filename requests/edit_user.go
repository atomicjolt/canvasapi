package requests

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/google/go-querystring/query"

	"github.com/atomicjolt/canvasapi"
	"github.com/atomicjolt/canvasapi/models"
)

// EditUser Modify an existing user. To modify a user's login, see the documentation for logins.
// https://canvas.instructure.com/doc/api/users.html
//
// Path Parameters:
// # ID (Required) ID
//
// Form Parameters:
// # User (Optional) The full name of the user. This name will be used by teacher for grading.
// # User (Optional) User's name as it will be displayed in discussions, messages, and comments.
// # User (Optional) User's name as used to sort alphabetically in lists.
// # User (Optional) The time zone for the user. Allowed time zones are
//    {http://www.iana.org/time-zones IANA time zones} or friendlier
//    {http://api.rubyonrails.org/classes/ActiveSupport/TimeZone.html Ruby on Rails time zones}.
// # User (Optional) The default email address of the user.
// # User (Optional) The user's preferred language, from the list of languages Canvas supports.
//    This is in RFC-5646 format.
// # User (Optional) A unique representation of the avatar record to assign as the user's
//    current avatar. This token can be obtained from the user avatars endpoint.
//    This supersedes the user [avatar] [url] argument, and if both are included
//    the url will be ignored. Note: this is an internal representation and is
//    subject to change without notice. It should be consumed with this api
//    endpoint and used in the user update endpoint, and should not be
//    constructed by the client.
// # User (Optional) To set the user's avatar to point to an external url, do not include a
//    token and instead pass the url here. Warning: For maximum compatibility,
//    please use 128 px square images.
// # User (Optional) Sets a title on the user profile. (See {api:ProfileController#settings Get user profile}.)
//    Profiles must be enabled on the root account.
// # User (Optional) Sets a bio on the user profile. (See {api:ProfileController#settings Get user profile}.)
//    Profiles must be enabled on the root account.
// # User (Optional) Sets pronouns on the user profile.
//    Passing an empty string will empty the user's pronouns
//    Only Available Pronouns set on the root account are allowed
//    Adding and changing pronouns must be enabled on the root account.
//
type EditUser struct {
	Path struct {
		ID string `json:"id"` //  (Required)
	} `json:"path"`

	Form struct {
		User struct {
			Name         string `json:"name"`          //  (Optional)
			ShortName    string `json:"short_name"`    //  (Optional)
			SortableName string `json:"sortable_name"` //  (Optional)
			TimeZone     string `json:"time_zone"`     //  (Optional)
			Email        string `json:"email"`         //  (Optional)
			Locale       string `json:"locale"`        //  (Optional)
			Avatar       struct {
				Token string `json:"token"` //  (Optional)
				Url   string `json:"url"`   //  (Optional)
			} `json:"avatar"`

			Title    string `json:"title"`    //  (Optional)
			Bio      string `json:"bio"`      //  (Optional)
			Pronouns string `json:"pronouns"` //  (Optional)
		} `json:"user"`
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

func (t *EditUser) GetBody() (string, error) {
	v, err := query.Values(t.Form)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%v", v.Encode()), nil
}

func (t *EditUser) HasErrors() error {
	errs := []string{}
	if t.Path.ID == "" {
		errs = append(errs, "'ID' is required")
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
