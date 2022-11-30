package utopia

import "time"

const (
	maxCharactersInPaymentComment = 148
	defaultCurrencyTag            = "CRP"
	defaultPort                   = 20000
	defaultWsPort                 = 25000
	defaultHost                   = "127.0.0.1"
	defaultTimeLayout             = time.RFC3339
	defaultReqRateLimitTimeout    = time.Second * 2

	reqDefault                     = "default"
	reqGetProfileStatus            = "getProfileStatus"
	reqGetSystemInfo               = "getSystemInfo"
	reqSetProfileStatus            = "setProfileStatus"
	reqGetOwnContact               = "getOwnContact"
	reqUseVoucher                  = "useVoucher"
	reqGetFinanceSystemInformation = "getFinanceSystemInformation"
	reqGetFinanceHistory           = "getFinanceHistory"
)
