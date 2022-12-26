package utopia

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLogData(t *testing.T) {
	// when everything is ok
	l := logData{}
	// then
	assert.NotEqual(t, "", l.getMessage())

	// when error is set
	l.Error = errors.New("test error")
	// then
	assert.NotEqual(t, "", l.getMessage())

	// when callback is set
	l.handle(func(logMessage string) {})
}
