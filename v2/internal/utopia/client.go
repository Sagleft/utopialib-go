package utopia

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/Sagleft/utopialib-go/v2/pkg/consts"
	"github.com/Sagleft/utopialib-go/v2/pkg/structs"
)

func NewUtopiaClient(data Config) *UtopiaClient {
	return &UtopiaClient{
		data: data,
	}
}

func (c *UtopiaClient) GetProfileStatus() (structs.ProfileStatus, error) {
	r := structs.ProfileStatus{}
	err := c.getSimpleStruct("getProfileStatus", &r)
	return r, err
}

// GetSystemInfo retrieves client system information
func (c *UtopiaClient) GetSystemInfo() (structs.SystemInfo, error) {
	r := structs.SystemInfo{}
	err := c.getSimpleStruct("getSystemInfo", &r)
	return r, err
}

func (c *UtopiaClient) SetProfileStatus(status string, mood string) error {
	result, err := c.queryResultToBool("setProfileStatus", uMap{
		"status": status,
		"mood":   mood,
	})
	if err != nil {
		return err
	}
	if !result {
		return errors.New("failed to set profile status")
	}
	return nil
}

func (c *UtopiaClient) GetOwnContact() (structs.OwnContactData, error) {
	r := structs.OwnContactData{}
	err := c.getSimpleStruct("getOwnContact", &r)
	return r, err
}

func (c *UtopiaClient) CheckClientConnection() bool {
	_, err := c.GetSystemInfo()
	return err == nil
}

func (c *UtopiaClient) UseVoucher(voucherID string) (string, error) {
	return c.queryResultToString("useVoucher", uMap{"voucherid": voucherID})
}

func (c *UtopiaClient) GetFinanceInfo() (structs.FinanceInfo, error) {
	r := structs.FinanceInfo{}
	err := c.getSimpleStruct("getFinanceSystemInformation", &r)
	return r, err
}

func (c *UtopiaClient) GetFinanceHistory(task structs.GetFinanceHistoryTask) (
	[]structs.FinanceHistoryData,
	error,
) {
	params := newMapBuilder().
		add(task.Currency, "", "currency").
		add(task.Filters, "", "filters").
		add(task.ReferenceNumber, "", "referenceNumber").
		add(task.BatchID, 0, "batchId").
		add(task.FromAmount, 0, "fromAmount").
		add(task.ToAmount, 0, "toAmount").
		add(task.SourcePubkey, "", "sourcePk").
		add(task.DestinationPubkey, "", "destinationPk").
		getMap()

	if !task.FromDate.IsZero() {
		params["fromDate"] = task.FromDate.Format(defaultTimeLayout)
	}
	if !task.ToDate.IsZero() {
		params["toDate"] = task.ToDate.Format(defaultTimeLayout)
	}

	filters := newMapBuilder().
		add(task.QueryOffset, 0, "offset").
		add(task.QueryLimitRows, 0, "limitRows").
		getMap()

	r := []structs.FinanceHistoryData{}
	err := c.retrieveStruct("getFinanceHistory", params, filters, &r)
	return r, err
}

func (c *UtopiaClient) GetBalance() (float64, error) {
	return c.queryResultToFloat64("getBalance", uMap{})
}

// GetUUSDBalance request account UUSD balance
func (c *UtopiaClient) GetUUSDBalance() (float64, error) {
	params := uMap{
		"currency": "USD",
	}
	result, err := c.queryResultToFloat64("getBalance", params)
	if err != nil {
		return 0, err
	}
	return result, nil
}

func (c *UtopiaClient) createCoinVoucher(amount float64, coin string) (string, error) {
	params := uMap{
		"amount":   amount,
		"currency": coin,
	}
	result, err := c.queryResultToString("createVoucher", params)
	if err != nil {
		return "", err
	}
	if result == "" {
		return "", errors.New("failed to create voucher, empty string in client response")
	}
	return result, nil
}

// CreateVoucher requests the creation of a new Crypton voucher. it returns referenceNumber
func (c *UtopiaClient) CreateVoucher(amount float64) (string, error) {
	return c.createCoinVoucher(amount, "CRP")
}

// CreateUUSDVoucher requests the creation of a new UUSD voucher. it returns referenceNumber
func (c *UtopiaClient) CreateUUSDVoucher(amount float64) (string, error) {
	return c.createCoinVoucher(amount, "UUSD")
}

// SetWebSocketState - set WSS Notification state
func (c *UtopiaClient) SetWebSocketState(task structs.SetWsStateTask) error {
	params := uMap{
		"enabled": strconv.FormatBool(task.Enabled),
		"port":    strconv.Itoa(task.Port),
	}
	if task.EnableSSL {
		params["enablessl"] = strconv.FormatBool(task.EnableSSL)
	}
	if task.Notifications != "" {
		params["notifications"] = task.Notifications
	}

	result, err := c.queryResultToString("setWebSocketState", params)
	if err != nil {
		return err
	}
	if result == "" {
		return errors.New("failed to set websocker state")
	}
	return nil
}

