package utopiago

import "github.com/Sagleft/utopialib-go/v2/internal/utopia"

type Client interface {
	SetProfileStatus(status string, mood string) error
}

type Config = utopia.Config

func NewUtopiaClient(c Config) Client {
	return utopia.NewUtopiaClient(c)
}
