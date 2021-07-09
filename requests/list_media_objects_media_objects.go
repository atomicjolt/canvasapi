package requests

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/google/go-querystring/query"

	"github.com/atomicjolt/canvasapi"
	"github.com/atomicjolt/canvasapi/models"
	"github.com/atomicjolt/string_utils"
)

// ListMediaObjectsMediaObjects Returns media objects created by the user making the request. When
// using the second version, returns media objects associated with
// the given course.
// https://canvas.instructure.com/doc/api/media_objects.html
//
// Query Parameters:
// # Query.Sort (Optional) . Must be one of title, created_atField to sort on. Default is "title"
//
//    title:: sorts on user_entered_title if available, title if not.
//
//    created_at:: sorts on the object's creation time.
// # Query.Order (Optional) . Must be one of asc, descSort direction. Default is "asc"
// # Query.Exclude (Optional) . Must be one of sources, tracksArray of data to exclude. By excluding "sources" and "tracks",
//    the api will not need to query kaltura, which greatly
//    speeds up its response.
//
//    sources:: Do not query kaltura for media_sources
//    tracks:: Do not query kaltura for media_tracks
//
type ListMediaObjectsMediaObjects struct {
	Query struct {
		Sort    string   `json:"sort" url:"sort,omitempty"`       //  (Optional) . Must be one of title, created_at
		Order   string   `json:"order" url:"order,omitempty"`     //  (Optional) . Must be one of asc, desc
		Exclude []string `json:"exclude" url:"exclude,omitempty"` //  (Optional) . Must be one of sources, tracks
	} `json:"query"`
}

func (t *ListMediaObjectsMediaObjects) GetMethod() string {
	return "GET"
}

func (t *ListMediaObjectsMediaObjects) GetURLPath() string {
	return ""
}

func (t *ListMediaObjectsMediaObjects) GetQuery() (string, error) {
	v, err := query.Values(t.Query)
	if err != nil {
		return "", err
	}
	return v.Encode(), nil
}

func (t *ListMediaObjectsMediaObjects) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *ListMediaObjectsMediaObjects) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *ListMediaObjectsMediaObjects) HasErrors() error {
	errs := []string{}
	if t.Query.Sort != "" && !string_utils.Include([]string{"title", "created_at"}, t.Query.Sort) {
		errs = append(errs, "Sort must be one of title, created_at")
	}
	if t.Query.Order != "" && !string_utils.Include([]string{"asc", "desc"}, t.Query.Order) {
		errs = append(errs, "Order must be one of asc, desc")
	}
	for _, v := range t.Query.Exclude {
		if v != "" && !string_utils.Include([]string{"sources", "tracks"}, v) {
			errs = append(errs, "Exclude must be one of sources, tracks")
		}
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *ListMediaObjectsMediaObjects) Do(c *canvasapi.Canvas, next *url.URL) ([]*models.MediaObject, *canvasapi.PagedResource, error) {
	var err error
	var response *http.Response
	if next != nil {
		response, err = c.Send(next, t.GetMethod(), nil)
	} else {
		response, err = c.SendRequest(t)
	}

	if err != nil {
		return nil, nil, err
	}
	if err != nil {
		return nil, nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		return nil, nil, err
	}
	ret := []*models.MediaObject{}
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
