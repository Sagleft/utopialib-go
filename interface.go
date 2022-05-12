package utopiago

// UtopiaClientInterface contains an enumeration of methods
type UtopiaClientInterface interface {
	apiQuery(methodName string) map[string]interface{}
	// profile
	GetProfileStatus() map[string]interface{}
	GetSystemInfo() map[string]interface{}
	GetOwnContact() map[string]interface{}
	// crypton
	GetBalance() (float64, error)
	UseVoucher(voucherCode string) error
	GetFinanceHistory() map[string]interface{}
	CheckClientConnection() bool
	CreateVoucher(amount float64) error
	// channels
	SendChannelMessage(channelID, message string) (string, error)
	SendChannelPicture(channelID, base64Image, comment, filenameForImage string) (string, error)
}
