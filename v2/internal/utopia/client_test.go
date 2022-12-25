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

func TestSendInstantMessage(t *testing.T) {
	handlerMock, c := getTestClient(t)

	handlerMock.EXPECT().Send(gomock.Any(), gomock.Any(), gomock.Any()).
		Return([]byte(`{"result":0}`), nil)

	_, err := c.SendInstantMessage("", "")
	require.Nil(t, err)
}

func TestGetContacts(t *testing.T) {
	handlerMock, c := getTestClient(t)

	handlerMock.EXPECT().Send(gomock.Any(), gomock.Any(), gomock.Any()).
		Return([]byte(`{"result":[]}`), nil)

	_, err := c.GetContacts("")
	require.Nil(t, err)
}

func TestGetContactsError(t *testing.T) {
	handlerMock, c := getTestClient(t)

	handlerMock.EXPECT().Send(gomock.Any(), gomock.Any(), gomock.Any()).
		Return([]byte(`invalid json`), nil)

	_, err := c.GetContacts("")
	require.Error(t, err)
}

func TestGetContactsError2(t *testing.T) {
	handlerMock, c := getTestClient(t)

	handlerMock.EXPECT().Send(gomock.Any(), gomock.Any(), gomock.Any()).
		Return([]byte(`{}`), nil)

	_, err := c.GetContacts("")
	require.Error(t, err)
}

func TestGetContact(t *testing.T) {
	handlerMock, c := getTestClient(t)

	handlerMock.EXPECT().Send(gomock.Any(), gomock.Any(), gomock.Any()).
		Return([]byte(`{"result": [{}]}`), nil)

	_, err := c.GetContact("")
	require.Nil(t, err)
}

func TestGetContactError(t *testing.T) {
	handlerMock, c := getTestClient(t)

	handlerMock.EXPECT().Send(gomock.Any(), gomock.Any(), gomock.Any()).
		Return([]byte(`invalid json`), nil)

	_, err := c.GetContact("")
	require.Error(t, err)
}

func TestGetContactNotFound(t *testing.T) {
	handlerMock, c := getTestClient(t)

	handlerMock.EXPECT().Send(gomock.Any(), gomock.Any(), gomock.Any()).
		Return([]byte(`{"result": []}`), nil)

	_, err := c.GetContact("")
	require.Error(t, err)
}

func TestJoinChannel(t *testing.T) {
	handlerMock, c := getTestClient(t)

	handlerMock.EXPECT().Send(gomock.Any(), gomock.Any(), gomock.Any()).
		Return([]byte(`{"result": true}`), nil)

	_, err := c.JoinChannel("", "")
	require.Nil(t, err)
}

func TestGetChannelContacts(t *testing.T) {
	handlerMock, c := getTestClient(t)

	handlerMock.EXPECT().Send(gomock.Any(), gomock.Any(), gomock.Any()).
		Return([]byte(`{"result": []}`), nil)

	_, err := c.GetChannelContacts("")
	require.Nil(t, err)
}

func TestGetChannelContactsError(t *testing.T) {
	handlerMock, c := getTestClient(t)

	handlerMock.EXPECT().Send(gomock.Any(), gomock.Any(), gomock.Any()).
		Return([]byte(`{}`), nil)

	_, err := c.GetChannelContacts("")
	require.Error(t, err)
}

func TestGetChannelContactsError2(t *testing.T) {
	handlerMock, c := getTestClient(t)

	handlerMock.EXPECT().Send(gomock.Any(), gomock.Any(), gomock.Any()).
		Return([]byte(`invalid json`), nil)

	_, err := c.GetChannelContacts("")
	require.Error(t, err)
}

func TestEnableChannelReadOnly(t *testing.T) {
	handlerMock, c := getTestClient(t)

	handlerMock.EXPECT().Send(gomock.Any(), gomock.Any(), gomock.Any()).
		Return([]byte(`{"result": true}`), nil)

	require.Nil(t, c.EnableChannelReadOnly("", true))
}

func TestRemoveChannelMessage(t *testing.T) {
	handlerMock, c := getTestClient(t)

	handlerMock.EXPECT().Send(gomock.Any(), gomock.Any(), gomock.Any()).
		Return([]byte(`{"result": ""}`), nil)

	require.Nil(t, c.RemoveChannelMessage("", 1000000))
}

func TestGetChannelMessages(t *testing.T) {
	handlerMock, c := getTestClient(t)

	handlerMock.EXPECT().Send(gomock.Any(), gomock.Any(), gomock.Any()).
		Return([]byte(`{"result": [{}]}`), nil)

	_, err := c.GetChannelMessages("", 0, 1)
	require.Nil(t, err)
}

func TestGetChannelMessagesNotFound(t *testing.T) {
	handlerMock, c := getTestClient(t)

	handlerMock.EXPECT().Send(gomock.Any(), gomock.Any(), gomock.Any()).
		Return([]byte(`{"result": []}`), nil)

	data, err := c.GetChannelMessages("", 0, 1)
	require.Nil(t, err)
	assert.Equal(t, 0, len(data))
}

func TestGetChannelMessagesError(t *testing.T) {
	handlerMock, c := getTestClient(t)

	handlerMock.EXPECT().Send(gomock.Any(), gomock.Any(), gomock.Any()).
		Return([]byte(`invalid json`), nil)

	_, err := c.GetChannelMessages("", 0, 1)
	require.Error(t, err)

	handlerMock.EXPECT().Send(gomock.Any(), gomock.Any(), gomock.Any()).
		Return([]byte(`{}`), nil)

	_, err = c.GetChannelMessages("", 0, 1)
	require.Error(t, err)
}

func TestSendPayment(t *testing.T) {
	handlerMock, c := getTestClient(t)

	handlerMock.EXPECT().Send(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().
		Return([]byte(`{"result": ""}`), nil)

	// when comment is too long
	_, err := c.SendPayment(structs.SendPaymentTask{
		To:     "pubkey",
		Amount: 100,
		Comment: "XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX" +
			"XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX" +
			"XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
	})
	require.Error(t, err)

	// when `amount` is not set
	task := structs.SendPaymentTask{}
	// then
	_, err = c.SendPayment(task)
	require.Error(t, err)

	// when `to` is not set
	task = structs.SendPaymentTask{
		Amount: 100,
	}
	_, err = c.SendPayment(task)
	require.Error(t, err)

	// when everything is ok
	task = structs.SendPaymentTask{
		To:     "pubkey",
		Amount: 100,
	}
	_, err = c.SendPayment(task)
	require.Nil(t, err)
}

func TestGetChannelInfo(t *testing.T) {
	handlerMock, c := getTestClient(t)

	handlerMock.EXPECT().Send(gomock.Any(), gomock.Any(), gomock.Any()).
		Return([]byte(`{"result": {}}`), nil)

	_, err := c.GetChannelInfo("")
	require.Nil(t, err)
}

func TestGetChannelInfoError(t *testing.T) {
	handlerMock, c := getTestClient(t)

	handlerMock.EXPECT().Send(gomock.Any(), gomock.Any(), gomock.Any()).
		Return([]byte(`{"result": false}`), nil)

	_, err := c.GetChannelInfo("")
	require.Error(t, err)

	handlerMock.EXPECT().Send(gomock.Any(), gomock.Any(), gomock.Any()).
		Return([]byte(`invalid json`), nil)

	_, err = c.GetChannelInfo("")
	require.Error(t, err)
}
