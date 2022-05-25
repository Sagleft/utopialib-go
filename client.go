package utopiago

import (
	"errors"
)

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
	if result == false {
		return errors.New("failed to set profile status")
	}
	return nil
}

// GetOwnContact asks for full details of the current account
func (c *UtopiaClient) GetOwnContact() (map[string]interface{}, error) {
	return c.apiQuery("getOwnContact", nil)
}

// CheckClientConnection - checks if there are any errors when contacting the client
func (c *UtopiaClient) CheckClientConnection() bool {
	_, err := c.GetSystemInfo()
	if err != nil {
		return false
	}
	return true
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
func (c *UtopiaClient) SetWebSocketState(enabled bool, port int) error {
	var enabledStr string
	if enabled {
		enabledStr = "1"
	}
	params := map[string]interface{}{
		"enabled": enabledStr,
		"port":    port,
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