package utopia

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetWsURL(t *testing.T) {
	_, c := getTestClient(t)
	assert.NotEqual(t, "", c.getWsURL())
}
