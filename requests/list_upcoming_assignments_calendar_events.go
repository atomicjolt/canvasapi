package requests

import (
	"github.com/atomicjolt/canvasapi"
)

// ListUpcomingAssignmentsCalendarEvents A paginated list of the current user's upcoming events.
// https://canvas.instructure.com/doc/api/users.html
//
type ListUpcomingAssignmentsCalendarEvents struct {
}

func (t *ListUpcomingAssignmentsCalendarEvents) GetMethod() string {
	return "GET"
}

func (t *ListUpcomingAssignmentsCalendarEvents) GetURLPath() string {
	return ""
}

func (t *ListUpcomingAssignmentsCalendarEvents) GetQuery() (string, error) {
	return "", nil
}

func (t *ListUpcomingAssignmentsCalendarEvents) GetBody() (string, error) {
	return "", nil
}

func (t *ListUpcomingAssignmentsCalendarEvents) HasErrors() error {
	return nil
}

func (t *ListUpcomingAssignmentsCalendarEvents) Do(c *canvasapi.Canvas) error {
	_, err := c.SendRequest(t)
	if err != nil {
		return err
	}

	return nil
}
