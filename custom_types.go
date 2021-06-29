package canvasapi

type QuizUserConversation struct {
	Body       string `json:"body" url:"body,omitempty"`
	Subject    string `json:"subject" url:"subject,omitempty"`
	Recipients string `json:"recipients" url:"recipients,omitempty"`
}
