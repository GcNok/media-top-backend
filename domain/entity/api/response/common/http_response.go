package common

type (
	HttpResponse struct {
		ResultCode int    `json:"resultCode"`
		Message    string `json:"message"`
	}
)
