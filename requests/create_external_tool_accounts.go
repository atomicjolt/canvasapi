package requests

import (
	"fmt"
	"strings"

	"github.com/google/go-querystring/query"

	"github.com/atomicjolt/canvasapi"
	"github.com/atomicjolt/string_utils"
)

// CreateExternalToolAccounts Create an external tool in the specified course/account.
// The created tool will be returned, see the "show" endpoint for an example.
// If a client ID is supplied canvas will attempt to create a context external
// tool using the LTI 1.3 standard.
// https://canvas.instructure.com/doc/api/external_tools.html
//
// Path Parameters:
// # AccountID (Required) ID
//
// Form Parameters:
// # ClientID (Required) The client id is attached to the developer key.
//    If supplied all other parameters are unnecessary and will be ignored
// # Name (Required) The name of the tool
// # PrivacyLevel (Required) . Must be one of anonymous, name_only, publicWhat information to send to the external tool.
// # ConsumerKey (Required) The consumer key for the external tool
// # SharedSecret (Required) The shared secret with the external tool
// # Description (Optional) A description of the tool
// # Url (Optional) The url to match links against. Either "url" or "domain" should be set,
//    not both.
// # Domain (Optional) The domain to match links against. Either "url" or "domain" should be
//    set, not both.
// # IconUrl (Optional) The url of the icon to show for this tool
// # Text (Optional) The default text to show for this tool
// # CustomFields (Optional) Custom fields that will be sent to the tool consumer; can be used
//    multiple times
// # IsRceFavorite (Optional) (Deprecated in favor of {api:ExternalToolsController#add_rce_favorite Add tool to RCE Favorites} and
//    {api:ExternalToolsController#remove_rce_favorite Remove tool from RCE Favorites})
//    Whether this tool should appear in a preferred location in the RCE.
//    This only applies to tools in root account contexts that have an editor
//    button placement.
// # AccountNavigation (Optional) The url of the external tool for account navigation
// # AccountNavigation (Optional) Set this to enable this feature
// # AccountNavigation (Optional) The text that will show on the left-tab in the account navigation
// # AccountNavigation (Optional) The width of the dialog the tool is launched in
// # AccountNavigation (Optional) The height of the dialog the tool is launched in
// # AccountNavigation (Optional) The layout type to use when launching the tool. Must be
//    "full_width", "full_width_in_context", "borderless", or "default"
// # UserNavigation (Optional) The url of the external tool for user navigation
// # UserNavigation (Optional) Set this to enable this feature
// # UserNavigation (Optional) The text that will show on the left-tab in the user navigation
// # UserNavigation (Optional) . Must be one of admins, members, publicWho will see the navigation tab. "admins" for admins, "public" or
//    "members" for everyone
// # CourseHomeSubNavigation (Optional) The url of the external tool for right-side course home navigation menu
// # CourseHomeSubNavigation (Optional) Set this to enable this feature
// # CourseHomeSubNavigation (Optional) The text that will show on the right-side course home navigation menu
// # CourseHomeSubNavigation (Optional) The url of the icon to show in the right-side course home navigation menu
// # CourseNavigation (Optional) Set this to enable this feature
// # CourseNavigation (Optional) The text that will show on the left-tab in the course navigation
// # CourseNavigation (Optional) . Must be one of admins, membersWho will see the navigation tab. "admins" for course admins, "members" for
//    students, null for everyone
// # CourseNavigation (Optional) . Must be one of _blank, _selfDetermines how the navigation tab will be opened.
//    "_blank"	Launches the external tool in a new window or tab.
//    "_self"	(Default) Launches the external tool in an iframe inside of Canvas.
// # CourseNavigation (Optional) . Must be one of disabled, enabledIf set to "disabled" the tool will not appear in the course navigation
//    until a teacher explicitly enables it.
//
//    If set to "enabled" the tool will appear in the course navigation
//    without requiring a teacher to explicitly enable it.
//
//    defaults to "enabled"
// # CourseNavigation (Optional) The layout type to use when launching the tool. Must be
//    "full_width", "full_width_in_context", "borderless", or "default"
// # EditorButton (Optional) The url of the external tool
// # EditorButton (Optional) Set this to enable this feature
// # EditorButton (Optional) The url of the icon to show in the WYSIWYG editor
// # EditorButton (Optional) The width of the dialog the tool is launched in
// # EditorButton (Optional) The height of the dialog the tool is launched in
// # EditorButton (Optional) Set this to ContentItemSelectionRequest to tell the tool to use
//    content-item; otherwise, omit
// # HomeworkSubmission (Optional) The url of the external tool
// # HomeworkSubmission (Optional) Set this to enable this feature
// # HomeworkSubmission (Optional) The text that will show on the homework submission tab
// # HomeworkSubmission (Optional) Set this to ContentItemSelectionRequest to tell the tool to use
//    content-item; otherwise, omit
// # LinkSelection (Optional) The url of the external tool
// # LinkSelection (Optional) Set this to enable this feature
// # LinkSelection (Optional) The text that will show for the link selection text
// # LinkSelection (Optional) Set this to ContentItemSelectionRequest to tell the tool to use
//    content-item; otherwise, omit
// # MigrationSelection (Optional) The url of the external tool
// # MigrationSelection (Optional) Set this to enable this feature
// # MigrationSelection (Optional) Set this to ContentItemSelectionRequest to tell the tool to use
//    content-item; otherwise, omit
// # ToolConfiguration (Optional) The url of the external tool
// # ToolConfiguration (Optional) Set this to enable this feature
// # ToolConfiguration (Optional) Set this to ContentItemSelectionRequest to tell the tool to use
//    content-item; otherwise, omit
// # ToolConfiguration (Optional) Set this to default the lis_person_contact_email_primary to prefer
//    provisioned sis_email; otherwise, omit
// # ResourceSelection (Optional) The url of the external tool
// # ResourceSelection (Optional) Set this to enable this feature. If set to false,
//    not_selectable must also be set to true in order to hide this tool
//    from the selection UI in modules and assignments.
// # ResourceSelection (Optional) The url of the icon to show in the module external tool list
// # ResourceSelection (Optional) The width of the dialog the tool is launched in
// # ResourceSelection (Optional) The height of the dialog the tool is launched in
// # ConfigType (Optional) Configuration can be passed in as CC xml instead of using query
//    parameters. If this value is "by_url" or "by_xml" then an xml
//    configuration will be expected in either the "config_xml" or "config_url"
//    parameter. Note that the name parameter overrides the tool name provided
//    in the xml
// # ConfigXml (Optional) XML tool configuration, as specified in the CC xml specification. This is
//    required if "config_type" is set to "by_xml"
// # ConfigUrl (Optional) URL where the server can retrieve an XML tool configuration, as specified
//    in the CC xml specification. This is required if "config_type" is set to
//    "by_url"
// # NotSelectable (Optional) Default: false. If set to true, and if resource_selection is set to false,
//    the tool won't show up in the external tool
//    selection UI in modules and assignments
// # OauthCompliant (Optional) Default: false, if set to true LTI query params will not be copied to the
//    post body.
//
type CreateExternalToolAccounts struct {
	Path struct {
		AccountID string `json:"account_id"` //  (Required)
	} `json:"path"`

	Form struct {
		ClientID     string `json:"client_id"`     //  (Required)
		Name         string `json:"name"`          //  (Required)
		PrivacyLevel string `json:"privacy_level"` //  (Required) . Must be one of anonymous, name_only, public
		ConsumerKey  string `json:"consumer_key"`  //  (Required)
		SharedSecret string `json:"shared_secret"` //  (Required)
		Description  string `json:"description"`   //  (Optional)
		Url          string `json:"url"`           //  (Optional)
		Domain       string `json:"domain"`        //  (Optional)
		IconUrl      string `json:"icon_url"`      //  (Optional)
		Text         string `json:"text"`          //  (Optional)
		CustomFields struct {
			FieldName string `json:"field_name"` //  (Optional)
		} `json:"custom_fields"`

		IsRceFavorite     bool `json:"is_rce_favorite"` //  (Optional)
		AccountNavigation struct {
			Url             string `json:"url"`              //  (Optional)
			Enabled         bool   `json:"enabled"`          //  (Optional)
			Text            string `json:"text"`             //  (Optional)
			SelectionWidth  string `json:"selection_width"`  //  (Optional)
			SelectionHeight string `json:"selection_height"` //  (Optional)
			DisplayType     string `json:"display_type"`     //  (Optional)
		} `json:"account_navigation"`

		UserNavigation struct {
			Url        string `json:"url"`        //  (Optional)
			Enabled    bool   `json:"enabled"`    //  (Optional)
			Text       string `json:"text"`       //  (Optional)
			Visibility string `json:"visibility"` //  (Optional) . Must be one of admins, members, public
		} `json:"user_navigation"`

		CourseHomeSubNavigation struct {
			Url     string `json:"url"`      //  (Optional)
			Enabled bool   `json:"enabled"`  //  (Optional)
			Text    string `json:"text"`     //  (Optional)
			IconUrl string `json:"icon_url"` //  (Optional)
		} `json:"course_home_sub_navigation"`

		CourseNavigation struct {
			Enabled      bool   `json:"enabled"`       //  (Optional)
			Text         string `json:"text"`          //  (Optional)
			Visibility   string `json:"visibility"`    //  (Optional) . Must be one of admins, members
			WindowTarget string `json:"window_target"` //  (Optional) . Must be one of _blank, _self
			Default      string `json:"default"`       //  (Optional) . Must be one of disabled, enabled
			DisplayType  string `json:"display_type"`  //  (Optional)
		} `json:"course_navigation"`

		EditorButton struct {
			Url             string `json:"url"`              //  (Optional)
			Enabled         bool   `json:"enabled"`          //  (Optional)
			IconUrl         string `json:"icon_url"`         //  (Optional)
			SelectionWidth  string `json:"selection_width"`  //  (Optional)
			SelectionHeight string `json:"selection_height"` //  (Optional)
			MessageType     string `json:"message_type"`     //  (Optional)
		} `json:"editor_button"`

		HomeworkSubmission struct {
			Url         string `json:"url"`          //  (Optional)
			Enabled     bool   `json:"enabled"`      //  (Optional)
			Text        string `json:"text"`         //  (Optional)
			MessageType string `json:"message_type"` //  (Optional)
		} `json:"homework_submission"`

		LinkSelection struct {
			Url         string `json:"url"`          //  (Optional)
			Enabled     bool   `json:"enabled"`      //  (Optional)
			Text        string `json:"text"`         //  (Optional)
			MessageType string `json:"message_type"` //  (Optional)
		} `json:"link_selection"`

		MigrationSelection struct {
			Url         string `json:"url"`          //  (Optional)
			Enabled     bool   `json:"enabled"`      //  (Optional)
			MessageType string `json:"message_type"` //  (Optional)
		} `json:"migration_selection"`

		ToolConfiguration struct {
			Url            string `json:"url"`              //  (Optional)
			Enabled        bool   `json:"enabled"`          //  (Optional)
			MessageType    string `json:"message_type"`     //  (Optional)
			PreferSISEmail bool   `json:"prefer_sis_email"` //  (Optional)
		} `json:"tool_configuration"`

		ResourceSelection struct {
			Url             string `json:"url"`              //  (Optional)
			Enabled         bool   `json:"enabled"`          //  (Optional)
			IconUrl         string `json:"icon_url"`         //  (Optional)
			SelectionWidth  string `json:"selection_width"`  //  (Optional)
			SelectionHeight string `json:"selection_height"` //  (Optional)
		} `json:"resource_selection"`

		ConfigType     string `json:"config_type"`     //  (Optional)
		ConfigXml      string `json:"config_xml"`      //  (Optional)
		ConfigUrl      string `json:"config_url"`      //  (Optional)
		NotSelectable  bool   `json:"not_selectable"`  //  (Optional)
		OauthCompliant bool   `json:"oauth_compliant"` //  (Optional)
	} `json:"form"`
}

