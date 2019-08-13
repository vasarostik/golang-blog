package go_blog

type PublishPostMessage struct {
	PostID int `json:"post_id"`
	UserID int `json:"user_id"`
	Action string `json:"action"`
	Timestamp string `json:"timestamp"`
}