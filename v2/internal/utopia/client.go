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
	params := uMap{}.
		add("currency", task.Currency).
		add("filters", task.Filters).
		add("referenceNumber", task.ReferenceNumber).
		add("batchId", task.BatchID).
		add("fromAmount", task.FromAmount).
		add("toAmount", task.ToAmount).
		add("sourcePk", task.SourcePubkey).
		add("destinationPk", task.DestinationPubkey)

	if !task.FromDate.IsZero() {
		params["fromDate"] = task.FromDate.Format(defaultTimeLayout)
	}
	if !task.ToDate.IsZero() {
		params["toDate"] = task.ToDate.Format(defaultTimeLayout)
	}

	filters := uMap{}.
		add("offset", task.QueryOffset).
		add("limitRows", task.QueryLimitRows)

	r := []structs.FinanceHistoryData{}
	err := c.retrieveStruct("getFinanceHistory", params, filters, &r)
	return r, err
}

func (c *UtopiaClient) GetBalance() (float64, error) {
	return c.queryResultToFloat64("getBalance", uMap{})
}

func (c *UtopiaClient) GetUUSDBalance() (float64, error) {
	return c.queryResultToFloat64("getBalance", uMap{"currency": "USD"})
}

func (c *UtopiaClient) createCoinVoucher(amount float64, coin string) (string, error) {
	params := uMap{}.set("amount", amount).set("currency", coin)
	return c.queryResultToString("createVoucher", params)
}

func (c *UtopiaClient) CreateVoucher(amount float64) (string, error) {
	return c.createCoinVoucher(amount, "CRP")
}

func (c *UtopiaClient) CreateUUSDVoucher(amount float64) (string, error) {
	return c.createCoinVoucher(amount, "UUSD")
}

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

func (c *UtopiaClient) GetWebSocketState() (int64, error) {
	result, err := c.queryResultToInt("getWebSocketState", nil)
	if err != nil {
		return 0, err
	}
	return result, nil
}

func (c *UtopiaClient) SendChannelMessage(channelID, message string) (string, error) {
	params := uMap{
		"channelid": channelID,
		"message":   message,
	}
	return c.queryResultToString("sendChannelMessage", params)
}

func (c *UtopiaClient) SendChannelContactMessage(channelID, contactPubkeyHash, message string) (string, error) {
	params := uMap{
		"channelid":       channelID,
		"contactHashedPk": contactPubkeyHash,
		"message":         message,
	}
	return c.queryResultToString("sendChannelPrivateMessageToContact", params)
}

func (c *UtopiaClient) SendChannelPicture(channelID, base64Image, comment, filenameForImage string) (string, error) {
	params := uMap{
		"channelid":      channelID,
		"base64_image":   base64Image,
		"comment":        comment,
		"filename_image": filenameForImage,
	}
	return c.queryResultToString("sendChannelPicture", params)
}

func (c *UtopiaClient) GetStickerNamesByCollection(collectionName string) ([]string, error) {
	params := uMap{
		"collection_name": collectionName,
	}
	return c.queryResultToStringsArray("getStickerNamesByCollection", params)
}

func (c *UtopiaClient) GetStickerImage(collectionName, stickerName string) (string, error) {
	params := uMap{
		"collection_name": collectionName,
		"sticker_name":    stickerName,
		"coder":           "BASE64",
	}
	return c.queryResultToString("getImageSticker", params)
}

func (c *UtopiaClient) UCodeEncode(dataHexCode, coder, format string, imageSize int) (string, error) {
	return c.queryResultToString("ucodeEncode", uMap{
		"hex_code":   dataHexCode,
		"size_image": imageSize,
		"coder":      "BASE64",
		"format":     format,
	})
}

func (c *UtopiaClient) SendAuthRequest(pubkey, message string) (bool, error) {
	params := uMap{
		"pk":      pubkey,
		"message": message,
	}
	return c.queryResultToBool("sendAuthorizationRequest", params)
}

func (c *UtopiaClient) AcceptAuthRequest(pubkey, message string) (bool, error) {
	params := uMap{
		"pk":      pubkey,
		"message": message,
	}
	return c.queryResultToBool("acceptAuthorizationRequest", params)
}

func (c *UtopiaClient) RejectAuthRequest(pubkey, message string) (bool, error) {
	params := uMap{
		"pk":      pubkey,
		"message": message,
	}
	return c.queryResultToBool("rejectAuthorizationRequest", params)
}

func (c *UtopiaClient) SendInstantMessage(to string, message string) (int64, error) {
	params := uMap{
		"to":   to,
		"text": message,
	}
	return c.queryResultToInt("sendInstantMessage", params)
}

func (c *UtopiaClient) GetContacts(filter string) ([]structs.ContactData, error) {
	// send request
	params := uMap{}.add("filter", filter)
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

func (c *UtopiaClient) JoinChannel(channelID string, password ...string) (bool, error) {
	params := uMap{
		"ident": channelID,
	}
	if len(password) > 0 {
		params["password"] = password[0]
	}
	return c.queryResultToBool("joinChannel", params)
}

func (c *UtopiaClient) GetChannelContacts(channelID string) ([]structs.ChannelContactData, error) {
	params := uMap{"channelid": channelID}
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

func (c *UtopiaClient) EnableChannelReadOnly(channelID string, readOnly bool) error {
	_, err := c.queryResultToBool("modifyChannel", uMap{
		"channelid": channelID,
		"read_only": readOnly,
	})
	return err
}

func (c *UtopiaClient) RemoveChannelMessage(channelID string, messageID int64) error {
	params := uMap{
		"channelid":  channelID,
		"id_message": messageID,
	}
	_, err := c.queryResultToString("removeChannelMessage", params)
	return err
}

func (c *UtopiaClient) GetChannelMessages(
	channelID string,
	offset int,
	maxMessages int,
) ([]structs.ChannelMessage, error) {
	params := uMap{"channelid": channelID}
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

func (c *UtopiaClient) SendPayment(task structs.SendPaymentTask) (string, error) {
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

func (c *UtopiaClient) GetChannels(task structs.GetChannelsTask) ([]structs.SearchChannelData, error) {
	params := uMap{
		"filter":       task.SearchFilter,
		"channel_type": task.ChannelType,
	}
	if !task.FromDate.IsZero() {
		params["from"] = task.FromDate.Format(defaultTimeLayout)
	}
	if !task.ToDate.IsZero() {
		params["to"] = task.ToDate.Format(defaultTimeLayout)
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
