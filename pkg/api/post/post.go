package post

import (
	"context"
	"encoding/json"
	"github.com/labstack/echo"
	"github.com/vasarostik/go_blog/pkg/grpc/service"
	go_blog "github.com/vasarostik/go_blog/pkg/utl/model"
	"log"
)

// Create creates a new post
func (u *Post) Create(c echo.Context, req go_blog.Post) (*go_blog.Post, error) {

	req.UserID = c.Get("id").(int)
	marsh, err := json.Marshal(req)
	if err != nil {
		log.Println(err)
	}

	Resp,err := u.grpcClient.Create(context.Background(), &service.Post{Data: marsh})
	if err != nil {
		log.Println(err)
	}
	log.Println(Resp)

	return u.udb.Create(u.db, req)
}

// MyList returns list of user`s post
func (u *Post) MyListGRPC(c echo.Context, id int) ([]go_blog.Post, error) {
	var post go_blog.Post

	var postList []go_blog.Post

	posts,err := u.grpcClient.List(context.Background(), &service.Request{UserID: int32(id)})

	if err != nil {
		panic(err)
	}

	for i := range posts.Posts{

		err = json.Unmarshal([]byte(posts.Posts[i]), &post)
		if err != nil {
			panic(err)
		}

		postList = append(postList, post)
	}

	return postList, nil
}

// MyList returns list of user`s post
func (u *Post) MyList(c echo.Context, id int) ([]go_blog.Post, error) {


	return u.udb.MyList(u.db, id)
}

// List returns list of all post
func (u *Post) List(c echo.Context) ([]go_blog.Post, error) {

	return u.udb.List(u.db)
}

// View returns single post
func (u *Post) View(c echo.Context, id int) (*go_blog.Post, error) {

	return u.udb.View(u.db, id)
}

//Delete deletes a post (only user`s post or another posts if this is admin)
func (u *Post) Delete(c echo.Context, id int) error {
	post, err := u.udb.View(u.db, id)
	if err != nil {
		return err
	}

	if err := u.rbac.EnforceUser(c, post.UserID); err != nil {
		return err
	}

	return u.udb.Delete(u.db, post)
}


// Update contains post's information used for updating
type Update struct {
	PostID    int
	Title string
	Content string
}

// Update updates user's post information
func (u *Post) Update(c echo.Context, r *Update) (*go_blog.Post, error) {
	post, err := u.udb.View(u.db, r.PostID)

	if err != nil {
		return nil,err
	}

	if err := u.rbac.EnforceUser(c, post.UserID); err != nil {
		return nil, err
	}

	if err := u.udb.Update(u.db, &go_blog.Post{
		Base:      go_blog.Base{ID: r.PostID},
		UserID: post.UserID,
		Title: r.Title,
		Content:  r.Content,
	}); err != nil {
		return nil, err
	}

	return u.udb.View(u.db, r.PostID)
}

