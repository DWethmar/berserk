package sagemcomfast5359

type Errors struct {
	Error       int    `json:"error"`
	Description string `json:"description"`
	Info        string `json:"info"`
}

type ErrResponse struct {
	Errors []Errors `json:"errors"`
}
