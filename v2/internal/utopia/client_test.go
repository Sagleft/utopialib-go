package utopia

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func getTestClient() *UtopiaClient {
	return NewUtopiaClient(Config{
		Host:                  "localhost",
		Port:                  20000,
		WsPort:                25000,
		RequestTimeoutSeconds: 1,
	})
}

func TestNewUtopiaClient(t *testing.T) {
	c := getTestClient()
	assert.Equal(t, float64(1), c.httpClient.Timeout.Seconds())
}

func TestLimitRate(t *testing.T) {
	c := getTestClient()
	c.limitRate("test")
	c.limitRate(reqDefault)
}

func TestGetProfileStatus(t *testing.T) {
	c := getTestClient()
	_, err := c.GetProfileStatus()
	require.NoError(t, err)
}
