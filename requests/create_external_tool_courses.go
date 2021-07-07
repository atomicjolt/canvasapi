package requests

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strings"

	"github.com/google/go-querystring/query"

	"github.com/atomicjolt/canvasapi"
	"github.com/atomicjolt/string_utils"
)

// CreateExternalToolCourses Create an external tool in the specified course/account.
// The created tool will be returned, see the "show" endpoint for an example.
// If a client ID is supplied canvas will attempt to create a context external
// tool using the LTI 1.3 standard.
// https://canvas.instructure.com/doc/api/external_tools.html
//
// Path Parameters:
// # Path.CourseID (Required) ID
//
// Form Parameters:
// # Form.ClientID (Required) The client id is attached to the developer key.
//    If supplied all other parameters are unnecessary and will be ignored
// # Form.Name (Required) The name of the tool
// # Form.PrivacyLevel (Required) . Must be one of anonymous, name_only, publicWhat information to send to the external tool.
// # Form.ConsumerKey (Required) The consumer key for the external tool
// # Form.SharedSecret (Required) The shared secret with the external tool
// # Form.Description (Optional) A description of the tool
// # Form.Url (Optional) The url to match links against. Either "url" or "domain" should be set,
//    not both.
// # Form.Domain (Optional) The domain to match links against. Either "url" or "domain" should be
//    set, not both.
// # Form.IconUrl (Optional) The url of the icon to show for this tool
// # Form.Text (Optional) The default text to show for this tool
// # Form.CustomFields.FieldName (Optional) Custom fields that will be sent to the tool consumer; can be used
//    multiple times
// # Form.IsRceFavorite (Optional) (Deprecated in favor of {api:ExternalToolsController#add_rce_favorite Add tool to RCE Favorites} and
//    {api:ExternalToolsController#remove_rce_favorite Remove tool from RCE Favorites})
//    Whether this tool should appear in a preferred location in the RCE.
//    This only applies to tools in root account contexts that have an editor
//    button placement.
// # Form.AccountNavigation.Url (Optional) The url of the external tool for account navigation
// # Form.AccountNavigation.Enabled (Optional) Set this to enable this feature
// # Form.AccountNavigation.Text (Optional) The text that will show on the left-tab in the account navigation
// # Form.AccountNavigation.SelectionWidth (Optional) The width of the dialog the tool is launched in
// # Form.AccountNavigation.SelectionHeight (Optional) The height of the dialog the tool is launched in
// # Form.AccountNavigation.DisplayType (Optional) The layout type to use when launching the tool. Must be
//    "full_width", "full_width_in_context", "in_nav_context", "borderless", or "default"
// # Form.UserNavigation.Url (Optional) The url of the external tool for user navigation
// # Form.UserNavigation.Enabled (Optional) Set this to enable this feature
// # Form.UserNavigation.Text (Optional) The text that will show on the left-tab in the user navigation
// # Form.UserNavigation.Visibility (Optional) . Must be one of admins, members, publicWho will see the navigation tab. "admins" for admins, "public" or
//    "members" for everyone
// # Form.CourseHomeSubNavigation.Url (Optional) The url of the external tool for right-side course home navigation menu
// # Form.CourseHomeSubNavigation.Enabled (Optional) Set this to enable this feature
// # Form.CourseHomeSubNavigation.Text (Optional) The text that will show on the right-side course home navigation menu
// # Form.CourseHomeSubNavigation.IconUrl (Optional) The url of the icon to show in the right-side course home navigation menu
// # Form.CourseNavigation.Enabled (Optional) Set this to enable this feature
// # Form.CourseNavigation.Text (Optional) The text that will show on the left-tab in the course navigation
// # Form.CourseNavigation.Visibility (Optional) . Must be one of admins, membersWho will see the navigation tab. "admins" for course admins, "members" for
//    students, null for everyone
// # Form.CourseNavigation.WindowTarget (Optional) . Must be one of _blank, _selfDetermines how the navigation tab will be opened.
//    "_blank"	Launches the external tool in a new window or tab.
//    "_self"	(Default) Launches the external tool in an iframe inside of Canvas.
// # Form.CourseNavigation.Default (Optional) . Must be one of disabled, enabledIf set to "disabled" the tool will not appear in the course navigation
//    until a teacher explicitly enables it.
//
//    If set to "enabled" the tool will appear in the course navigation
//    without requiring a teacher to explicitly enable it.
//
//    defaults to "enabled"
// # Form.CourseNavigation.DisplayType (Optional) The layout type to use when launching the tool. Must be
//    "full_width", "full_width_in_context", "in_nav_context", "borderless", or "default"
// # Form.EditorButton.Url (Optional) The url of the external tool
// # Form.EditorButton.Enabled (Optional) Set this to enable this feature
// # Form.EditorButton.IconUrl (Optional) The url of the icon to show in the WYSIWYG editor
// # Form.EditorButton.SelectionWidth (Optional) The width of the dialog the tool is launched in
// # Form.EditorButton.SelectionHeight (Optional) The height of the dialog the tool is launched in
// # Form.EditorButton.MessageType (Optional) Set this to ContentItemSelectionRequest to tell the tool to use
//    content-item; otherwise, omit
// # Form.HomeworkSubmission.Url (Optional) The url of the external tool
// # Form.HomeworkSubmission.Enabled (Optional) Set this to enable this feature
// # Form.HomeworkSubmission.Text (Optional) The text that will show on the homework submission tab
// # Form.HomeworkSubmission.MessageType (Optional) Set this to ContentItemSelectionRequest to tell the tool to use
//    content-item; otherwise, omit
// # Form.LinkSelection.Url (Optional) The url of the external tool
// # Form.LinkSelection.Enabled (Optional) Set this to enable this feature
// # Form.LinkSelection.Text (Optional) The text that will show for the link selection text
// # Form.LinkSelection.MessageType (Optional) Set this to ContentItemSelectionRequest to tell the tool to use
//    content-item; otherwise, omit
// # Form.MigrationSelection.Url (Optional) The url of the external tool
// # Form.MigrationSelection.Enabled (Optional) Set this to enable this feature
// # Form.MigrationSelection.MessageType (Optional) Set this to ContentItemSelectionRequest to tell the tool to use
//    content-item; otherwise, omit
// # Form.ToolConfiguration.Url (Optional) The url of the external tool
// # Form.ToolConfiguration.Enabled (Optional) Set this to enable this feature
// # Form.ToolConfiguration.MessageType (Optional) Set this to ContentItemSelectionRequest to tell the tool to use
//    content-item; otherwise, omit
// # Form.ToolConfiguration.PreferSISEmail (Optional) Set this to default the lis_person_contact_email_primary to prefer
//    provisioned sis_email; otherwise, omit
// # Form.ResourceSelection.Url (Optional) The url of the external tool
// # Form.ResourceSelection.Enabled (Optional) Set this to enable this feature. If set to false,
//    not_selectable must also be set to true in order to hide this tool
//    from the selection UI in modules and assignments.
// # Form.ResourceSelection.IconUrl (Optional) The url of the icon to show in the module external tool list
// # Form.ResourceSelection.SelectionWidth (Optional) The width of the dialog the tool is launched in
// # Form.ResourceSelection.SelectionHeight (Optional) The height of the dialog the tool is launched in
// # Form.ConfigType (Optional) Configuration can be passed in as CC xml instead of using query
//    parameters. If this value is "by_url" or "by_xml" then an xml
//    configuration will be expected in either the "config_xml" or "config_url"
//    parameter. Note that the name parameter overrides the tool name provided
//    in the xml
// # Form.ConfigXml (Optional) XML tool configuration, as specified in the CC xml specification. This is
//    required if "config_type" is set to "by_xml"
// # Form.ConfigUrl (Optional) URL where the server can retrieve an XML tool configuration, as specified
//    in the CC xml specification. This is required if "config_type" is set to
//    "by_url"
// # Form.NotSelectable (Optional) Default: false. If set to true, and if resource_selection is set to false,
//    the tool won't show up in the external tool
//    selection UI in modules and assignments
// # Form.OauthCompliant (Optional) Default: false, if set to true LTI query params will not be copied to the
//    post body.
//
type CreateExternalToolCourses struct {
	Path struct {
		CourseID string `json:"course_id" url:"course_id,omitempty"` //  (Required)
	} `json:"path"`

	Form struct {
		ClientID     string `json:"client_id" url:"client_id,omitempty"`         //  (Required)
		Name         string `json:"name" url:"name,omitempty"`                   //  (Required)
		PrivacyLevel string `json:"privacy_level" url:"privacy_level,omitempty"` //  (Required) . Must be one of anonymous, name_only, public
		ConsumerKey  string `json:"consumer_key" url:"consumer_key,omitempty"`   //  (Required)
		SharedSecret string `json:"shared_secret" url:"shared_secret,omitempty"` //  (Required)
		Description  string `json:"description" url:"description,omitempty"`     //  (Optional)
		Url          string `json:"url" url:"url,omitempty"`                     //  (Optional)
		Domain       string `json:"domain" url:"domain,omitempty"`               //  (Optional)
		IconUrl      string `json:"icon_url" url:"icon_url,omitempty"`           //  (Optional)
		Text         string `json:"text" url:"text,omitempty"`                   //  (Optional)
		CustomFields struct {
			FieldName string `json:"field_name" url:"field_name,omitempty"` //  (Optional)
		} `json:"custom_fields" url:"custom_fields,omitempty"`

		IsRceFavorite     bool `json:"is_rce_favorite" url:"is_rce_favorite,omitempty"` //  (Optional)
		AccountNavigation struct {
			Url             string `json:"url" url:"url,omitempty"`                           //  (Optional)
			Enabled         bool   `json:"enabled" url:"enabled,omitempty"`                   //  (Optional)
			Text            string `json:"text" url:"text,omitempty"`                         //  (Optional)
			SelectionWidth  string `json:"selection_width" url:"selection_width,omitempty"`   //  (Optional)
			SelectionHeight string `json:"selection_height" url:"selection_height,omitempty"` //  (Optional)
			DisplayType     string `json:"display_type" url:"display_type,omitempty"`         //  (Optional)
		} `json:"account_navigation" url:"account_navigation,omitempty"`

		UserNavigation struct {
			Url        string `json:"url" url:"url,omitempty"`               //  (Optional)
			Enabled    bool   `json:"enabled" url:"enabled,omitempty"`       //  (Optional)
			Text       string `json:"text" url:"text,omitempty"`             //  (Optional)
			Visibility string `json:"visibility" url:"visibility,omitempty"` //  (Optional) . Must be one of admins, members, public
		} `json:"user_navigation" url:"user_navigation,omitempty"`

		CourseHomeSubNavigation struct {
			Url     string `json:"url" url:"url,omitempty"`           //  (Optional)
			Enabled bool   `json:"enabled" url:"enabled,omitempty"`   //  (Optional)
			Text    string `json:"text" url:"text,omitempty"`         //  (Optional)
			IconUrl string `json:"icon_url" url:"icon_url,omitempty"` //  (Optional)
		} `json:"course_home_sub_navigation" url:"course_home_sub_navigation,omitempty"`

		CourseNavigation struct {
			Enabled      bool   `json:"enabled" url:"enabled,omitempty"`             //  (Optional)
			Text         string `json:"text" url:"text,omitempty"`                   //  (Optional)
			Visibility   string `json:"visibility" url:"visibility,omitempty"`       //  (Optional) . Must be one of admins, members
			WindowTarget string `json:"window_target" url:"window_target,omitempty"` //  (Optional) . Must be one of _blank, _self
			Default      string `json:"default" url:"default,omitempty"`             //  (Optional) . Must be one of disabled, enabled
			DisplayType  string `json:"display_type" url:"display_type,omitempty"`   //  (Optional)
		} `json:"course_navigation" url:"course_navigation,omitempty"`

		EditorButton struct {
			Url             string `json:"url" url:"url,omitempty"`                           //  (Optional)
			Enabled         bool   `json:"enabled" url:"enabled,omitempty"`                   //  (Optional)
			IconUrl         string `json:"icon_url" url:"icon_url,omitempty"`                 //  (Optional)
			SelectionWidth  string `json:"selection_width" url:"selection_width,omitempty"`   //  (Optional)
			SelectionHeight string `json:"selection_height" url:"selection_height,omitempty"` //  (Optional)
			MessageType     string `json:"message_type" url:"message_type,omitempty"`         //  (Optional)
		} `json:"editor_button" url:"editor_button,omitempty"`

		HomeworkSubmission struct {
			Url         string `json:"url" url:"url,omitempty"`                   //  (Optional)
			Enabled     bool   `json:"enabled" url:"enabled,omitempty"`           //  (Optional)
			Text        string `json:"text" url:"text,omitempty"`                 //  (Optional)
			MessageType string `json:"message_type" url:"message_type,omitempty"` //  (Optional)
		} `json:"homework_submission" url:"homework_submission,omitempty"`

		LinkSelection struct {
			Url         string `json:"url" url:"url,omitempty"`                   //  (Optional)
			Enabled     bool   `json:"enabled" url:"enabled,omitempty"`           //  (Optional)
			Text        string `json:"text" url:"text,omitempty"`                 //  (Optional)
			MessageType string `json:"message_type" url:"message_type,omitempty"` //  (Optional)
		} `json:"link_selection" url:"link_selection,omitempty"`

		MigrationSelection struct {
			Url         string `json:"url" url:"url,omitempty"`                   //  (Optional)
			Enabled     bool   `json:"enabled" url:"enabled,omitempty"`           //  (Optional)
			MessageType string `json:"message_type" url:"message_type,omitempty"` //  (Optional)
		} `json:"migration_selection" url:"migration_selection,omitempty"`

		ToolConfiguration struct {
			Url            string `json:"url" url:"url,omitempty"`                           //  (Optional)
			Enabled        bool   `json:"enabled" url:"enabled,omitempty"`                   //  (Optional)
			MessageType    string `json:"message_type" url:"message_type,omitempty"`         //  (Optional)
			PreferSISEmail bool   `json:"prefer_sis_email" url:"prefer_sis_email,omitempty"` //  (Optional)
		} `json:"tool_configuration" url:"tool_configuration,omitempty"`

		ResourceSelection struct {
			Url             string `json:"url" url:"url,omitempty"`                           //  (Optional)
			Enabled         bool   `json:"enabled" url:"enabled,omitempty"`                   //  (Optional)
			IconUrl         string `json:"icon_url" url:"icon_url,omitempty"`                 //  (Optional)
			SelectionWidth  string `json:"selection_width" url:"selection_width,omitempty"`   //  (Optional)
			SelectionHeight string `json:"selection_height" url:"selection_height,omitempty"` //  (Optional)
		} `json:"resource_selection" url:"resource_selection,omitempty"`

		ConfigType     string `json:"config_type" url:"config_type,omitempty"`         //  (Optional)
		ConfigXml      string `json:"config_xml" url:"config_xml,omitempty"`           //  (Optional)
		ConfigUrl      string `json:"config_url" url:"config_url,omitempty"`           //  (Optional)
		NotSelectable  bool   `json:"not_selectable" url:"not_selectable,omitempty"`   //  (Optional)
		OauthCompliant bool   `json:"oauth_compliant" url:"oauth_compliant,omitempty"` //  (Optional)
	} `json:"form"`
}

