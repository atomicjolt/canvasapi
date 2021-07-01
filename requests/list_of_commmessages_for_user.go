package requests

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/url"
	"strings"
	"time"

	"github.com/google/go-querystring/query"

	"github.com/atomicjolt/canvasapi"
	"github.com/atomicjolt/canvasapi/models"
)

// ListOfCommmessagesForUser Retrieve a paginated list of messages sent to a user.
// https://canvas.instructure.com/doc/api/comm_messages.html
//
// Query Parameters:
// # Query.UserID (Required) The user id for whom you want to retrieve CommMessages
// # Query.StartTime (Optional) The beginning of the time range you want to retrieve message from.
//    Up to a year prior to the current date is available.
// # Query.EndTime (Optional) The end of the time range you want to retrieve messages for.
//    Up to a year prior to the current date is available.
//
type ListOfCommmessagesForUser struct {
	Query struct {
		UserID    string    `json:"user_id" url:"user_id,omitempty"`       //  (Required)
		StartTime time.Time `json:"start_time" url:"start_time,omitempty"` //  (Optional)
		EndTime   time.Time `json:"end_time" url:"end_time,omitempty"`     //  (Optional)
	} `json:"query"`
}

func (t *ListOfCommmessagesForUser) GetMethod() string {
	return "GET"
}

func (t *ListOfCommmessagesForUser) GetURLPath() string {
	return ""
}

func (t *ListOfCommmessagesForUser) GetQuery() (string, error) {
	v, err := query.Values(t.Query)
	if err != nil {
		return "", err
	}
	return v.Encode(), nil
}

func (t *ListOfCommmessagesForUser) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *ListOfCommmessagesForUser) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *ListOfCommmessagesForUser) HasErrors() error {
	errs := []string{}
	if t.Query.UserID == "" {
		errs = append(errs, "'Query.UserID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *ListOfCommmessagesForUser) Do(c *canvasapi.Canvas) ([]*models.CommMessage, *canvasapi.PagedResource, error) {
	response, err := c.SendRequest(t)
	if err != nil {
		return nil, nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		return nil, nil, err
	}
	ret := []*models.CommMessage{}
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
