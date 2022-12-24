package utopia

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	mocks "github.com/Sagleft/utopialib-go/v2/internal/mocks"
)

func getTestClient(t *testing.T) (*mocks.MockRequestHandler, *UtopiaClient) {
	ctrl := gomock.NewController(t)
	handlerMock := mocks.NewMockRequestHandler(ctrl)

	c := NewUtopiaClient(Config{
		Port:                  20000,
		WsPort:                25000,
		RequestTimeoutSeconds: 1,
	})
	c.reqHandler = handlerMock
	return handlerMock, c
}

func TestNewUtopiaClient(t *testing.T) {
	getTestClient(t)
}

func TestLimitRate(t *testing.T) {
	_, c := getTestClient(t)
	c.limitRate("test")
	c.limitRate(reqDefault)
}

func TestGetProfileStatus(t *testing.T) {
	handlerMock, c := getTestClient(t)

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

func TestGetSystemInfo(t *testing.T) {
	handlerMock, c := getTestClient(t)

	handlerMock.EXPECT().Send(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		[]byte(`{"result": {}}`), nil,
	)

	_, err := c.GetSystemInfo()
	require.NoError(t, err)
}

func TestSetProfileStatusNoError(t *testing.T) {
	handlerMock, c := getTestClient(t)

	// when all is ok
	handlerMock.EXPECT().Send(gomock.Any(), gomock.Any(), gomock.Any()).
		AnyTimes().Return([]byte(`{"result": {}}`), nil)

	// then
	require.NoError(t, c.SetProfileStatus("test", "test"))
}

func TestSetProfileStatusError(t *testing.T) {
	handlerMock, c := getTestClient(t)

	// when error was given
	handlerMock.EXPECT().Send(gomock.Any(), gomock.Any(), gomock.Any()).
		AnyTimes().Return(nil, errors.New("test error"))

	// then
	require.Error(t, c.SetProfileStatus("test", "test"))
}
