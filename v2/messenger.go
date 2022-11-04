package utopiago

import (
	"github.com/Sagleft/utopialib-go/v2/internal/utopia"
	"github.com/Sagleft/utopialib-go/v2/pkg/structs"
)

type Client interface {
	// GetProfileStatus gets data about the status of the current account
	GetProfileStatus() (structs.ProfileStatus, error)

	// SetProfileStatus updates data about the status of the current account
	SetProfileStatus(status string, mood string) error

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
}

type Config = utopia.Config

func NewUtopiaClient(c Config) Client {
	return utopia.NewUtopiaClient(c)
}
