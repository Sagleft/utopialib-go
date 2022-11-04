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