// GetWebSocketState - returns WSS Notifications state, 0 - disabled or active listening port number.
func (c *UtopiaClient) GetWebSocketState() (int64, error) {
	result, err := c.queryResultToInt("getWebSocketState", nil)
	if err != nil {
		return 0, err
	}
	return result, nil
}

// SendChannelMessage - send channel message & get message ID
func (c *UtopiaClient) SendChannelMessage(channelID, message string) (string, error) {
	params := uMap{
		"channelid": channelID,
		"message":   message,
	}
	return c.queryResultToString("sendChannelMessage", params)
}

// SendChannelContactMessage - send channel message to contact in private mode
func (c *UtopiaClient) SendChannelContactMessage(channelID, contactPubkeyHash, message string) (string, error) {
	params := uMap{
		"channelid":       channelID,
		"contactHashedPk": contactPubkeyHash,
		"message":         message,
	}
	return c.queryResultToString("sendChannelPrivateMessageToContact", params)
}

// SendChannelPicture - send channel picture & get message ID
func (c *UtopiaClient) SendChannelPicture(channelID, base64Image, comment, filenameForImage string) (string, error) {
	params := uMap{
		"channelid":      channelID,
		"base64_image":   base64Image,
		"comment":        comment,
		"filename_image": filenameForImage,
	}
	return c.queryResultToString("sendChannelPicture", params)
}

// GetStickerNamesByCollection returns available names from corresponded collection
func (c *UtopiaClient) GetStickerNamesByCollection(collectionName string) ([]string, error) {
	params := uMap{
		"collection_name": collectionName,
	}
	return c.queryResultToStringsArray("getStickerNamesByCollection", params)
}

// GetStickerImage returns sticker image in base64
func (c *UtopiaClient) GetStickerImage(collectionName, stickerName string) (string, error) {
	params := uMap{
		"collection_name": collectionName,
		"sticker_name":    stickerName,
		"coder":           "BASE64",
	}
	return c.queryResultToString("getImageSticker", params)
}

// UCodeEncode - encode data to uCode image.
// coder: BASE64 for example
// format: JPG or PNG
func (c *UtopiaClient) UCodeEncode(dataHexCode, coder, format string, imageSize int) (string, error) {
	return c.queryResultToString("ucodeEncode", uMap{
		"hex_code":   dataHexCode,
		"size_image": imageSize,
		"coder":      "BASE64",
		"format":     format,
	})
}

// SendAuthRequest - send auth request to user
func (c *UtopiaClient) SendAuthRequest(pubkey, message string) (bool, error) {
	params := uMap{
		"pk":      pubkey,
		"message": message,
	}
	return c.queryResultToBool("sendAuthorizationRequest", params)
}

// AcceptAuthRequest - accept auth request
func (c *UtopiaClient) AcceptAuthRequest(pubkey, message string) (bool, error) {
	params := uMap{
		"pk":      pubkey,
		"message": message,
	}
	return c.queryResultToBool("acceptAuthorizationRequest", params)
}

// RejectAuthRequest - reject user auth request
func (c *UtopiaClient) RejectAuthRequest(pubkey, message string) (bool, error) {
	params := uMap{
		"pk":      pubkey,
		"message": message,
	}
	return c.queryResultToBool("rejectAuthorizationRequest", params)
}

// SendInstantMessage - send message to contact (PM).
// to -- pubkey or uNS entry name
func (c *UtopiaClient) SendInstantMessage(to string, message string) (int64, error) {
	params := uMap{
		"to":   to,
		"text": message,
	}
	return c.queryResultToInt("sendInstantMessage", params)
}

// GetContacts - get account contacts.
// params: filter - contact pubkey or nickname
func (c *UtopiaClient) GetContacts(filter string) ([]structs.ContactData, error) {
	// send request
	params := uMap{}
	if filter != "" {
		params["filter"] = filter
	}
	response, err := c.apiQuery("getContacts", params)
	if err != nil {
		return nil, err
	}

	data := []structs.ContactData{}
	if err := convertResult(response, &data); err != nil {
		return nil, err
	}
	return data, nil
}

// GetContact data
func (c *UtopiaClient) GetContact(pubkeyOrNick string) (structs.ContactData, error) {
	contacts, err := c.GetContacts(pubkeyOrNick)
	if err != nil {
		return structs.ContactData{}, err
	}

	if len(contacts) == 0 {
		return structs.ContactData{}, errors.New("contact not found")
	}
	return contacts[0], nil
}

// JoinChannel - join to channel or chat.
// password is optional. returns join status (bool) and error
func (c *UtopiaClient) JoinChannel(channelID string, password ...string) (bool, error) {
	params := uMap{
		"ident": channelID,
	}
	if len(password) > 0 {
		params["password"] = password[0]
	}
	return c.queryResultToBool("joinChannel", params)
}

