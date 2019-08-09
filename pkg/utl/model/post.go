package go_blog

type Post struct {
	Base
	ID int `json:"post_id"`
	Title string `json:"title"`
	Content string `json:"content"`
	UserID int `json:"id"`
}
