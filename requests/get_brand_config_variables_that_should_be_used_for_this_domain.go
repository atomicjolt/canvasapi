package requests

import (
	"net/url"

	"github.com/atomicjolt/canvasapi"
)

// GetBrandConfigVariablesThatShouldBeUsedForThisDomain Will redirect to a static json file that has all of the brand
// variables used by this account. Even though this is a redirect,
// do not store the redirected url since if the account makes any changes
// it will redirect to a new url. Needs no authentication.
// https://canvas.instructure.com/doc/api/brand_configs.html
//
type GetBrandConfigVariablesThatShouldBeUsedForThisDomain struct {
}

func (t *GetBrandConfigVariablesThatShouldBeUsedForThisDomain) GetMethod() string {
	return "GET"
}

func (t *GetBrandConfigVariablesThatShouldBeUsedForThisDomain) GetURLPath() string {
	return ""
}

func (t *GetBrandConfigVariablesThatShouldBeUsedForThisDomain) GetQuery() (string, error) {
	return "", nil
}

func (t *GetBrandConfigVariablesThatShouldBeUsedForThisDomain) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *GetBrandConfigVariablesThatShouldBeUsedForThisDomain) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *GetBrandConfigVariablesThatShouldBeUsedForThisDomain) HasErrors() error {
	return nil
}

func (t *GetBrandConfigVariablesThatShouldBeUsedForThisDomain) Do(c *canvasapi.Canvas) error {
	_, err := c.SendRequest(t)
	if err != nil {
		return err
	}

	return nil
}
