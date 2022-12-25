package utopia

import (
	"errors"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	mocks "github.com/Sagleft/utopialib-go/v2/internal/mocks"
	"github.com/Sagleft/utopialib-go/v2/pkg/structs"
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

func TestSetProfileStatusUnsuccess(t *testing.T) {
	handlerMock, c := getTestClient(t)

	// when error was given
	handlerMock.EXPECT().Send(gomock.Any(), gomock.Any(), gomock.Any()).
		AnyTimes().Return([]byte(`{"result":false}`), nil)

	// then
	err := c.SetProfileStatus("test", "test")
	require.ErrorIs(t, err, ErrorSetProfileStatus)
}

func TestGetOwnContact(t *testing.T) {
	handlerMock, c := getTestClient(t)

	handlerMock.EXPECT().Send(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		[]byte(`{
			"result": {
				"avatarMd5": "8AFDAB98B48A90F7D3B18AFF96F0852C",
				"hashedPk": "809262B77E2EF657F04C7FA9822426D6",
				"isFriend": false,
				"nick": "contact",
				"pk": "CFF4DB80DCA10BD2317D538FF790A03EDA26274768E5EB04E0FDA51989131F32",
				"status": 4096
			},
			"resultExtraInfo": {
				"elapsed": "0"
			}
		}`), nil,
	)

	contact, err := c.GetOwnContact()
	require.NoError(t, err)

	assert.Equal(t, "contact", contact.Nick)
	assert.Equal(t, 4096, contact.Status)
}

func TestCheckClientConnectionSuccess(t *testing.T) {
	handlerMock, c := getTestClient(t)

	handlerMock.EXPECT().Send(gomock.Any(), gomock.Any(), gomock.Any()).
		Return([]byte(`{"result":{}}`), nil)

	require.True(t, c.CheckClientConnection())
}

func TestCheckClientConnectionUnsuccess(t *testing.T) {
	handlerMock, c := getTestClient(t)

	handlerMock.EXPECT().Send(gomock.Any(), gomock.Any(), gomock.Any()).
		Return(nil, ErrorClientDisconnected)

	require.False(t, c.CheckClientConnection())
}

func TestUseVoucher(t *testing.T) {
	handlerMock, c := getTestClient(t)

	handlerMock.EXPECT().Send(gomock.Any(), gomock.Any(), gomock.Any()).
		Return([]byte(`{"result":{}}`), nil)

	_, err := c.UseVoucher("123-456-789")
	require.NoError(t, err)
}

func TestGetFinanceInfo(t *testing.T) {
	handlerMock, c := getTestClient(t)

	handlerMock.EXPECT().Send(gomock.Any(), gomock.Any(), gomock.Any()).
		Return([]byte(`{"result":{}}`), nil)

	_, err := c.GetFinanceInfo()
	require.NoError(t, err)
}

func TestGetFinanceHistory(t *testing.T) {
	handlerMock, c := getTestClient(t)

	handlerMock.EXPECT().Send(gomock.Any(), gomock.Any(), gomock.Any()).
		Return([]byte(`{"result":[{},{}]}`), nil)

	data, err := c.GetFinanceHistory(structs.GetFinanceHistoryTask{})
	require.NoError(t, err)
	assert.Equal(t, 2, len(data))

	handlerMock.EXPECT().Send(gomock.Any(), gomock.Any(), gomock.Any()).
		Return([]byte(`{"result":[]}`), nil)

	_, err = c.GetFinanceHistory(structs.GetFinanceHistoryTask{
		FromDate: time.Now(),
		ToDate:   time.Now(),
	})
	require.NoError(t, err)
}

func TestGetBalance(t *testing.T) {
	handlerMock, c := getTestClient(t)

	handlerMock.EXPECT().Send(gomock.Any(), gomock.Any(), gomock.Any()).
		Return([]byte(`{"result":0}`), nil)

	_, err := c.GetBalance()
	require.NoError(t, err)
}

func TestGetUUSDBalance(t *testing.T) {
	handlerMock, c := getTestClient(t)

	handlerMock.EXPECT().Send(gomock.Any(), gomock.Any(), gomock.Any()).
		Return([]byte(`{"result":0}`), nil)

	_, err := c.GetUUSDBalance()
	require.NoError(t, err)
}

func TestCreateVoucher(t *testing.T) {
	handlerMock, c := getTestClient(t)

	handlerMock.EXPECT().Send(gomock.Any(), gomock.Any(), gomock.Any()).
		Return([]byte(`{"result":""}`), nil)

	_, err := c.CreateVoucher(100)
	require.NoError(t, err)
}

func TestCreateUUSDVoucher(t *testing.T) {
	handlerMock, c := getTestClient(t)

	handlerMock.EXPECT().Send(gomock.Any(), gomock.Any(), gomock.Any()).
		Return([]byte(`{"result":""}`), nil)

	_, err := c.CreateUUSDVoucher(100)
	require.NoError(t, err)
}

func TestSetWebSocketState(t *testing.T) {
	handlerMock, c := getTestClient(t)

	handlerMock.EXPECT().Send(gomock.Any(), gomock.Any(), gomock.Any()).
		Return([]byte(`{"result":"ok"}`), nil)

	require.NoError(t, c.SetWebSocketState(structs.SetWsStateTask{
		EnableSSL:     true,
		Notifications: "test",
	}))
}

func TestSetWebSocketStateEmptyResult(t *testing.T) {
	handlerMock, c := getTestClient(t)

	handlerMock.EXPECT().Send(gomock.Any(), gomock.Any(), gomock.Any()).
		Return([]byte(`{"result":""}`), nil)

	require.Error(t, c.SetWebSocketState(structs.SetWsStateTask{}))
}

func TestSetWebSocketStateError(t *testing.T) {
	handlerMock, c := getTestClient(t)

	handlerMock.EXPECT().Send(gomock.Any(), gomock.Any(), gomock.Any()).
		Return([]byte(`{}`), nil)

	require.Error(t, c.SetWebSocketState(structs.SetWsStateTask{}))
}

func TestGetWebSocketState(t *testing.T) {
	handlerMock, c := getTestClient(t)

	handlerMock.EXPECT().Send(gomock.Any(), gomock.Any(), gomock.Any()).
		Return([]byte(`{"result":0}`), nil)

	_, err := c.GetWebSocketState()
	require.Nil(t, err)
}

func TestGetWebSocketStateError(t *testing.T) {
	handlerMock, c := getTestClient(t)

	handlerMock.EXPECT().Send(gomock.Any(), gomock.Any(), gomock.Any()).
		Return([]byte(`{}`), nil)

	_, err := c.GetWebSocketState()
	require.Error(t, err)
}

func TestSendChannelMessage(t *testing.T) {
	handlerMock, c := getTestClient(t)

	handlerMock.EXPECT().Send(gomock.Any(), gomock.Any(), gomock.Any()).
		Return([]byte(`{"result":""}`), nil)

	_, err := c.SendChannelMessage("", "")
	require.Nil(t, err)
}

func TestSendChannelContactMessage(t *testing.T) {
	handlerMock, c := getTestClient(t)

	handlerMock.EXPECT().Send(gomock.Any(), gomock.Any(), gomock.Any()).
		Return([]byte(`{"result":""}`), nil)

	_, err := c.SendChannelContactMessage("", "", "")
	require.Nil(t, err)
}

func TestSendChannelPicture(t *testing.T) {
	handlerMock, c := getTestClient(t)

	handlerMock.EXPECT().Send(gomock.Any(), gomock.Any(), gomock.Any()).
		Return([]byte(`{"result":""}`), nil)

	_, err := c.SendChannelPicture("", "", "", "")
	require.Nil(t, err)
}

func TestGetStickerNamesByCollection(t *testing.T) {
	handlerMock, c := getTestClient(t)

	handlerMock.EXPECT().Send(gomock.Any(), gomock.Any(), gomock.Any()).
		Return([]byte(`{"result":[]}`), nil)

	_, err := c.GetStickerNamesByCollection("")
	require.Nil(t, err)
}

func TestGetStickerImage(t *testing.T) {
	handlerMock, c := getTestClient(t)

	handlerMock.EXPECT().Send(gomock.Any(), gomock.Any(), gomock.Any()).
		Return([]byte(`{"result":""}`), nil)

	_, err := c.GetStickerImage("", "")
	require.Nil(t, err)
}

func TestUCodeEncode(t *testing.T) {
	handlerMock, c := getTestClient(t)

	handlerMock.EXPECT().Send(gomock.Any(), gomock.Any(), gomock.Any()).
		Return([]byte(`{"result":""}`), nil)

	_, err := c.UCodeEncode("", "", "", 256)
	require.Nil(t, err)
}

func TestSendAuthRequest(t *testing.T) {
	handlerMock, c := getTestClient(t)

	handlerMock.EXPECT().Send(gomock.Any(), gomock.Any(), gomock.Any()).
		Return([]byte(`{"result":true}`), nil)

	_, err := c.SendAuthRequest("", "")
	require.Nil(t, err)
}

func TestAcceptAuthRequest(t *testing.T) {
	handlerMock, c := getTestClient(t)

	handlerMock.EXPECT().Send(gomock.Any(), gomock.Any(), gomock.Any()).
		Return([]byte(`{"result":true}`), nil)

	_, err := c.AcceptAuthRequest("", "")
	require.Nil(t, err)
}

func TestRejectAuthRequest(t *testing.T) {
	handlerMock, c := getTestClient(t)

	handlerMock.EXPECT().Send(gomock.Any(), gomock.Any(), gomock.Any()).
		Return([]byte(`{"result":true}`), nil)

	_, err := c.RejectAuthRequest("", "")
	require.Nil(t, err)
}
