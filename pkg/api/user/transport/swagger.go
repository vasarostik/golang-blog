package transport

import go_blog "github.com/vasarostik/go_blog/pkg/utl/model"

// User model response
// swagger:response userResp
type swaggUserResponse struct {
	// in:body
	Body struct {
		*go_blog.User
	}
}

// Users model response
// swagger:response userListResp
type swaggUserListResponse struct {
	// in:body
	Body struct {
		Users []go_blog.User `json:"users"`
		Page  int          `json:"page"`
	}
}