// GetChannelContacts - get channel contacts
func (c *UtopiaClient) GetChannelContacts(channelID string) ([]structs.ChannelContactData, error) {
	// send request
	params := uMap{
		"channelid": channelID,
	}
	response, err := c.apiQuery("getChannelContacts", params)
	if err != nil {
		return nil, err
	}

	data := []structs.ChannelContactData{}
	if err := convertResult(response, &data); err != nil {
		return nil, err
	}
	return data, nil
}

func (c *UtopiaClient) EnableReadOnly(channelID string, readOnly bool) error {
	_, err := c.queryResultToBool("modifyChannel", uMap{
		"channelid": channelID,
		"read_only": readOnly,
	})
	return err
}

// RemoveChannelMessage - remove channel message
func (c *UtopiaClient) RemoveChannelMessage(channelID string, messageID int64) error {
	params := uMap{
		"channelid":  channelID,
		"id_message": messageID,
	}
	_, err := c.queryResultToString("removeChannelMessage", params)
	return err
}

// GetChannelMessages - get channel messages with filter (offset, max messages count)
func (c *UtopiaClient) GetChannelMessages(
	channelID string,
	offset int,
	maxMessages int,
) ([]structs.ChannelMessage, error) {
	// send request
	params := uMap{
		"channelid": channelID,
	}
	filters := uMap{
		"offset": offset,
		"limit":  maxMessages,
	}
	response, err := c.apiQueryWithFilters("getChannelMessages", params, filters)
	if err != nil {
		return nil, err
	}

	data := []structs.ChannelMessage{}
	if err := convertResult(response, &data); err != nil {
		return nil, err
	}
	return data, nil
}

type SendPaymentTask struct {
	// required
	To     string  `json:"to"`     // pubkey, nickname or card ID
	Amount float64 `json:"amount"` // more than zero, no more than 9 decimal places

	// optional
	CurrencyTag string `json:"currency"`   // example: "CRP", "UUSD". by default: "CRP"
	FromCardID  string `json:"fromCardID"` // specify here your card ID
	Comment     string `json:"comment"`
}

// SendPayment - send coins
func (c *UtopiaClient) SendPayment(task SendPaymentTask) (string, error) {
	if task.Comment != "" && len(task.Comment) > maxCharactersInPaymentComment {
		return "", fmt.Errorf("comment max length is %v characters", maxCharactersInPaymentComment)
	}

	if task.CurrencyTag == "" {
		task.CurrencyTag = defaultCurrencyTag
	}

	params := uMap{
		"to":       task.To,
		"comment":  task.Comment,
		"cardid":   task.FromCardID,
		"amount":   task.Amount,
		"currency": task.CurrencyTag,
	}
	return c.queryResultToString("sendPayment", params)
}

// GetChannelInfo - get specific channel info
func (c *UtopiaClient) GetChannelInfo(channelID string) (structs.ChannelData, error) {
	params := uMap{
		"channelid": channelID,
	}
	response, err := c.apiQuery("getChannelInfo", params)
	if err != nil {
		return structs.ChannelData{}, err
	}

	data := structs.ChannelData{}
	if err := convertResult(response, &data); err != nil {
		return structs.ChannelData{}, err
	}

	return data, nil
}

type GetChannelsTask struct {
	// optional
	SearchFilter string             // part of channel name or channel ID, etc
	ChannelType  consts.ChannelType // by default: 0 - registered
	FromDate     string             // date example: 2019-11-23T10:00:00.001
	ToDate       string
	SortBy       consts.SortChannelsBy
}

// GetChannels get available channels
func (c *UtopiaClient) GetChannels(task GetChannelsTask) ([]structs.SearchChannelData, error) {
	params := uMap{
		"filter":       task.SearchFilter,
		"channel_type": task.ChannelType,
	}
	if task.FromDate != "" {
		params["from"] = task.FromDate
	}
	if task.ToDate != "" {
		params["to"] = task.ToDate
	}

	filters := uMap{}
	switch task.SortBy {
	case consts.SortChannelsByCreated:
		filters["sortBy"] = "created"
	case consts.SortChannelsByIsPrivate:
		filters["sortBy"] = "isprivate"
	case consts.SortChannelsByName:
		filters["sortBy"] = "name"
	case consts.SortChannelsByModified:
		filters["sortBy"] = "modified"
	case consts.SortChannelsByDescription:
		filters["sortBy"] = "description"
	}

	response, err := c.apiQueryWithFilters("getChannels", params, filters)
	if err != nil {
		return nil, err
	}

	data := []structs.SearchChannelData{}
	if err := convertResult(response, &data); err != nil {
		return nil, err
	}

	return data, nil
}

func (c *UtopiaClient) ToogleChannelNotifications(channelID string, enabled bool) error {
	params := uMap{
		"channelid": channelID,
		"enabled":   enabled,
	}

	if _, err := c.queryResultToBool("enableChannelNotification", params); err != nil {
		return err
	}
	return nil
}
