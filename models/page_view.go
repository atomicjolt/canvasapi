package models

import (
	"time"
)

type PageView struct {
	ID                 string         `json:"id" url:"id,omitempty"`                                   // A UUID representing the page view.  This is also the unique request id.Example: 3e246700-e305-0130-51de-02e33aa501ef
	AppName            string         `json:"app_name" url:"app_name,omitempty"`                       // If the request is from an API request, the app that generated the access token.Example: Canvas for iOS
	Url                string         `json:"url" url:"url,omitempty"`                                 // The URL requested.Example: https://canvas.instructure.com/conversations
	ContextType        string         `json:"context_type" url:"context_type,omitempty"`               // The type of context for the request.Example: Course
	AssetType          string         `json:"asset_type" url:"asset_type,omitempty"`                   // The type of asset in the context for the request, if any.Example: Discussion
	Controller         string         `json:"controller" url:"controller,omitempty"`                   // The rails controller that handled the request.Example: discussions
	Action             string         `json:"action" url:"action,omitempty"`                           // The rails action that handled the request.Example: index
	Contributed        bool           `json:"contributed" url:"contributed,omitempty"`                 // This field is deprecated, and will always be false.Example: false
	InteractionSeconds float64        `json:"interaction_seconds" url:"interaction_seconds,omitempty"` // An approximation of how long the user spent on the page, in seconds.Example: 7.21
	CreatedAt          time.Time      `json:"created_at" url:"created_at,omitempty"`                   // When the request was made.Example: 2013-10-01T19:49:47Z
	UserRequest        bool           `json:"user_request" url:"user_request,omitempty"`               // A flag indicating whether the request was user-initiated, or automatic (such as an AJAX call).Example: true
	RenderTime         float64        `json:"render_time" url:"render_time,omitempty"`                 // How long the response took to render, in seconds.Example: 0.369
	UserAgent          string         `json:"user_agent" url:"user_agent,omitempty"`                   // The user-agent of the browser or program that made the request.Example: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_8_5) AppleWebKit/536.30.1 (KHTML, like Gecko) Version/6.0.5 Safari/536.30.1
	Participated       bool           `json:"participated" url:"participated,omitempty"`               // True if the request counted as participating, such as submitting homework.Example: false
	HttpMethod         string         `json:"http_method" url:"http_method,omitempty"`                 // The HTTP method such as GET or POST.Example: GET
	RemoteIp           string         `json:"remote_ip" url:"remote_ip,omitempty"`                     // The origin IP address of the request.Example: 173.194.46.71
	Links              *PageViewLinks `json:"links" url:"links,omitempty"`                             // The page view links to define the relationships.Example: 1234, 1234
}

func (t *PageView) HasError() error {
	return nil
}
