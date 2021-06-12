package requests

import (
	"fmt"
	"strings"

	"github.com/google/go-querystring/query"

	"github.com/atomicjolt/canvasapi"
)

// UpdateCustomColor Updates a custom color for a user for a given context.  This allows
// colors for the calendar and elsewhere to be customized on a user basis.
//
// The asset string parameter should be in the format 'context_id', for example
// 'course_42'
// https://canvas.instructure.com/doc/api/users.html
//
// Path Parameters:
// # ID (Required) ID
// # AssetString (Required) ID
//
// Form Parameters:
// # Hexcode (Optional) The hexcode of the color to set for the context, if you choose to pass the
//    hexcode as a query parameter rather than in the request body you should
//    NOT include the '#' unless you escape it first.
//
type UpdateCustomColor struct {
	Path struct {
		ID          string `json:"id"`           //  (Required)
		AssetString string `json:"asset_string"` //  (Required)
	} `json:"path"`

	Form struct {
		Hexcode string `json:"hexcode"` //  (Optional)
	} `json:"form"`
}

func (t *UpdateCustomColor) GetMethod() string {
	return "PUT"
}

func (t *UpdateCustomColor) GetURLPath() string {
	path := "users/{id}/colors/{asset_string}"
	path = strings.ReplaceAll(path, "{id}", fmt.Sprintf("%v", t.Path.ID))
	path = strings.ReplaceAll(path, "{asset_string}", fmt.Sprintf("%v", t.Path.AssetString))
	return path
}

func (t *UpdateCustomColor) GetQuery() (string, error) {
	return "", nil
}

func (t *UpdateCustomColor) GetBody() (string, error) {
	v, err := query.Values(t.Form)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%v", v.Encode()), nil
}

func (t *UpdateCustomColor) HasErrors() error {
	errs := []string{}
	if t.Path.ID == "" {
		errs = append(errs, "'ID' is required")
	}
	if t.Path.AssetString == "" {
		errs = append(errs, "'AssetString' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *UpdateCustomColor) Do(c *canvasapi.Canvas) error {
	_, err := c.SendRequest(t)
	if err != nil {
		return err
	}

	return nil
}
