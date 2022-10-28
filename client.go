package utopiago

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
)

const (
	MaxCharactersInPaymentComment = 148
	DefaultCurrencyTag            = "CRP"
)

const (
	ChannelTypeRegistered ChannelType = iota
	ChannelTypeRecent
	ChannelTypeMy
	ChannelTypeFriends
	ChannelTypeBookmarked
	ChannelTypeJoined
	ChannelTypeOpened
	ChannelTypeBlackList
	ChannelTypeDeleted
)

type ChannelType int

// NewClient - create client with default data:
// http protocol, localhost, default port (20000)
func NewClient(token string) *UtopiaClient {
	return &UtopiaClient{
		Protocol: "http",
		Host:     "localhost",
		Token:    token,
		Port:     20000,
		WsPort:   25000,
	}
}

func (c *UtopiaClient) SetToken(token string) *UtopiaClient {
	c.Token = token
	return c
}

func (c *UtopiaClient) SetProtocol(proto string) *UtopiaClient {
	c.Protocol = proto
	return c
}

func (c *UtopiaClient) SetPort(port int) *UtopiaClient {
	c.Port = port
	return c
}

func (c *UtopiaClient) SetWsPort(wsPort int) *UtopiaClient {
	c.WsPort = wsPort
	return c
}

func (c *UtopiaClient) SetLogsCallback(cb LogCallback) {
	c.logCallback = cb
}

// GetProfileStatus gets data about the status of the current account
func (c *UtopiaClient) GetProfileStatus() (map[string]interface{}, error) {
	return c.apiQuery("getProfileStatus", nil)
}

// GetSystemInfo retrieves client system information
func (c *UtopiaClient) GetSystemInfo() (map[string]interface{}, error) {
	return c.apiQuery("getSystemInfo", nil)
}

// SetProfileStatus updates data about the status of the current account
func (c *UtopiaClient) SetProfileStatus(status string, mood string) error {
	queryMap := make(map[string]interface{})
	queryMap["status"] = status
	queryMap["mood"] = mood

	result, err := c.queryResultToBool("setProfileStatus", queryMap)
	if err != nil {
		return err
	}
	if !result {
		return errors.New("failed to set profile status")
	}
	return nil
}

// GetOwnContact asks for full details of the current account
func (c *UtopiaClient) GetOwnContact() (OwnContactData, error) {
	response, err := c.apiQuery("getOwnContact", nil)
	if err != nil {
		return OwnContactData{}, err
	}

	// check result exists
	result, isResultFound := response["result"]
	if !isResultFound {
		return OwnContactData{}, errors.New("accaptable result doesn't exists in client response")
	}

	// convert result
	jsonBytes, err := json.Marshal(result)
	if err != nil {
		return OwnContactData{}, errors.New("failed to encode response result: " + err.Error())
	}

	ownContact := OwnContactData{}
	err = json.Unmarshal(jsonBytes, &ownContact)
	if err != nil {
		return OwnContactData{}, errors.New("failed to decode reconverted result: " + err.Error())
	}

	return ownContact, nil
}

// CheckClientConnection - checks if there are any errors when contacting the client
func (c *UtopiaClient) CheckClientConnection() bool {
	_, err := c.GetSystemInfo()
	return err == nil
}

// UseVoucher - uses the voucher and returns an error on failure
func (c *UtopiaClient) UseVoucher(voucherID string) (string, error) {
	params := map[string]interface{}{
		"voucherid": voucherID,
	}
	return c.queryResultToString("useVoucher", params)
}

// GetFinanceHistory request the necessary financial statistics
func (c *UtopiaClient) GetFinanceHistory(filters string, referenceNumber string) ([]interface{}, error) {
	params := map[string]interface{}{
		"filters":         filters,
		"referenceNumber": referenceNumber,
	}
	return c.queryResultToInterfaceArray("getFinanceSystemInformation", params)
}

// GetBalance request account Crypton balance
func (c *UtopiaClient) GetBalance() (float64, error) {
	result, err := c.queryResultToFloat64("getBalance", map[string]interface{}{})
	if err != nil {
		return 0, err
	}
	return result, nil
}

// GetUUSDBalance request account UUSD balance
func (c *UtopiaClient) GetUUSDBalance() (float64, error) {
	params := map[string]interface{}{
		"currency": "USD",
	}
	result, err := c.queryResultToFloat64("getBalance", params)
	if err != nil {
		return 0, err
	}
	return result, nil
}

