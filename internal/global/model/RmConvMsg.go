// internal/global/model/RmConvMsg.go

package model

import "time"

type Conversation struct {
	ConversationID string     `json:"conversation_id"`
	UpdateMask     []string   `json:"update_mask"`
	Status         StatusType `json:"status"`
}

type StatusType string

const (
	SoftDeleted StatusType = "SOFT_DELETED"
	// 其他的以后再加
)

func (typ StatusType) String() string {
	switch typ {
	case SoftDeleted:
		return "SOFT_DELETED"
	default:
		return "UNKNOWN"
	}
}

type RmConvResponse struct {
	Success bool       `json:"success"`
	Errors  []string   `json:"errors"`
	Data    RmConvData `json:"data"`
}

type RmConvData struct {
	UserID              int       `json:"user_id"`
	Type                string    `json:"type"`
	Title               string    `json:"title"`
	TagIDs              []string  `json:"tag_ids"`
	SystemPrompt        string    `json:"system_prompt"`
	Model               string    `json:"model"`
	ContextLength       int       `json:"context_length"`
	Temperature         float64   `json:"temperature"`
	MaxTokens           int       `json:"max_tokens"`
	IsPinned            bool      `json:"is_pinned"`
	Status              string    `json:"status"`
	CreatedTime         time.Time `json:"created_time"`
	UpdatedTime         time.Time `json:"updated_time"`
	LastUserMessageTime time.Time `json:"last_user_message_time"`
	DeletedTime         time.Time `json:"deleted_time"`
	LastMessage         string    `json:"last_message"`
	ID                  string    `json:"id"`
}
