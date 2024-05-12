package types

// 타입들 선언

type ApiResponse struct {
	Result      int    `json:"result"`
	Description string `json:"description"`
}

func NewApiResponse(description string, result int) *ApiResponse {
	return &ApiResponse{
		Result:      result,
		Description: description,
	}
}
