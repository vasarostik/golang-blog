package transport

import (
	go_blog "github.com/vasarostik/go_blog/pkg/utl/model"
)


// Chat model response
// swagger:response messagesListResp
type swaggChatListResponse struct {
	// in:body
	Body struct {
		Messages go_blog.MessagesList `json:"messages"`
		Page  int          `json:"page"`
	}
}