package utopiago

import "strings"

// CheckErrorConnBroken - check the text of the request error, determining whether the client must be restarted
func CheckErrorConnBroken(err error) bool {
	if err == nil {
		return false
	}

	return strings.Contains(err.Error(), "read: connection reset by peer") ||
		strings.Contains(err.Error(), "context deadline exceeded")
}
