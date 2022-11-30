package structs

import "github.com/Sagleft/utopialib-go/v2/pkg/consts"

type ProfileStatus struct {
	Mood       string `json:"mood"`
	Status     string `json:"status"`
	StatusCode int    `json:"status_code"`
}

// IsOnline - is contact online?
func (d ProfileStatus) IsOnline() bool {
	return d.StatusCode == consts.StatusCodeOnline
}

// IsOffline - is contact offline?
func (d ProfileStatus) IsOffline() bool {
	return d.StatusCode == consts.StatusCodeOffline
}

// IsAway - is contact away?
func (d ProfileStatus) IsAway() bool {
	return d.StatusCode == consts.StatusCodeAway
}

// IsDoNotDisturb - is contact marked to do not disturb mode?
func (d ProfileStatus) IsDoNotDisturb() bool {
	return d.StatusCode == consts.StatusCodeDoNotDisturb
}

// IsInvisible - is contact invisible?
func (d ProfileStatus) IsInvisible() bool {
	return d.StatusCode == consts.StatusCodeInvisible
}

// IsOnline - is contact online?
func (d *ContactData) IsOnline() bool {
	return d.Status == consts.StatusCodeOnline
}

// IsOffline - is contact offline?
func (d *ContactData) IsOffline() bool {
	return d.Status == consts.StatusCodeOffline
}

// IsAway - is contact away?
func (d *ContactData) IsAway() bool {
	return d.Status == consts.StatusCodeAway
}

// IsDoNotDisturb - is contact marked to do not disturb mode?
func (d *ContactData) IsDoNotDisturb() bool {
	return d.Status == consts.StatusCodeDoNotDisturb
}

// IsInvisible - is contact invisible?
func (d *ContactData) IsInvisible() bool {
	return d.Status == consts.StatusCodeInvisible
}
