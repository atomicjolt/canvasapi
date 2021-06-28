package requests

import (
	"net/url"

	"github.com/atomicjolt/canvasapi"
)

// RedirectToRootOutcomeGroupForContextGlobal Convenience redirect to find the root outcome group for a particular
// context. Will redirect to the appropriate outcome group's URL.
// https://canvas.instructure.com/doc/api/outcome_groups.html
//
type RedirectToRootOutcomeGroupForContextGlobal struct {
}

func (t *RedirectToRootOutcomeGroupForContextGlobal) GetMethod() string {
	return "GET"
}

func (t *RedirectToRootOutcomeGroupForContextGlobal) GetURLPath() string {
	return ""
}

func (t *RedirectToRootOutcomeGroupForContextGlobal) GetQuery() (string, error) {
	return "", nil
}

func (t *RedirectToRootOutcomeGroupForContextGlobal) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *RedirectToRootOutcomeGroupForContextGlobal) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *RedirectToRootOutcomeGroupForContextGlobal) HasErrors() error {
	return nil
}

func (t *RedirectToRootOutcomeGroupForContextGlobal) Do(c *canvasapi.Canvas) error {
	_, err := c.SendRequest(t)
	if err != nil {
		return err
	}

	return nil
}