func (t *CreateExternalToolCourses) GetMethod() string {
	return "POST"
}

func (t *CreateExternalToolCourses) GetURLPath() string {
	path := "courses/{course_id}/external_tools"
	path = strings.ReplaceAll(path, "{course_id}", fmt.Sprintf("%v", t.Path.CourseID))
	return path
}

func (t *CreateExternalToolCourses) GetQuery() (string, error) {
	return "", nil
}

func (t *CreateExternalToolCourses) GetBody() (url.Values, error) {
	return query.Values(t.Form)
}

func (t *CreateExternalToolCourses) GetJSON() ([]byte, error) {
	j, err := json.Marshal(t.Form)
	if err != nil {
		return nil, nil
	}
	return j, nil
}

func (t *CreateExternalToolCourses) HasErrors() error {
	errs := []string{}
	if t.Path.CourseID == "" {
		errs = append(errs, "'Path.CourseID' is required")
	}
	if t.Form.ClientID == "" {
		errs = append(errs, "'Form.ClientID' is required")
	}
	if t.Form.Name == "" {
		errs = append(errs, "'Form.Name' is required")
	}
	if t.Form.PrivacyLevel == "" {
		errs = append(errs, "'Form.PrivacyLevel' is required")
	}
	if t.Form.PrivacyLevel != "" && !string_utils.Include([]string{"anonymous", "name_only", "public"}, t.Form.PrivacyLevel) {
		errs = append(errs, "PrivacyLevel must be one of anonymous, name_only, public")
	}
	if t.Form.ConsumerKey == "" {
		errs = append(errs, "'Form.ConsumerKey' is required")
	}
	if t.Form.SharedSecret == "" {
		errs = append(errs, "'Form.SharedSecret' is required")
	}
	if t.Form.UserNavigation.Visibility != "" && !string_utils.Include([]string{"admins", "members", "public"}, t.Form.UserNavigation.Visibility) {
		errs = append(errs, "UserNavigation must be one of admins, members, public")
	}
	if t.Form.CourseNavigation.Visibility != "" && !string_utils.Include([]string{"admins", "members"}, t.Form.CourseNavigation.Visibility) {
		errs = append(errs, "CourseNavigation must be one of admins, members")
	}
	if t.Form.CourseNavigation.WindowTarget != "" && !string_utils.Include([]string{"_blank", "_self"}, t.Form.CourseNavigation.WindowTarget) {
		errs = append(errs, "CourseNavigation must be one of _blank, _self")
	}
	if t.Form.CourseNavigation.Default != "" && !string_utils.Include([]string{"disabled", "enabled"}, t.Form.CourseNavigation.Default) {
		errs = append(errs, "CourseNavigation must be one of disabled, enabled")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *CreateExternalToolCourses) Do(c *canvasapi.Canvas) error {
	_, err := c.SendRequest(t)
	if err != nil {
		return err
	}

	return nil
}
