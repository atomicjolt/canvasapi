package requests

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/google/go-querystring/query"

	"github.com/atomicjolt/canvasapi"
)

// DeleteCustomData Delete custom user data.
//
// Arbitrary JSON data can be stored for a User.  This API call
// deletes that data for a given scope.  Without a scope, all custom_data is deleted.
// See {api:UsersController#set_custom_data Store Custom Data} for details and
// examples of storage and retrieval.
//
// As an example, we'll store some data, then delete a subset of it.
//
// Example {api:UsersController#set_custom_data PUT} with valid JSON data:
//   curl 'https://<canvas>/api/v1/users/<user_id>/custom_data' \
//     -X PUT \
//     -F 'ns=com.my-organization.canvas-app' \
//     -F 'data[fruit][apple]=so tasty' \
//     -F 'data[fruit][kiwi]=a bit sour' \
//     -F 'data[veggies][root][onion]=tear-jerking' \
//     -H 'Authorization: Bearer <token>'
//
// Response:
//   !!!javascript
//   {
//     "data": {
//       "fruit": {
//         "apple": "so tasty",
//         "kiwi": "a bit sour"
//       },
//       "veggies": {
//         "root": {
//           "onion": "tear-jerking"
//         }
//       }
//     }
//   }
//
// Example DELETE:
//   curl 'https://<canvas>/api/v1/users/<user_id>/custom_data/fruit/kiwi' \
//     -X DELETE \
//     -F 'ns=com.my-organization.canvas-app' \
//     -H 'Authorization: Bearer <token>'
//
// Response:
//   !!!javascript
//   {
//     "data": "a bit sour"
//   }
//
// Example {api:UsersController#get_custom_data GET} following the above DELETE:
//   curl 'https://<canvas>/api/v1/users/<user_id>/custom_data' \
//     -X GET \
//     -F 'ns=com.my-organization.canvas-app' \
//     -H 'Authorization: Bearer <token>'
//
// Response:
//   !!!javascript
//   {
//     "data": {
//       "fruit": {
//         "apple": "so tasty"
//       },
//       "veggies": {
//         "root": {
//           "onion": "tear-jerking"
//         }
//       }
//     }
//   }
//
// Note that hashes left empty after a DELETE will get removed from the custom_data store.
// For example, following the previous commands, if we delete /custom_data/veggies/root/onion,
// then the entire /custom_data/veggies scope will be removed.
//
// Example DELETE that empties a parent scope:
//   curl 'https://<canvas>/api/v1/users/<user_id>/custom_data/veggies/root/onion' \
//     -X DELETE \
//     -F 'ns=com.my-organization.canvas-app' \
//     -H 'Authorization: Bearer <token>'
//
// Response:
//   !!!javascript
//   {
//     "data": "tear-jerking"
//   }
//
// Example {api:UsersController#get_custom_data GET} following the above DELETE:
//   curl 'https://<canvas>/api/v1/users/<user_id>/custom_data' \
//     -X GET \
//     -F 'ns=com.my-organization.canvas-app' \
//     -H 'Authorization: Bearer <token>'
//
// Response:
//   !!!javascript
//   {
//     "data": {
//       "fruit": {
//         "apple": "so tasty"
//       }
//     }
//   }
//
// On success, this endpoint returns an object containing the data that was deleted.
//
// Responds with status code 400 if the namespace parameter, +ns+, is missing or invalid,
// or if the specified scope does not contain any data.
// https://canvas.instructure.com/doc/api/users.html
//
// Path Parameters:
// # UserID (Required) ID
//
// Query Parameters:
// # Ns (Required) The namespace from which to delete the data.  This should be something other
//    Canvas API apps aren't likely to use, such as a reverse DNS for your organization.
//
type DeleteCustomData struct {
	Path struct {
		UserID string `json:"user_id" url:"user_id,omitempty"` //  (Required)
	} `json:"path"`

	Query struct {
		Ns string `json:"ns" url:"ns,omitempty"` //  (Required)
	} `json:"query"`
}

func (t *DeleteCustomData) GetMethod() string {
	return "DELETE"
}

func (t *DeleteCustomData) GetURLPath() string {
	path := "users/{user_id}/custom_data"
	path = strings.ReplaceAll(path, "{user_id}", fmt.Sprintf("%v", t.Path.UserID))
	return path
}

func (t *DeleteCustomData) GetQuery() (string, error) {
	v, err := query.Values(t.Query)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("?%v", v.Encode()), nil
}

func (t *DeleteCustomData) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *DeleteCustomData) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *DeleteCustomData) HasErrors() error {
	errs := []string{}
	if t.Path.UserID == "" {
		errs = append(errs, "'UserID' is required")
	}
	if t.Query.Ns == "" {
		errs = append(errs, "'Ns' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *DeleteCustomData) Do(c *canvasapi.Canvas) error {
	_, err := c.SendRequest(t)
	if err != nil {
		return err
	}

	return nil
}
