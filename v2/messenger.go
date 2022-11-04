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
}

type Config = utopia.Config

func NewUtopiaClient(c Config) Client {
	return utopia.NewUtopiaClient(c)
}
