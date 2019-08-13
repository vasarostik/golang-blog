package post

import (
	"context"
	"encoding/json"
	"github.com/labstack/echo"
	"github.com/pkg/errors"
	"github.com/vasarostik/go_blog/pkg/grpc/service"
	go_blog "github.com/vasarostik/go_blog/pkg/utl/model"
	"log"
	"time"
)


const (
	topicPublishPost = "posts:publish"
)

// Create creates a new post
func (u *Post) Create(c echo.Context, req go_blog.Post) (*go_blog.Post, error) {

	req.UserID = c.Get("id").(int)
	marsh, err := json.Marshal(req)

	checkErr(err)

	Resp,err := u.grpcClient.Create(context.Background(), &service.Post{Data: marsh})

	if err != nil {
		log.Println(err)
	}

	log.Println(Resp)

	res, err := u.udb.Create(u.db, req)

	checkErr(err)

	message := &go_blog.PublishPostMessage{
		PostID: res.ID,
		UserID: res.UserID,
		Timestamp: time.Now().Format(time.RFC850),
		Action: "created",
	}

	if err := u.PublishMessage(topicPublishPost, *message); err != nil {
		panic(err)
	}

	return res, nil
}


// MyList returns list of user`s post
func (u *Post) MyListGRPC(c echo.Context, id int) ([]go_blog.Post, error) {
	var postStruct go_blog.Post

	var postList []go_blog.Post

	posts,err := u.grpcClient.List(context.Background(), &service.Request{UserID: int32(id)})

	if err != nil {
		panic(err)
	}

	for i := range posts.Posts{

		err = json.Unmarshal([]byte(posts.Posts[i]), &postStruct)
		if err != nil {
			panic(err)
		}

		postList = append(postList, postStruct)
	}

	return postList, nil
}

func (u *Post) PublishMessage(topic string, msg go_blog.PublishPostMessage) error {
	bs, err := json.Marshal(msg)
	if err != nil {
		return errors.Wrap(err, "failed to marshal proto message")
	}

	if err := u.natsClient.Publish(topicPublishPost, bs); err != nil {
		return errors.Wrap(err, "failed to publish message")
	}

	if err := u.natsClient.Flush(); err != nil {
		return errors.Wrap(err, "failed to flush message")
	}

	if err := u.natsClient.LastError(); err != nil {
		return errors.Wrap(err, "received error after publishing")
	}

	return nil
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
	postStruct, err := u.udb.View(u.db, id)
	if err != nil {
		return err
	}

	if err := u.rbac.EnforceUser(c, postStruct.UserID); err != nil {
		return err
	}

	err = u.udb.Delete(u.db, postStruct)

	checkErr(err)

	message := &go_blog.PublishPostMessage{
		PostID: id,
		UserID: postStruct.UserID,
		Timestamp: time.Now().Format(time.RFC850),
		Action: "deleted",
	}

	if err := u.PublishMessage(topicPublishPost, *message); err != nil {
		panic(err)
	}

	return nil
}


// Update contains post's information used for updating
type Update struct {
	PostID    int
	Title string
	Content string
}

// Update updates user's post information
func (u *Post) Update(c echo.Context, r *Update) (*go_blog.Post, error) {
	postStruct, err := u.udb.View(u.db, r.PostID)

	if err != nil {
		return nil,err
	}

	if err := u.rbac.EnforceUser(c, postStruct.UserID); err != nil {
		return nil, err
	}

	if err := u.udb.Update(u.db, &go_blog.Post{
		Base:   go_blog.Base{ID: r.PostID},
		UserID: postStruct.UserID,
		Title: r.Title,
		Content:  r.Content,
	}); err != nil {
		return nil, err
	}

	res, err := u.udb.View(u.db, r.PostID)

	checkErr(err)

	message := &go_blog.PublishPostMessage{
		PostID: res.ID,
		UserID: res.UserID,
		Timestamp: time.Now().Format(time.RFC850),
		Action: "updated",
	}

	if err := u.PublishMessage(topicPublishPost, *message); err != nil {
		panic(err)
	}

	return res,nil
}


func checkErr(err error) {
	if err != nil {
		panic(err.Error())
	}
}
