package requests

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/google/go-querystring/query"

	"github.com/atomicjolt/canvasapi"
	"github.com/atomicjolt/canvasapi/models"
	"github.com/atomicjolt/string_utils"
)

// ListMediaObjectsCourses Returns media objects created by the user making the request. When
// using the second version, returns media objects associated with
// the given course.
// https://canvas.instructure.com/doc/api/media_objects.html
//
// Path Parameters:
// # CourseID (Required) ID
//
// Query Parameters:
// # Sort (Optional) . Must be one of title, created_atField to sort on. Default is "title"
//
//    title:: sorts on user_entered_title if available, title if not.
//
//    created_at:: sorts on the object's creation time.
// # Order (Optional) . Must be one of asc, descSort direction. Default is "asc"
// # Exclude (Optional) . Must be one of sources, tracksArray of data to exclude. By excluding "sources" and "tracks",
//    the api will not need to query kaltura, which greatly
//    speeds up its response.
//
//    sources:: Do not query kaltura for media_sources
//    tracks:: Do not query kaltura for media_tracks
//
type ListMediaObjectsCourses struct {
	Path struct {
		CourseID string `json:"course_id"` //  (Required)
	} `json:"path"`

	Query struct {
		Sort    string   `json:"sort"`    //  (Optional) . Must be one of title, created_at
		Order   string   `json:"order"`   //  (Optional) . Must be one of asc, desc
		Exclude []string `json:"exclude"` //  (Optional) . Must be one of sources, tracks
	} `json:"query"`
}

func (t *ListMediaObjectsCourses) GetMethod() string {
	return "GET"
}

func (t *ListMediaObjectsCourses) GetURLPath() string {
	path := "courses/{course_id}/media_objects"
	path = strings.ReplaceAll(path, "{course_id}", fmt.Sprintf("%v", t.Path.CourseID))
	return path
}

func (t *ListMediaObjectsCourses) GetQuery() (string, error) {
	v, err := query.Values(t.Query)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("?%v", v.Encode()), nil
}

func (t *ListMediaObjectsCourses) GetBody() (string, error) {
	return "", nil
}

func (t *ListMediaObjectsCourses) HasErrors() error {
	errs := []string{}
	if t.Path.CourseID == "" {
		errs = append(errs, "'CourseID' is required")
	}
	if !string_utils.Include([]string{"title", "created_at"}, t.Query.Sort) {
		errs = append(errs, "Sort must be one of title, created_at")
	}
	if !string_utils.Include([]string{"asc", "desc"}, t.Query.Order) {
		errs = append(errs, "Order must be one of asc, desc")
	}
	for _, v := range t.Query.Exclude {
		if !string_utils.Include([]string{"sources", "tracks"}, v) {
			errs = append(errs, "Exclude must be one of sources, tracks")
		}
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *ListMediaObjectsCourses) Do(c *canvasapi.Canvas) ([]*models.MediaObject, error) {
	response, err := c.SendRequest(t)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		return nil, err
	}
	ret := []*models.MediaObject{}
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return nil, err
	}

	return ret, nil
}
