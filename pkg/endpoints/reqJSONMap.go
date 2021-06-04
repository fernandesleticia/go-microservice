package endpoints

type GetRequest struct {
	Filters []internal.Filter `json:"filters,omitempty"`
}

type GetResponse struct {
	Transactions []internal.Transaction `json:"transactions"`
	Err 		 string 				`json:"err,omitempty"`
}

type StatusRequest struct {
	TransactionID string `json:"status"`
}

type StatusResponse struct {
	Status internal.Status `json:"status"`
	Err string 			   `json:"err,omitempty"`
}

type TransactionRequest struct {
	Transaction *internal.Transaction `json:"transaction"`
}

type TransactionResponse struct{
	TransactionID string `json:"transactionID"`
}

type ServiceStatusRequest struct{}

type ServiceStatusResponse struct {
	Code int   `json:"status"`
	Err string `json:"err,omitempty"`
}