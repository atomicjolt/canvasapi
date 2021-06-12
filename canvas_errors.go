package canvasapi

import (
	"errors"
)

// ErrRateLimitExceeded is returned when the api rate limit has been reached.
var ErrRateLimitExceeded = errors.New("403 Forbidden (Rate Limit Exceeded)")

// IsRateLimit returns true if the error given is a rate limit error.
func IsRateLimit(e error) bool {
	return e == ErrRateLimitExceeded
}