func (c *UtopiaClient) createCoinVoucher(amount float64, coin string) (string, error) {
	params := map[string]interface{}{
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
func (c *UtopiaClient) SetWebSocketState(task SetWsStateTask) error {
	params := map[string]interface{}{
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
	params := map[string]interface{}{
		"channelid": channelID,
		"message":   message,
	}
	return c.queryResultToString("sendChannelMessage", params)
}

// SendChannelContactMessage - send channel message to contact in private mode
func (c *UtopiaClient) SendChannelContactMessage(channelID, contactPubkeyHash, message string) (string, error) {
	params := map[string]interface{}{
		"channelid":       channelID,
		"contactHashedPk": contactPubkeyHash,
		"message":         message,
	}
	return c.queryResultToString("sendChannelPrivateMessageToContact", params)
}

// SendChannelPicture - send channel picture & get message ID
func (c *UtopiaClient) SendChannelPicture(channelID, base64Image, comment, filenameForImage string) (string, error) {
	params := map[string]interface{}{
		"channelid":      channelID,
		"base64_image":   base64Image,
		"comment":        comment,
		"filename_image": filenameForImage,
	}
	return c.queryResultToString("sendChannelPicture", params)
}

// GetStickerNamesByCollection returns available names from corresponded collection
func (c *UtopiaClient) GetStickerNamesByCollection(collectionName string) ([]string, error) {
	params := map[string]interface{}{
		"collection_name": collectionName,
	}
	return c.queryResultToStringsArray("getStickerNamesByCollection", params)
}

// GetStickerImage returns sticker image in base64
func (c *UtopiaClient) GetStickerImage(collectionName, stickerName string) (string, error) {
	params := map[string]interface{}{
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
	return c.queryResultToString("ucodeEncode", map[string]interface{}{
		"hex_code":   dataHexCode,
		"size_image": imageSize,
		"coder":      "BASE64",
		"format":     format,
	})
}

// SendAuthRequest - send auth request to user
func (c *UtopiaClient) SendAuthRequest(pubkey, message string) (bool, error) {
	params := map[string]interface{}{
		"pk":      pubkey,
		"message": message,
	}
	return c.queryResultToBool("sendAuthorizationRequest", params)
}

// AcceptAuthRequest - accept auth request
func (c *UtopiaClient) AcceptAuthRequest(pubkey, message string) (bool, error) {
	params := map[string]interface{}{
		"pk":      pubkey,
		"message": message,
	}
	return c.queryResultToBool("acceptAuthorizationRequest", params)
}

// RejectAuthRequest - reject user auth request
func (c *UtopiaClient) RejectAuthRequest(pubkey, message string) (bool, error) {
	params := map[string]interface{}{
		"pk":      pubkey,
		"message": message,
	}
	return c.queryResultToBool("rejectAuthorizationRequest", params)
}

// SendInstantMessage - send message to contact (PM).
// to -- pubkey or uNS entry name
func (c *UtopiaClient) SendInstantMessage(to string, message string) (int64, error) {
	params := map[string]interface{}{
		"to":   to,
		"text": message,
	}
	return c.queryResultToInt("sendInstantMessage", params)
}

// GetContacts - get account contacts.
// params: filter - contact pubkey or nickname
func (c *UtopiaClient) GetContacts(filter string) ([]ContactData, error) {
	// send request
	params := map[string]interface{}{}
	if filter != "" {
		params["filter"] = filter
	}
	response, err := c.apiQuery("getContacts", params)
	if err != nil {
		return nil, err
	}

	data := []ContactData{}
	if err := convertResult(response, &data); err != nil {
		return nil, err
	}
	return data, nil
}

// GetContact data
func (c *UtopiaClient) GetContact(pubkeyOrNick string) (ContactData, error) {
	contacts, err := c.GetContacts(pubkeyOrNick)
	if err != nil {
		return ContactData{}, err
	}

	if len(contacts) == 0 {
		return ContactData{}, errors.New("contact not found")
	}
	return contacts[0], nil
}

// IsOnline - is contact online?
func (d *ContactData) IsOnline() bool {
	return d.Status == 4096
}

// IsOffline - is contact offline?
func (d *ContactData) IsOffline() bool {
	return d.Status == 65536
}

// IsAway - is contact away?
func (d *ContactData) IsAway() bool {
	return d.Status == 4097
}

// IsDoNotDisturb - is contact marked to do not disturb mode?
func (d *ContactData) IsDoNotDisturb() bool {
	return d.Status == 4099
}

// IsInvisible - is contact invisible?
func (d *ContactData) IsInvisible() bool {
	return d.Status == 32768
}

// JoinChannel - join to channel or chat.
// password is optional. returns join status (bool) and error
func (c *UtopiaClient) JoinChannel(channelID string, password ...string) (bool, error) {
	params := map[string]interface{}{
		"ident": channelID,
	}
	if len(password) > 0 {
		params["password"] = password[0]
	}
	return c.queryResultToBool("joinChannel", params)
}

// GetChannelContacts - get channel contacts
func (c *UtopiaClient) GetChannelContacts(channelID string) ([]ChannelContactData, error) {
	// send request
	params := map[string]interface{}{
		"channelid": channelID,
	}
	response, err := c.apiQuery("getChannelContacts", params)
	if err != nil {
		return nil, err
	}

	data := []ChannelContactData{}
	if err := convertResult(response, &data); err != nil {
		return nil, err
	}
	return data, nil
}

func (c *UtopiaClient) EnableReadOnly(channelID string, readOnly bool) error {
	_, err := c.queryResultToBool("modifyChannel", map[string]interface{}{
		"channelid": channelID,
		"read_only": readOnly,
	})
	return err
}

// RemoveChannelMessage - remove channel message
func (c *UtopiaClient) RemoveChannelMessage(channelID string, messageID int64) error {
	params := map[string]interface{}{
		"channelid":  channelID,
		"id_message": messageID,
	}
	_, err := c.queryResultToString("removeChannelMessage", params)
	return err
}

// GetChannelMessages - get channel messages with filter (offset, max messages count)
func (c *UtopiaClient) GetChannelMessages(channelID string, offset int, maxMessages int) ([]ChannelMessage, error) {
	// send request
	params := map[string]interface{}{
		"channelid": channelID,
	}
	filters := map[string]interface{}{
		"offset": offset,
		"limit":  maxMessages,
	}
	response, err := c.apiQueryWithFilters("getChannelMessages", params, filters)
	if err != nil {
		return nil, err
	}

	data := []ChannelMessage{}
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
	if task.Comment != "" && len(task.Comment) > MaxCharactersInPaymentComment {
		return "", fmt.Errorf("comment max length is %v characters", MaxCharactersInPaymentComment)
	}

	if task.CurrencyTag == "" {
		task.CurrencyTag = DefaultCurrencyTag
	}

	params := map[string]interface{}{
		"to":       task.To,
		"comment":  task.Comment,
		"cardid":   task.FromCardID,
		"amount":   task.Amount,
		"currency": task.CurrencyTag,
	}
	return c.queryResultToString("sendPayment", params)
}

// GetChannelInfo - get specific channel info
func (c *UtopiaClient) GetChannelInfo(channelID string) (ChannelData, error) {
	params := map[string]interface{}{
		"channelid": channelID,
	}
	response, err := c.apiQuery("getChannelInfo", params)
	if err != nil {
		return ChannelData{}, err
	}

	data := ChannelData{}
	if err := convertResult(response, &data); err != nil {
		return ChannelData{}, err
	}

	return data, nil
}

type GetChannelsTask struct {
	// optional
	SearchFilter string      // part of channel name or channel ID, etc
	ChannelType  ChannelType // by default: 0 - registered
	FromDate     string      // date example: 2019-11-23T10:00:00.001
	ToDate       string
}

// GetChannels get available channels
func (c *UtopiaClient) GetChannels(task GetChannelsTask) (SearchChannelData, error) {
	params := map[string]interface{}{
		"filter":       task.SearchFilter,
		"channel_type": task.ChannelType,
	}
	if task.FromDate != "" {
		params["from"] = task.FromDate
	}
	if task.ToDate != "" {
		params["to"] = task.ToDate
	}

	response, err := c.apiQuery("getChannels", params)
	if err != nil {
		return SearchChannelData{}, err
	}

	data := SearchChannelData{}
	if err := convertResult(response, &data); err != nil {
		return SearchChannelData{}, err
	}

	return data, nil
}
