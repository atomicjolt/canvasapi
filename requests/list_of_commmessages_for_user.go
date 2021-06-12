package requests

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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
// # UserID (Required) The user id for whom you want to retrieve CommMessages
// # StartTime (Optional) The beginning of the time range you want to retrieve message from.
//    Up to a year prior to the current date is available.
// # EndTime (Optional) The end of the time range you want to retrieve messages for.
//    Up to a year prior to the current date is available.
//
type ListOfCommmessagesForUser struct {
	Query struct {
		UserID    string    `json:"user_id"`    //  (Required)
		StartTime time.Time `json:"start_time"` //  (Optional)
		EndTime   time.Time `json:"end_time"`   //  (Optional)
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
	return fmt.Sprintf("?%v", v.Encode()), nil
}

func (t *ListOfCommmessagesForUser) GetBody() (string, error) {
	return "", nil
}

func (t *ListOfCommmessagesForUser) HasErrors() error {
	errs := []string{}
	if t.Query.UserID == "" {
		errs = append(errs, "'UserID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *ListOfCommmessagesForUser) Do(c *canvasapi.Canvas) ([]*models.CommMessage, error) {
	response, err := c.SendRequest(t)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		return nil, err
	}
	ret := []*models.CommMessage{}
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return nil, err
	}

	return ret, nil
}
