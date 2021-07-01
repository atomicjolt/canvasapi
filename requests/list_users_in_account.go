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

// ListUsersInAccount A paginated list of of users associated with this account.
//
//  @example_request
//    curl https://<canvas>/api/v1/accounts/self/users?search_term=<search value> \
//       -X GET \
//       -H 'Authorization: Bearer <token>'
// https://canvas.instructure.com/doc/api/users.html
//
// Path Parameters:
// # Path.AccountID (Required) ID
//
// Query Parameters:
// # Query.SearchTerm (Optional) The partial name or full ID of the users to match and return in the
//    results list. Must be at least 3 characters.
//
//    Note that the API will prefer matching on canonical user ID if the ID has
//    a numeric form. It will only search against other fields if non-numeric
//    in form, or if the numeric value doesn't yield any matches. Queries by
//    administrative users will search on SIS ID, login ID, name, or email
//    address
// # Query.EnrollmentType (Optional) When set, only return users enrolled with the specified course-level base role.
//    This can be a base role type of 'student', 'teacher',
//    'ta', 'observer', or 'designer'.
// # Query.Sort (Optional) . Must be one of username, email, sis_id, last_loginThe column to sort results by.
// # Query.Order (Optional) . Must be one of asc, descThe order to sort the given column by.
//
type ListUsersInAccount struct {
	Path struct {
		AccountID string `json:"account_id" url:"account_id,omitempty"` //  (Required)
	} `json:"path"`

	Query struct {
		SearchTerm     string `json:"search_term" url:"search_term,omitempty"`         //  (Optional)
		EnrollmentType string `json:"enrollment_type" url:"enrollment_type,omitempty"` //  (Optional)
		Sort           string `json:"sort" url:"sort,omitempty"`                       //  (Optional) . Must be one of username, email, sis_id, last_login
		Order          string `json:"order" url:"order,omitempty"`                     //  (Optional) . Must be one of asc, desc
	} `json:"query"`
}

func (t *ListUsersInAccount) GetMethod() string {
	return "GET"
}

func (t *ListUsersInAccount) GetURLPath() string {
	path := "accounts/{account_id}/users"
	path = strings.ReplaceAll(path, "{account_id}", fmt.Sprintf("%v", t.Path.AccountID))
	return path
}

func (t *ListUsersInAccount) GetQuery() (string, error) {
	v, err := query.Values(t.Query)
	if err != nil {
		return "", err
	}
	return v.Encode(), nil
}

func (t *ListUsersInAccount) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *ListUsersInAccount) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *ListUsersInAccount) HasErrors() error {
	errs := []string{}
	if t.Path.AccountID == "" {
		errs = append(errs, "'Path.AccountID' is required")
	}
	if t.Query.Sort != "" && !string_utils.Include([]string{"username", "email", "sis_id", "last_login"}, t.Query.Sort) {
		errs = append(errs, "Sort must be one of username, email, sis_id, last_login")
	}
	if t.Query.Order != "" && !string_utils.Include([]string{"asc", "desc"}, t.Query.Order) {
		errs = append(errs, "Order must be one of asc, desc")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *ListUsersInAccount) Do(c *canvasapi.Canvas) ([]*models.User, *canvasapi.PagedResource, error) {
	response, err := c.SendRequest(t)
	if err != nil {
		return nil, nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		return nil, nil, err
	}
	ret := []*models.User{}
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return nil, nil, err
	}

	pagedResource, err := canvasapi.ExtractPagedResource(response.Header)
	if err != nil {
		return nil, nil, err
	}

	return ret, pagedResource, nil
}
