package canvasapi

type SuccessResponse struct {
	Success bool `json:"success"`
}

type UnreadCount struct {
	UnreadCount int64 `json:"unread_count"`
}
