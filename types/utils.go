package types

// 타입들 선언

type ApiResponse struct {
	Result      int         `json:"result"`
	Description string      `json:"description"`
	ErrCode     interface{} `json:"errCode"`
}

func NewApiResponse(description string, result int, errCode any) *ApiResponse {
	return &ApiResponse{
		Result:      result,
		Description: description,
		ErrCode:     errCode,
	}
}
