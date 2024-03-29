package utopiago

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCallbackCheck(t *testing.T) {
	c := UtopiaClient{}
	assert.Nil(t, c.logCallback)
}

func TestLogGetMessage(t *testing.T) {
	l := logData{}
	require.NotPanics(t, func() {
		l.getMessage()
	})
}
