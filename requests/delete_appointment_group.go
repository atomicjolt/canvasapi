package requests

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/google/go-querystring/query"

	"github.com/atomicjolt/canvasapi"
)

// DeleteAppointmentGroup Delete an appointment group (and associated time slots and reservations)
// and return the deleted group
// https://canvas.instructure.com/doc/api/appointment_groups.html
//
// Path Parameters:
// # Path.ID (Required) ID
//
// Query Parameters:
// # Query.CancelReason (Optional) Reason for deleting/canceling the appointment group.
//
type DeleteAppointmentGroup struct {
	Path struct {
		ID string `json:"id" url:"id,omitempty"` //  (Required)
	} `json:"path"`

	Query struct {
		CancelReason string `json:"cancel_reason" url:"cancel_reason,omitempty"` //  (Optional)
	} `json:"query"`
}

func (t *DeleteAppointmentGroup) GetMethod() string {
	return "DELETE"
}

func (t *DeleteAppointmentGroup) GetURLPath() string {
	path := "appointment_groups/{id}"
	path = strings.ReplaceAll(path, "{id}", fmt.Sprintf("%v", t.Path.ID))
	return path
}

func (t *DeleteAppointmentGroup) GetQuery() (string, error) {
	v, err := query.Values(t.Query)
	if err != nil {
		return "", err
	}
	return v.Encode(), nil
}

func (t *DeleteAppointmentGroup) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *DeleteAppointmentGroup) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *DeleteAppointmentGroup) HasErrors() error {
	errs := []string{}
	if t.Path.ID == "" {
		errs = append(errs, "'Path.ID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *DeleteAppointmentGroup) Do(c *canvasapi.Canvas) error {
	_, err := c.SendRequest(t)
	if err != nil {
		return err
	}

	return nil
}
