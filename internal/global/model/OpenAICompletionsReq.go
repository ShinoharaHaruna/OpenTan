// internal/global/model/OpenAICompletionsReq.go

package model

type OpenAICompletionsRequest struct {
	Model    string          `json:"model"`
	Messages []OpenAIMessage `json:"messages"`
	Stream   *bool           `json:"stream,omitempty"` // 指针类型，允许为空
}

type OpenAIMessage struct {
	Role    string      `json:"role"`
	Content interface{} `json:"content"`
}
