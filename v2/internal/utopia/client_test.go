package utopia

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	mocks "github.com/Sagleft/utopialib-go/v2/internal/mocks"
)

func getTestClient() *UtopiaClient {
	return NewUtopiaClient(Config{
		Port:                  20000,
		WsPort:                25000,
		RequestTimeoutSeconds: 1,
	})
}

func TestNewUtopiaClient(t *testing.T) {
	getTestClient()
}

func TestLimitRate(t *testing.T) {
	c := getTestClient()
	c.limitRate("test")
	c.limitRate(reqDefault)
}

func TestGetProfileStatus(t *testing.T) {
	ctrl := gomock.NewController(t)
	c := getTestClient()
	handlerMock := mocks.NewMockRequestHandler(ctrl)
	c.reqHandler = handlerMock

	handlerMock.EXPECT().Send(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		[]byte(`{
			"result": {
				"mood": "[snowleo]",
				"status": "Available",
				"status_code": 4096
			},
			"resultExtraInfo": {
				"elapsed": "0"
			}
		}`), nil,
	)

	_, err := c.GetProfileStatus()
	require.NoError(t, err)
}
