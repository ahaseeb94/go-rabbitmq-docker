package Core

type UnformedResponse struct {
	StatusCode string                 `json:"status_code"`
	HTTPCode   int                    `json:"http_code"`
	Content    map[string]interface{} `json:"content"`
}

func UnformedResponseDraft(statusCode string, httpCode int, content map[string]interface{}) *UnformedResponse {
	return &UnformedResponse{
		StatusCode: statusCode,
		HTTPCode:   httpCode,
		Content:    content,
	}
}
