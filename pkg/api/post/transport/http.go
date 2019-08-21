package transport

import (
	"github.com/labstack/echo"
	"github.com/vasarostik/go_blog/pkg/api/post"
	go_blog "github.com/vasarostik/go_blog/pkg/utl/model"
	"net/http"
	"strconv"
)

// HTTP represents user http service
type HTTP struct {
	svc post.Service
}

// NewHTTP creates new user http service
func NewHTTP(svc post.Service, er *echo.Group) {
	h := HTTP{svc}

	ur := er.Group("/post")

	// swagger:route POST /v1/post/create posts createPost
	// Creates new post.
	// responses:
	//  200: postResp
	//  400: errMsg
	//  401: err
	//  403: errMsg
	//  500: err
	ur.POST("/create", h.create)

	// swagger:operation GET /v1/posts/my posts MyPosts
	// ---
	// summary: Returns list of users`s posts.
	// description: Returns list of posts.
	// parameters:
	// - name: limit
	//   in: query
	//   description: number of results
	//   type: int
	//   required: false
	// - name: page
	//   in: query
	//   description: page number
	//   type: int
	//   required: false
	// responses:
	//   "200":
	//     "$ref": "#/responses/postListResp"
	//   "400":
	//     "$ref": "#/responses/errMsg"
	//   "401":
	//     "$ref": "#/responses/err"
	//   "403":
	//     "$ref": "#/responses/err"
	//   "500":
	//     "$ref": "#/responses/err"
	ur.GET("s/my", h.myList)

	// swagger:operation GET /v1/posts/my/grpc posts myPostsGRPC
	// ---
	// summary: Returns list of user`s posts with GRPC.
	// description: Returns list of users.
	// parameters:
	// - name: limit
	//   in: query
	//   description: number of results
	//   type: int
	//   required: false
	// - name: page
	//   in: query
	//   description: page number
	//   type: int
	//   required: false
	// responses:
	//   "200":
	//     "$ref": "#/responses/postListResp"
	//   "400":
	//     "$ref": "#/responses/errMsg"
	//   "401":
	//     "$ref": "#/responses/err"
	//   "403":
	//     "$ref": "#/responses/err"
	//   "500":
	//     "$ref": "#/responses/err"
	ur.GET("s/my/grpc", h.myListGRPC)

	// swagger:operation GET /v1/posts posts allPosts
	// ---
	// summary: Returns list of posts.
	// description: Returns list of post.
	// parameters:
	// - name: limit
	//   in: query
	//   description: number of results
	//   type: int
	//   required: false
	// - name: page
	//   in: query
	//   description: page number
	//   type: int
	//   required: false
	// responses:
	//   "200":
	//     "$ref": "#/responses/postListResp"
	//   "400":
	//     "$ref": "#/responses/errMsg"
	//   "401":
	//     "$ref": "#/responses/err"
	//   "403":
	//     "$ref": "#/responses/err"
	//   "500":
	//     "$ref": "#/responses/err"
	ur.GET("s", h.list)

	// swagger:operation GET /v1/posts/{id} posts singlePost
	// ---
	// summary: Returns a single post.
	// description: Returns a single post by its ID.
	// parameters:
	// - name: id
	//   in: path
	//   description: id of post
	//   type: int
	//   required: true
	// responses:
	//   "200":
	//     "$ref": "#/responses/postResp"
	//   "400":
	//     "$ref": "#/responses/err"
	//   "401":
	//     "$ref": "#/responses/err"
	//   "403":
	//     "$ref": "#/responses/err"
	//   "404":
	//     "$ref": "#/responses/err"
	//   "500":
	//     "$ref": "#/responses/err"
	ur.GET("/:id", h.view)

	// swagger:operation PATCH /v1/post/{id} posts singleUpdatePost
	// ---
	// summary: Updates post's information
	// parameters:
	// - name: id
	//   in: path
	//   description: id of post
	//   type: int
	//   required: true
	// - name: title
	//   in: body
	//   description: Title of post
	//   required: true
	// - name: content
	//   in: body
	//   description: Conent of post
	//   required: true
	// responses:
	//   "200":
	//     "$ref": "#/responses/postResp"
	//   "400":
	//     "$ref": "#/responses/errMsg"
	//   "401":
	//     "$ref": "#/responses/err"
	//   "403":
	//     "$ref": "#/responses/err"
	//   "500":
	//     "$ref": "#/responses/err"
	ur.PATCH("/:id", h.update)
	//
	// swagger:operation DELETE /v1/post/{id} posts deletePost
	// ---
	// summary: Deletes a post
	// description: Deletes a post with requested ID.
	// parameters:
	// - name: id
	//   in: path
	//   description: id of post
	//   type: int
	//   required: true
	// responses:
	//   "200":
	//     "$ref": "#/responses/ok"
	//   "400":
	//     "$ref": "#/responses/err"
	//   "401":
	//     "$ref": "#/responses/err"
	//   "403":
	//     "$ref": "#/responses/err"
	//   "500":
	//     "$ref": "#/responses/err"


	ur.DELETE("/:id", h.delete)
}

const (
	topicPublishPost = "posts:publish"
)

// swagger:model userCreate
type createReq struct {
	Title string `json:"title"`
	Content string `json:"content"`
	UserID int `json:"id"`
}

func (h *HTTP) create(c echo.Context) error {
	r := new(createReq)

	if err := c.Bind(r); err != nil {

		return err
	}
	id := c.Get("id").(int)

	postCreated, err := h.svc.Create(c, go_blog.Post{
		Title:   r.Title,
		Content:   r.Content,
		UserID: id,
	})

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, postCreated)
}

type listResponse struct {
	Posts []go_blog.Post `json:"posts"`
}



func (h *HTTP) myList(c echo.Context) error {

	id := c.Get("id").(int)

	result, err := h.svc.MyList(c, id)

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, listResponse{result})
}


func (h *HTTP) myListGRPC(c echo.Context) error {

	id := c.Get("id").(int)

	result, err := h.svc.MyListGRPC(c, id)

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, listResponse{result})
}


func (h *HTTP) list(c echo.Context) error {

	result, err := h.svc.List(c)

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, listResponse{result})
}

func (h *HTTP) view(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return go_blog.ErrBadRequest
	}

	result, err := h.svc.View(c, id)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, result)
}

// swagger:model userUpdate
type updateReq struct {
	ID        int    `json:"-"`
	FTitle string `json:"title"`
	Content string `json:"content"`
}

func (h *HTTP) update(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return go_blog.ErrBadRequest
	}

	req := new(updateReq)
	if err := c.Bind(req); err != nil {
		return err
	}

	post, err := h.svc.Update(c, &post.Update{
		PostID:        id,
		Title: req.FTitle,
		Content:  req.Content,
	})

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, post)
}


func (h *HTTP) delete(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return go_blog.ErrBadRequest
	}

	if err := h.svc.Delete(c, id); err != nil {
		return err
	}

	return c.NoContent(http.StatusOK)
}
