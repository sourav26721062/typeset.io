package models

//Response - Model description
type Response struct {
	Status string `json:"status,omitempty"`
	Error  string `json:"error,omitempty" `
}
