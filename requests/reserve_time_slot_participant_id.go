package requests

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strings"

	"github.com/google/go-querystring/query"

	"github.com/atomicjolt/canvasapi"
)

// ReserveTimeSlotParticipantID Reserves a particular time slot and return the new reservation
// https://canvas.instructure.com/doc/api/calendar_events.html
//
// Path Parameters:
// # ID (Required) ID
// # ParticipantID (Required) User or group id for whom you are making the reservation (depends on the
//    participant type). Defaults to the current user (or user's candidate group).
//
// Form Parameters:
// # Comments (Optional) Comments to associate with this reservation
// # CancelExisting (Optional) Defaults to false. If true, cancel any previous reservation(s) for this
//    participant and appointment group.
//
type ReserveTimeSlotParticipantID struct {
	Path struct {
		ID            string `json:"id" url:"id,omitempty"`                         //  (Required)
		ParticipantID string `json:"participant_id" url:"participant_id,omitempty"` //  (Required)
	} `json:"path"`

	Form struct {
		Comments       string `json:"comments" url:"comments,omitempty"`               //  (Optional)
		CancelExisting bool   `json:"cancel_existing" url:"cancel_existing,omitempty"` //  (Optional)
	} `json:"form"`
}

func (t *ReserveTimeSlotParticipantID) GetMethod() string {
	return "POST"
}

func (t *ReserveTimeSlotParticipantID) GetURLPath() string {
	path := "calendar_events/{id}/reservations/{participant_id}"
	path = strings.ReplaceAll(path, "{id}", fmt.Sprintf("%v", t.Path.ID))
	path = strings.ReplaceAll(path, "{participant_id}", fmt.Sprintf("%v", t.Path.ParticipantID))
	return path
}

func (t *ReserveTimeSlotParticipantID) GetQuery() (string, error) {
	return "", nil
}

func (t *ReserveTimeSlotParticipantID) GetBody() (url.Values, error) {
	return query.Values(t.Form)
}

func (t *ReserveTimeSlotParticipantID) GetJSON() ([]byte, error) {
	j, err := json.Marshal(t.Form)
	if err != nil {
		return nil, nil
	}
	return j, nil
}

func (t *ReserveTimeSlotParticipantID) HasErrors() error {
	errs := []string{}
	if t.Path.ID == "" {
		errs = append(errs, "'ID' is required")
	}
	if t.Path.ParticipantID == "" {
		errs = append(errs, "'ParticipantID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *ReserveTimeSlotParticipantID) Do(c *canvasapi.Canvas) error {
	_, err := c.SendRequest(t)
	if err != nil {
		return err
	}

	return nil
}
