package models

type NamesAndRoleMessage struct {
	MessageType       string                   `json:"message_type" url:"message_type,omitempty"`                 // The type of LTI message being described. Always set to 'LtiResourceLinkRequest'.Example: LtiResourceLinkRequest
	Locale            string                   `json:"locale" url:"locale,omitempty"`                             // The member's preferred locale.Example: en
	CanvasUserID      int64                    `json:"canvas_user_id" url:"canvas_user_id,omitempty"`             // The member's API ID.Example: 1
	CanvasUserLoginID string                   `json:"canvas_user_login_id" url:"canvas_user_login_id,omitempty"` // The member's primary login username.Example: showell@school.edu
	Custom            map[string](interface{}) `json:"custom" url:"custom,omitempty"`                             // Expanded LTI custom parameters that pertain to the member (as opposed to the Context).Example: en, America/Denver
}

func (t *NamesAndRoleMessage) HasErrors() error {
	return nil
}
