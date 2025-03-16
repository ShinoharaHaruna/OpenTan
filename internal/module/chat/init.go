// internal/module/chat/init.go

package chat

type ChatModule struct{}

func (m *ChatModule) GetName() string {
	return "chat"
}

func (m *ChatModule) Init() {}
