package structs

import "time"

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
