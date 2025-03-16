// internal/global/model/NonStreamResponse.go

package model

type NonStreamResponse struct {
	Success bool          `json:"success"`
	Errors  []interface{} `json:"errors"`
	Data    NonStreamData `json:"data"`
}

type NonStreamData struct {
	ConversationID string    `json:"conversation_id"`
	UserID         int       `json:"user_id"`
	Type           string    `json:"type"`
	Role           string    `json:"role"`
	Content        string    `json:"content"`
	Model          string    `json:"model"`
	ParentID       string    `json:"parent_id"`
	QuestionID     string    `json:"question_id"`
	ContextLength  int       `json:"context_length"`
	Messages       []Message `json:"messages"`
	Usage          Usage     `json:"usage"`
	Status         string    `json:"status"`
	CreatedTime    string    `json:"created_time"`
	UpdatedTime    string    `json:"updated_time"`
	ContentWords   string    `json:"content_words"`
	Duration       int       `json:"duration"`
	FinishReason   string    `json:"finish_reason"`
	ID             string    `json:"id"`
}

type Message struct {
	ID      string      `json:"_id,omitempty"` // 注意omitempty，因为有的message没有_id
	Role    string      `json:"role"`
	Content interface{} `json:"content"` // 可以是字符串或对象数组
}

type Usage struct {
	PromptTokens     int `json:"prompt_tokens"`
	CompletionTokens int `json:"completion_tokens"`
	TotalTokens      int `json:"total_tokens"`
}
