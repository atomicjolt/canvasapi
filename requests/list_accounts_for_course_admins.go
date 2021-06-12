package requests

import (
	"encoding/json"
	"io/ioutil"

	"github.com/atomicjolt/canvasapi"
	"github.com/atomicjolt/canvasapi/models"
)

// ListAccountsForCourseAdmins A paginated list of accounts that the current user can view through their
// admin course enrollments. (Teacher, TA, or designer enrollments).
// Only returns "id", "name", "workflow_state", "root_account_id" and "parent_account_id"
// https://canvas.instructure.com/doc/api/accounts.html
//
type ListAccountsForCourseAdmins struct {
}

func (t *ListAccountsForCourseAdmins) GetMethod() string {
	return "GET"
}

func (t *ListAccountsForCourseAdmins) GetURLPath() string {
	return ""
}

func (t *ListAccountsForCourseAdmins) GetQuery() (string, error) {
	return "", nil
}

func (t *ListAccountsForCourseAdmins) GetBody() (string, error) {
	return "", nil
}

func (t *ListAccountsForCourseAdmins) HasErrors() error {
	return nil
}

func (t *ListAccountsForCourseAdmins) Do(c *canvasapi.Canvas) ([]*models.Account, error) {
	response, err := c.SendRequest(t)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		return nil, err
	}
	ret := []*models.Account{}
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return nil, err
	}

	return ret, nil
}