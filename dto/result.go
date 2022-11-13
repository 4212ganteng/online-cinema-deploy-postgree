package dto

type SuccessResult struct {
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
}

type ErrResult struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}
