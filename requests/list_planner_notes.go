package requests

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/url"
	"time"

	"github.com/google/go-querystring/query"

	"github.com/atomicjolt/canvasapi"
	"github.com/atomicjolt/canvasapi/models"
)

// ListPlannerNotes Retrieve the paginated list of planner notes
//
// Retrieve planner note for a user
// https://canvas.instructure.com/doc/api/planner.html
//
// Query Parameters:
// # StartDate (Optional) Only return notes with todo dates since the start_date (inclusive).
//    No default. The value should be formatted as: yyyy-mm-dd or
//    ISO 8601 YYYY-MM-DDTHH:MM:SSZ.
// # EndDate (Optional) Only return notes with todo dates before the end_date (inclusive).
//    No default. The value should be formatted as: yyyy-mm-dd or
//    ISO 8601 YYYY-MM-DDTHH:MM:SSZ.
//    If end_date and start_date are both specified and equivalent,
//    then only notes with todo dates on that day are returned.
// # ContextCodes (Optional) List of context codes of courses whose notes you want to see.
//    If not specified, defaults to all contexts that the user belongs to.
//    The format of this field is the context type, followed by an
//    underscore, followed by the context id. For example: course_42
//    Including a code matching the user's own context code (e.g. user_1)
//    will include notes that are not associated with any particular course.
//
type ListPlannerNotes struct {
	Query struct {
		StartDate    time.Time `json:"start_date" url:"start_date,omitempty"`       //  (Optional)
		EndDate      time.Time `json:"end_date" url:"end_date,omitempty"`           //  (Optional)
		ContextCodes []string  `json:"context_codes" url:"context_codes,omitempty"` //  (Optional)
	} `json:"query"`
}

func (t *ListPlannerNotes) GetMethod() string {
	return "GET"
}

func (t *ListPlannerNotes) GetURLPath() string {
	return ""
}

func (t *ListPlannerNotes) GetQuery() (string, error) {
	v, err := query.Values(t.Query)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("?%v", v.Encode()), nil
}

func (t *ListPlannerNotes) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *ListPlannerNotes) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *ListPlannerNotes) HasErrors() error {
	return nil
}

func (t *ListPlannerNotes) Do(c *canvasapi.Canvas) ([]*models.PlannerNote, error) {
	response, err := c.SendRequest(t)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		return nil, err
	}
	ret := []*models.PlannerNote{}
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return nil, err
	}

	return ret, nil
}
