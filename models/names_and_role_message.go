package models

type NamesAndRoleMessage struct {
	MessageType       string `json:"message_type"`         // The type of LTI message being described. Always set to 'LtiResourceLinkRequest'.Example: LtiResourceLinkRequest
	Locale            string `json:"locale"`               // The member's preferred locale.Example: en
	CanvasUserID      int64  `json:"canvas_user_id"`       // The member's API ID.Example: 1
	CanvasUserLoginID string `json:"canvas_user_login_id"` // The member's primary login username.Example: showell@school.edu
	Custom            string `json:"custom"`               // Expanded LTI custom parameters that pertain to the member (as opposed to the Context).Example: en, America/Denver
}

func (t *NamesAndRoleMessage) HasError() error {
	return nil
}
