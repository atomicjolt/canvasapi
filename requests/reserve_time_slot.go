package requests

import (
	"fmt"
	"strings"

	"github.com/google/go-querystring/query"

	"github.com/atomicjolt/canvasapi"
)

// ReserveTimeSlot Reserves a particular time slot and return the new reservation
// https://canvas.instructure.com/doc/api/calendar_events.html
//
// Path Parameters:
// # ID (Required) ID
//
// Form Parameters:
// # ParticipantID (Optional) User or group id for whom you are making the reservation (depends on the
//    participant type). Defaults to the current user (or user's candidate group).
// # Comments (Optional) Comments to associate with this reservation
// # CancelExisting (Optional) Defaults to false. If true, cancel any previous reservation(s) for this
//    participant and appointment group.
//
type ReserveTimeSlot struct {
	Path struct {
		ID string `json:"id"` //  (Required)
	} `json:"path"`

	Form struct {
		ParticipantID  string `json:"participant_id"`  //  (Optional)
		Comments       string `json:"comments"`        //  (Optional)
		CancelExisting bool   `json:"cancel_existing"` //  (Optional)
	} `json:"form"`
}

func (t *ReserveTimeSlot) GetMethod() string {
	return "POST"
}

func (t *ReserveTimeSlot) GetURLPath() string {
	path := "calendar_events/{id}/reservations"
	path = strings.ReplaceAll(path, "{id}", fmt.Sprintf("%v", t.Path.ID))
	return path
}

func (t *ReserveTimeSlot) GetQuery() (string, error) {
	return "", nil
}

func (t *ReserveTimeSlot) GetBody() (string, error) {
	v, err := query.Values(t.Form)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%v", v.Encode()), nil
}

func (t *ReserveTimeSlot) HasErrors() error {
	errs := []string{}
	if t.Path.ID == "" {
		errs = append(errs, "'ID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *ReserveTimeSlot) Do(c *canvasapi.Canvas) error {
	_, err := c.SendRequest(t)
	if err != nil {
		return err
	}

	return nil
}
