// internal/global/model/NewConvBody.go

package model

type NewConvBody struct {
	Content      string       `json:"content"`
	Stream       bool         `json:"stream"`
	Conversation ConvMetadata `json:"conversation"`
}

type ConvMetadata struct {
	Title string `json:"title"`
	Model string `json:"model"`
}
