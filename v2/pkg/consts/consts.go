package consts

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

const (
	SortChannelsByCreated SortChannelsBy = iota + 1
	SortChannelsByIsPrivate
	SortChannelsByModified
	SortChannelsByName
	SortChannelsByDescription
)

type SortChannelsBy int

type ChannelType int

const (
	StatusCodeOnline       = 4096
	StatusCodeOffline      = 65536
	StatusCodeAway         = 4097
	StatusCodeDoNotDisturb = 4099
	StatusCodeInvisible    = 32768
)
