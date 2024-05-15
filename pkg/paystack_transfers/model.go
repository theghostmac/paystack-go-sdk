package paystack_transfers

// InitiateTransferRequest represents the request body for initiating a transfer.
type InitiateTransferRequest struct {
	Source    string `json:"source"`
	Amount    int    `json:"amount"`
	Recipient string `json:"recipient"`
	Reason    string `json:"reason,omitempty"`
}

// InitiateTransferResponse represents the response from the Paystack API for initiating a transfer.
type InitiateTransferResponse struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
	Data    struct {
		Integration  int    `json:"integration"`
		Domain       string `json:"domain"`
		Amount       int    `json:"amount"`
		Currency     string `json:"currency"`
		Source       string `json:"source"`
		Reason       string `json:"reason"`
		Recipient    int    `json:"recipient"`
		Status       string `json:"status"`
		TransferCode string `json:"transfer_code"`
		ID           int    `json:"id"`
		CreatedAt    string `json:"createdAt"`
		UpdatedAt    string `json:"updatedAt"`
	} `json:"data"`
}

// FinalizeTransferRequest represents the request body for finalizing a transfer.
type FinalizeTransferRequest struct {
	TransferCode string `json:"transfer_code"`
	OTP          string `json:"otp"`
}

// FinalizeTransferResponse represents the response from the Paystack API for finalizing a transfer.
type FinalizeTransferResponse struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
	Data    struct {
		Domain        string `json:"domain"`
		Amount        int    `json:"amount"`
		Currency      string `json:"currency"`
		Reference     string `json:"reference"`
		Source        string `json:"source"`
		SourceDetails string `json:"source_details"`
		Reason        string `json:"reason"`
		Status        string `json:"status"`
		Failures      string `json:"failures"`
		TransferCode  string `json:"transfer_code"`
		TitanCode     string `json:"titan_code"`
		TransferredAt string `json:"transferred_at"`
		ID            int    `json:"id"`
		Integration   int    `json:"integration"`
		Recipient     int    `json:"recipient"`
		CreatedAt     string `json:"createdAt"`
		UpdatedAt     string `json:"updatedAt"`
	} `json:"data"`
}

// Transfer represents a single transfer object in a bulk transfer request.
type Transfer struct {
	Amount    int    `json:"amount"`
	Reference string `json:"reference"`
	Reason    string `json:"reason"`
	Recipient string `json:"recipient"`
}

// InitiateBulkTransferRequest represents the request body for initiating a bulk transfer.
type InitiateBulkTransferRequest struct {
	Source    string     `json:"source"`
	Transfers []Transfer `json:"transfers"`
}

// InitiateBulkTransferResponse represents the response from the Paystack API for initiating a bulk transfer.
type InitiateBulkTransferResponse struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
	Data    []struct {
		Reference    string `json:"reference"`
		Recipient    string `json:"recipient"`
		Amount       int    `json:"amount"`
		TransferCode string `json:"transfer_code"`
		Currency     string `json:"currency"`
		Status       string `json:"status"`
	} `json:"data"`
}
