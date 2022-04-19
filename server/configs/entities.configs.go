package configs

type ICommonResponse struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
}
