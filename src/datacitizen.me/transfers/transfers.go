package transfers

import (
	"encoding/json"
)

// DataTransferRequest A request to make a data transfer
type DataTransferRequest struct {
	Producer          string   `json:"producer"`
	Consumer          string   `json:"consumer"`
	DataCategory      string   `json:"data-category"`
	CitizenIdentities []string `json:"citizen-identities"`
}

// DataTransferStatusResponse Base Response, providing a simple status
type DataTransferStatusResponse struct {
	Status string `json:"status"`
}

// ToJSON Gets a Json representation of the response.
func (dtsr DataTransferStatusResponse) ToJSON() ([]byte, error) {
	return json.Marshall(dtsr)
}

// DataTransferConditionalResponse Provides status and a list of data records to be scrubbed.
type DataTransferConditionalResponse struct {
	DataTransferStatusResponse
	ScrubList []string `json:"scrub-list"`
}

// DataTransferMessageResponse Provides a status and a message (used for errors)
type DataTransferMessageResponse struct {
	DataTransferStatusResponse
	Message string `json:"message"`
}

// DataTransferResponse Defines the contract of a response to a DataTransferRequest
type DataTransferResponse interface {
	ToJSON() ([]byte, error)
}

// ===========
// = Methods =
// ===========

// GetResponse Gets a response to a DataTransferRequest
func (dtr DataTransferRequest) GetResponse() DataTransferResponse {
	// TODO
}
