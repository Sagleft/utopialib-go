package utopia

import (
	"errors"
	"time"
)

const (
	maxCharactersInPaymentComment = 148
	defaultCurrencyTag            = "CRP"
	defaultPort                   = 20000
	defaultWsPort                 = 25000
	defaultHost                   = "127.0.0.1"
	defaultProtocol               = "http"
	defaultTimeLayout             = time.RFC3339
	defaultRequestsPerSecond      = 5

	reqDefault                     = "default"
	reqGetProfileStatus            = "getProfileStatus"
	reqGetSystemInfo               = "getSystemInfo"
	reqSetProfileStatus            = "setProfileStatus"
	reqGetOwnContact               = "getOwnContact"
	reqUseVoucher                  = "useVoucher"
	reqGetFinanceSystemInformation = "getFinanceSystemInformation"
	reqGetFinanceHistory           = "getFinanceHistory"
	reqGetChannels                 = "getChannels"
	reqGetChannelInfo              = "getChannelInfo"
	reqJoinChannel                 = "joinChannel"
	reqGetBalance                  = "getBalance"
	reqCreateVoucher               = "createVoucher"
	reqSetWebSocketState           = "setWebSocketState"
	reqGetWebSocketState           = "getWebSocketState"
	reqSendChannelMessage          = "sendChannelMessage"
	reqSendPrivateChannelMessage   = "sendChannelPrivateMessageToContact"
	reqSendChannelPicture          = "sendChannelPicture"
	reqGetStickerNamesByCollection = "getStickerNamesByCollection"
	reqGetImageSticker             = "getImageSticker"
	reqUcodeEncode                 = "ucodeEncode"
	reqGetChannelContacts          = "getChannelContacts"
)

var (
	ErrorSetProfileStatus   = errors.New("failed to set profile status")
	ErrorClientDisconnected = errors.New("client disconected")
)
