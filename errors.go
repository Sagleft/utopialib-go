package utopiago

import (
	"strings"
)

var connBrokenErrorInfo = []string{
	"read: connection reset by peer",
	"context deadline exceeded",
	"client disconected",
	"connection refused",
	"EOF",
}

// CheckErrorConnBroken - check the text of the request error, determining whether the client must be restarted
func CheckErrorConnBroken(err error) bool {
	if err == nil {
		return false
	}

	for _, info := range connBrokenErrorInfo {
		if strings.Contains(err.Error(), info) {
			return true
		}
	}
	return false
}
