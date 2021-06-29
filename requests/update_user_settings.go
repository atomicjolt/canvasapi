package requests

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/google/go-querystring/query"

	"github.com/atomicjolt/canvasapi"
)

// UpdateUserSettings Update an existing user's settings.
// https://canvas.instructure.com/doc/api/users.html
//
// Path Parameters:
// # ID (Required) ID
//
// Query Parameters:
// # ManualMarkAsRead (Optional) If true, require user to manually mark discussion posts as read (don't
//    auto-mark as read).
// # ReleaseNotesBadgeDisabled (Optional) If true, hide the badge for new release notes.
// # CollapseGlobalNav (Optional) If true, the user's page loads with the global navigation collapsed
// # HideDashcardColorOverlays (Optional) If true, images on course cards will be presented without being tinted
//    to match the course color.
// # CommentLibrarySuggestionsEnabled (Optional) If true, suggestions within the comment library will be shown.
// # ElementaryDashboardDisabled (Optional) If true, will display the user's preferred class Canvas dashboard
//    view instead of the canvas for elementary view.
//
type UpdateUserSettings struct {
	Path struct {
		ID string `json:"id" url:"id,omitempty"` //  (Required)
	} `json:"path"`

	Query struct {
		ManualMarkAsRead                 bool `json:"manual_mark_as_read" url:"manual_mark_as_read,omitempty"`                                 //  (Optional)
		ReleaseNotesBadgeDisabled        bool `json:"release_notes_badge_disabled" url:"release_notes_badge_disabled,omitempty"`               //  (Optional)
		CollapseGlobalNav                bool `json:"collapse_global_nav" url:"collapse_global_nav,omitempty"`                                 //  (Optional)
		HideDashcardColorOverlays        bool `json:"hide_dashcard_color_overlays" url:"hide_dashcard_color_overlays,omitempty"`               //  (Optional)
		CommentLibrarySuggestionsEnabled bool `json:"comment_library_suggestions_enabled" url:"comment_library_suggestions_enabled,omitempty"` //  (Optional)
		ElementaryDashboardDisabled      bool `json:"elementary_dashboard_disabled" url:"elementary_dashboard_disabled,omitempty"`             //  (Optional)
	} `json:"query"`
}

func (t *UpdateUserSettings) GetMethod() string {
	return "GET"
}

func (t *UpdateUserSettings) GetURLPath() string {
	path := "users/{id}/settings"
	path = strings.ReplaceAll(path, "{id}", fmt.Sprintf("%v", t.Path.ID))
	return path
}

func (t *UpdateUserSettings) GetQuery() (string, error) {
	v, err := query.Values(t.Query)
	if err != nil {
		return "", err
	}
	return v.Encode(), nil
}

func (t *UpdateUserSettings) GetBody() (url.Values, error) {
	return nil, nil
}

func (t *UpdateUserSettings) GetJSON() ([]byte, error) {
	return nil, nil
}

func (t *UpdateUserSettings) HasErrors() error {
	errs := []string{}
	if t.Path.ID == "" {
		errs = append(errs, "'ID' is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *UpdateUserSettings) Do(c *canvasapi.Canvas) error {
	_, err := c.SendRequest(t)
	if err != nil {
		return err
	}

	return nil
}
