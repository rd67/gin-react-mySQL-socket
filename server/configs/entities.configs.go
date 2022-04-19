package configs

type CommonResponseStruct struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
}
