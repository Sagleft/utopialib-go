package utopiago

import (
	utopia "github.com/Sagleft/utopialib-go/v2/internal/utopia"
)

type Client interface {
	SetProfileStatus(status string, mood string) error
}

func NewUtopiaClient() Client {
	return &utopia.UtopiaClient{}
}
