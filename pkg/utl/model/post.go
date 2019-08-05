package go_blog

type Post struct {
	Base
	Title string `json:"title"`
	Content string `json:"content"`
	UserID int `json:"id"`
}
