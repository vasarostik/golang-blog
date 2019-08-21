package transport

import (
	go_blog "github.com/vasarostik/go_blog/pkg/utl/model"
)

// Post response
// swagger:response postResp
type swaggPostResp struct {
	// in:body
	Body struct {
		*go_blog.Post
	}
}

// Post model response
// swagger:response postListResp
type swaggUserListResponse struct {
	// in:body
	Body struct {
		Posts []go_blog.Post `json:"posts"`
		Page  int          `json:"page"`
	}
}

