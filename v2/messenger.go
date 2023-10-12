package utopiago

import (
	"github.com/Sagleft/utopialib-go/v2/internal/utopia"
	"github.com/Sagleft/utopialib-go/v2/pkg/structs"
	"github.com/Sagleft/utopialib-go/v2/pkg/websocket"
)

type Client interface {
	// GetProfileStatus gets data about the status of the current account
	GetProfileStatus() (structs.ProfileStatus, error)

	// SetProfileStatus updates data about the status of the current account
	SetProfileStatus(status string, mood string) error

	// SetProfileData - update account name
	SetProfileData(nick, firstName, lastName string) error

	// GetOwnContact asks for full details of the current account
	GetOwnContact() (structs.OwnContactData, error)

	// CheckClientConnection - checks if there are any errors when contacting the client
	CheckClientConnection() bool

	// UseVoucher - uses the voucher and returns an error on failure
	UseVoucher(voucherID string) (string, error)

	// GetFinanceInfo request financial info
	GetFinanceInfo() (structs.FinanceInfo, error)

	// GetFinanceHistory request the necessary financial statistics
	GetFinanceHistory(task structs.GetFinanceHistoryTask) (
		[]structs.FinanceHistoryData,
		error,
	)

	// GetBalance request account Crypton balance
	GetBalance() (float64, error)

	// GetUUSDBalance request account UUSD balance
	GetUUSDBalance() (float64, error)

	// CreateVoucher requests the creation of a new Crypton voucher. it returns referenceNumber
	CreateVoucher(amount float64) (string, error)
	CreateVoucherBatch(amount float64, count int) (string, error)

	// CreateUUSDVoucher requests the creation of a new UUSD voucher. it returns referenceNumber
	CreateUUSDVoucher(amount float64) (string, error)
	CreateUUSDVoucherBatch(amount float64, count int) (string, error)

	// GetWebSocketState - returns WSS Notifications state.
	// 0 - disabled or active listening port number
	GetWebSocketState() (int64, error)

	// SetWebSocketState - set WSS Notification state
	SetWebSocketState(task structs.SetWsStateTask) error

	// WsSubscribe - connect to websocket & receive messages.
	// NOTE: it's blocking method
	WsSubscribe(task websocket.WsSubscribeTask) (websocket.Handler, error)

	// SendChannelMessage - send channel message & get message ID
	SendChannelMessage(channelID, message string) (string, error)

	// SendChannelContactMessage - send channel message to contact in private mode
	SendChannelContactMessage(channelID, contactPubkeyHash, message string) (string, error)

	// SendChannelPicture - send channel picture & get message ID
	SendChannelPicture(channelID, base64Image, comment, filenameForImage string) (string, error)

	// GetStickerNamesByCollection returns available names from corresponded collection
	GetStickerNamesByCollection(collectionName string) ([]string, error)

	// GetStickerImage returns sticker image in base64
	GetStickerImage(collectionName, stickerName string) (string, error)

	// UCodeEncode - encode data to uCode image.
	// coder: BASE64 for example
	// format: JPG or PNG
	UCodeEncode(dataHexCode, coder, format string, imageSize int) (string, error)

	// SendAuthRequest - send auth request to user
	SendAuthRequest(pubkey, message string) (bool, error)

	// AcceptAuthRequest - accept auth request
	AcceptAuthRequest(pubkey, message string) (bool, error)

	// RejectAuthRequest - reject user auth request
	RejectAuthRequest(pubkey, message string) (bool, error)

	// SendInstantMessage - send message to contact (PM).
	// to - pubkey or uNS entry name
	SendInstantMessage(to string, message string) (int64, error)

	// GetContacts - get account contacts.
	// params: filter - contact pubkey or nickname
	GetContacts(filter string) ([]structs.ContactData, error)

	// GetContact data
	GetContact(pubkeyOrNick string) (structs.ContactData, error)

	// JoinChannel - join to channel or chat.
	// password is optional. returns join status (bool) and error
	JoinChannel(channelID string, password ...string) (bool, error)

	// GetChannelContacts - get channel contacts
	GetChannelContacts(channelID string) ([]structs.ChannelContactData, error)

	// EnableChannelReadOnly - toogle channel readonly mode
	EnableChannelReadOnly(channelID string, readOnly bool) error

	// RemoveChannelMessage - remove channel message
	RemoveChannelMessage(channelID string, messageID uint64) error

	// GetChannelMessages - get channel messages with filter (offset, max messages count)
	GetChannelMessages(
		channelID string,
		offset int,
		maxMessages int,
	) ([]structs.ChannelMessage, error)

	// SendPayment - send coins
	SendPayment(task structs.SendPaymentTask) (string, error)

	// GetChannelInfo - get specific channel info
	GetChannelInfo(channelID string) (structs.ChannelData, error)

	// GetChannels get available channels
	GetChannels(task structs.GetChannelsTask) ([]structs.SearchChannelData, error)

	// ToogleChannelNotifications - enable or disable channel notifications
	ToogleChannelNotifications(channelID string, enabled bool) error

	// GetNetworkConnections - get current network peers
	GetNetworkConnections() ([]structs.PeerInfo, error)

	// EnableReadOnly - convert chat to channel
	EnableReadOnly(channelID string, readOnly bool) error

	/*
		GetChannelModeratorRights - find out if the user has moderator rights
		in the channel and get the data about them
	*/
	GetChannelModeratorRights(
		channelID string,
		moderatorPubkey string,
	) (structs.ModeratorRights, error)

	GetChannelModerators(channelID string) ([]string, error)
}

type Config = utopia.Config

func NewUtopiaClient(c Config) Client {
	return utopia.NewUtopiaClient(c)
}
