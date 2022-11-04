package structs

type FinanceInfo struct {
	CRP  CryptonFinanceInfo `json:"CRP"`
	UUSD UUSDFinanceInfo    `json:"USD"`

	ProofOfStake      bool `json:"PoS"`
	EnableToUseMining bool `json:"enableToUseMining"`
	MinintPeriod      int  `json:"miningPeriod"`
	SettingsVersion   int  `json:"settingsVersion"`
}

type CryptonFinanceInfo struct {
	CardCreatePriceDefault        float64 `json:"cardCreatePrice"`
	CardCreatePrice1Symbol        float64 `json:"cardCreatePrice10"`
	CardCreatePrice2Symbols       float64 `json:"cardCreatePrice100"`
	CardCreatePrice3Symbols       float64 `json:"cardCreatePrice1000"`
	CardCreatePrice4Symbols       float64 `json:"cardCreatePrice10000"`
	CardsCreationEnabled          bool    `json:"cardsCreationEnabled"`
	CardsMaxActive                uint    `json:"cardsMaxActive"`
	CardsMaxPerDay                uint    `json:"cardsMaxPerDay"`
	InvestorMinAmount             int64   `json:"investorMinAmount"`
	InvoicesDefaultTtl            uint    `json:"invoicesDefaultTtl"`
	InvoicesEnabled               bool    `json:"invoicesEnabled"`
	InvoicesMaxTotal              uint    `json:"invoicesMaxTotal"`
	InvoicesMaxTotalFromMerchant  uint    `json:"invoicesMaxTotalFromMerchant"`
	InvoicesMinAmount             float64 `json:"invoicesMinAmount"`
	TransferCardFee               float64 `json:"transferCardFee"`
	TransferCheckFee              bool    `json:"transferCheckFee"`
	TransferExternalFee           float64 `json:"transferExternalFee"`
	TransferInternalFee           float64 `json:"transferInternalFee"`
	TransfersEnabled              bool    `json:"transfersEnabled"`
	UnsDefaultTtl                 uint    `json:"unsDefaultTtl"`
	UnsDeleteNameFee              float64 `json:"unsDeleteNameFee"`
	UnsModifyNameFee              float64 `json:"unsModifyNameFee"`
	UnsName1SymbolRegistrationFee float64 `json:"unsName1RegistrationFee"`
	UnsName2SymbolRegistrationFee float64 `json:"unsName2RegistrationFee"`
	UnsName3SymbolRegistrationFee float64 `json:"unsName3RegistrationFee"`
	UnsName4SymbolRegistrationFee float64 `json:"unsName4RegistrationFee"`
	UnsProxyEnabled               bool    `json:"unsProxyEnabled"`
	UnsTransferFee                float64 `json:"unsTransferFee"`
	VouchersCreateEnabled         bool    `json:"vouchersCreateEnabled"`
	VouchersMaxActive             uint    `json:"vouchersMaxActive"`
	VouchersMaxPerBatch           uint    `json:"vouchersMaxPerBatch"`
	VouchersMinAmount             float64 `json:"vouchersMinAmount"`
	VouchersMinPerBatch           uint    `json:"vouchersMinPerBatch"`
	VouchersUseEnabled            bool    `json:"vouchersUseEnabled"`
}

type UUSDFinanceInfo struct {
	TransferExternalFee   float64 `json:"transferExternalFee"`
	TransferInternalFee   float64 `json:"transferInternalFee"`
	TransfersEnabled      bool    `json:"transfersEnabled"`
	VouchersCreateEnabled bool    `json:"vouchersCreateEnabled"`
	VouchersMaxActive     uint    `json:"vouchersMaxActive"`
	VouchersMaxPerBatch   uint    `json:"vouchersMaxPerBatch"`
	VouchersMinAmount     float64 `json:"vouchersMinAmount"`
	VouchersMinPerBatch   uint    `json:"vouchersMinPerBatch"`
	VouchersUseEnabled    bool    `json:"vouchersUseEnabled"`
}

type FinanceHistoryData struct{}
