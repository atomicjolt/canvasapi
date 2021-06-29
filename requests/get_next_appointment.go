package requests

import (
	"encoding/json"
	"io/ioutil"
	"net/url"

	"github.com/google/go-querystring/query"

	"github.com/atomicjolt/canvasapi"
	"github.com/atomicjolt/canvasapi/models"
)

// GetNextAppointment Return the next appointment available to sign up for. The appointment
// is returned in a one-element array. If no future appointments are
// available, an empty array is returned.
// https://canvas.instructure.com/doc/api/appointment_groups.html
//
// Query Parameters:
// # Query.AppointmentGroupIDs (Optional) List of ids of appointment groups to search.
//
type GetNextAppointment struct {
	Query struct {
		AppointmentGroupIDs []string `json:"appointment_group_ids" url:"appointment_group_ids,omitempty"` //  (Optional)
	} `json:"query"`
}

func (t *GetNextAppointment) GetMethod() string {
	return "GET"
}

func (t *GetNextAppointment) GetURLPath() string {
	return ""
}

func (t *GetNextAppointment) GetQuery() (string, error) {
	v, err := query.Values(t.Query)
	if err != nil {
		return "", err
	}
	return v.Encode(), nil
}

func (t *GetNextAppointment) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *GetNextAppointment) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *GetNextAppointment) HasErrors() error {
	return nil
}

func (t *GetNextAppointment) Do(c *canvasapi.Canvas) ([]*models.CalendarEvent, error) {
	response, err := c.SendRequest(t)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		return nil, err
	}
	ret := []*models.CalendarEvent{}
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return nil, err
	}

	return ret, nil
}
