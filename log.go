package utopiago

import (
	"fmt"
	"time"
)

type LogCallback func(logMessage string)

type logData struct {
	TimeCreated time.Time
	Timestamp   int64 // unix milli
	APIURL      string
	APIMethod   string
	RequestType string
	RequestData map[string]interface{}
	Filters     map[string]interface{}
	Response    []byte
	Error       error
	Elapsed     string
}

func (l *logData) setElapsedTime() {
	l.Elapsed = time.Since(l.TimeCreated).String()
}

func (l *logData) useError(err error) error {
	l.Error = err
	l.setElapsedTime()
	return err
}

/*func (l *logData) useResponse(response []byte) {
	l.Response = response
	l.setElapsedTime()
}*/

func (l *logData) getMessage() string {
	status := "success"
	if l.Error != nil {
		status = "error: " + l.Error.Error()
	}

	return fmt.Sprintf(
		"%v: %s %s -> %s. elapsed %s: %s\nrequest: %q\nfilters: %q\nresponse: %q",
		l.Timestamp,
		l.RequestType,
		l.APIURL,
		l.APIMethod,
		l.Elapsed,
		status,
		l.RequestData,
		l.Filters,
		string(l.Response),
	)
}

func (l *logData) handle(cb LogCallback) {
	if cb == nil {
		return
	}

	cb(l.getMessage())
}
