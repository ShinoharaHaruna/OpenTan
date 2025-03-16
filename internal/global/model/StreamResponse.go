// internal/global/model/StreamResponse.go

package model

type StreamResponse struct {
	ID             string   `json:"id"`
	ParentID       string   `json:"parent_id"`
	ConversationID string   `json:"conversation_id"`
	QuestionID     string   `json:"question_id"`
	Model          string   `json:"model"`
	Choices        []Choice `json:"choices"`
	CreatedTime    string   `json:"created_time"`
}

type Choice struct {
	Delta        Delta       `json:"delta"`
	Logprobs     interface{} `json:"logprobs"`      // 可以是 null，所以使用 interface{}
	FinishReason interface{} `json:"finish_reason"` // 可以是 null，所以使用 interface{}
	Index        int         `json:"index"`
}

type Delta struct {
	Content string `json:"content"`
	Role    string `json:"role"`
}
