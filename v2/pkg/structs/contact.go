package structs

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
