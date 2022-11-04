package structs

import (
	"time"

	"github.com/Sagleft/utopialib-go/v2/pkg/consts"
)

type GetFinanceHistoryTask struct {
	// optional
	Currency          string    `json:"currency"`
	Filters           string    `json:"filters"`
	ReferenceNumber   string    `json:"referenceNumber"`
	FromDate          time.Time `json:"fromDate"`
	ToDate            time.Time `json:"toDate"`
	BatchID           int       `json:"batchId"`
	FromAmount        float64   `json:"fromAmount"`
	ToAmount          float64   `json:"toAmount"`
	SourcePubkey      string    `json:"sourcePk"`
	DestinationPubkey string    `json:"destinationPk"`
	QueryOffset       uint      `json:"offset"`
	QueryLimitRows    uint      `json:"limitRows"`
}

type SendPaymentTask struct {
	// required
	To     string  `json:"to"`     // pubkey, nickname or card ID
	Amount float64 `json:"amount"` // more than zero, no more than 9 decimal places

	// optional
	CurrencyTag string `json:"currency"`   // example: "CRP", "UUSD". by default: "CRP"
	FromCardID  string `json:"fromCardID"` // specify here your card ID
	Comment     string `json:"comment"`
}

type GetChannelsTask struct {
	// optional
	SearchFilter string             // part of channel name or channel ID, etc
	ChannelType  consts.ChannelType // by default: 0 - registered
	FromDate     time.Time
	ToDate       time.Time
	SortBy       consts.SortChannelsBy
}
