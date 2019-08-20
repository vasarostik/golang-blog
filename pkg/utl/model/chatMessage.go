package go_blog

type ChatMessage struct {
	Username string `json:"username"`
	Message  string `json:"message"`
}

type MessagesList struct {
	Messages []string `json:"messages,omitempty"`
}