func (t *CreateExternalToolAccounts) GetMethod() string {
	return "POST"
}

func (t *CreateExternalToolAccounts) GetURLPath() string {
	path := "accounts/{account_id}/external_tools"
	path = strings.ReplaceAll(path, "{account_id}", fmt.Sprintf("%v", t.Path.AccountID))
	return path
}

func (t *CreateExternalToolAccounts) GetQuery() (string, error) {
	return "", nil
}

func (t *CreateExternalToolAccounts) GetBody() (string, error) {
	v, err := query.Values(t.Form)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%v", v.Encode()), nil
}

func (t *CreateExternalToolAccounts) HasErrors() error {
	errs := []string{}
	if t.Path.AccountID == "" {
		errs = append(errs, "'AccountID' is required")
	}
	if t.Form.ClientID == "" {
		errs = append(errs, "'ClientID' is required")
	}
	if t.Form.Name == "" {
		errs = append(errs, "'Name' is required")
	}
	if t.Form.PrivacyLevel == "" {
		errs = append(errs, "'PrivacyLevel' is required")
	}
	if !string_utils.Include([]string{"anonymous", "name_only", "public"}, t.Form.PrivacyLevel) {
		errs = append(errs, "PrivacyLevel must be one of anonymous, name_only, public")
	}
	if t.Form.ConsumerKey == "" {
		errs = append(errs, "'ConsumerKey' is required")
	}
	if t.Form.SharedSecret == "" {
		errs = append(errs, "'SharedSecret' is required")
	}
	if !string_utils.Include([]string{"admins", "members", "public"}, t.Form.UserNavigation.Visibility) {
		errs = append(errs, "UserNavigation must be one of admins, members, public")
	}
	if !string_utils.Include([]string{"admins", "members"}, t.Form.CourseNavigation.Visibility) {
		errs = append(errs, "CourseNavigation must be one of admins, members")
	}
	if !string_utils.Include([]string{"_blank", "_self"}, t.Form.CourseNavigation.WindowTarget) {
		errs = append(errs, "CourseNavigation must be one of _blank, _self")
	}
	if !string_utils.Include([]string{"disabled", "enabled"}, t.Form.CourseNavigation.Default) {
		errs = append(errs, "CourseNavigation must be one of disabled, enabled")
	}
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

func (t *CreateExternalToolAccounts) Do(c *canvasapi.Canvas) error {
	_, err := c.SendRequest(t)
	if err != nil {
		return err
	}

	return nil
}
