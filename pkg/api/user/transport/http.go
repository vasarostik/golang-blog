package transport

import (
	"net/http"
	"strconv"

	"github.com/vasarostik/go_blog/pkg/api/user"

	go_blog "github.com/vasarostik/go_blog/pkg/utl/model"

	"github.com/labstack/echo"
)

// HTTP represents user http service
type HTTP struct {
	svc user.Service
}

// NewHTTP creates new user http service
func NewHTTP(svc user.Service, er *echo.Group, e *echo.Echo) {
	h := HTTP{svc}

	e.POST("/users", h.create) // user creates account

	ur := er.Group("/users")

	// swagger:route POST /v1/users users userCreate
	// Creates new user account.
	// responses:
	//  200: userResp
	//  400: errMsg
	//  401: err
	//  403: errMsg
	//  500: err
	ur.POST("", h.create) // admin creates user`s account(missed some RBAC logic)

	// swagger:operation GET /v1/users users listUsers
	// ---
	// summary: Returns list of users.
	// description: Returns list of users. Depending on the user role requesting it, it may return all users for SuperAdmin/Admin users, all company/location users for Company/Location admins, and an error for non-admin users.
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
	//     "$ref": "#/responses/userListResp"
	//   "400":
	//     "$ref": "#/responses/errMsg"
	//   "401":
	//     "$ref": "#/responses/err"
	//   "403":
	//     "$ref": "#/responses/err"
	//   "500":
	//     "$ref": "#/responses/err"
	ur.GET("", h.list)

	// swagger:operation GET /v1/users/{id} users getUser
	// ---
	// summary: Returns a single user.
	// description: Returns a single user by its ID.
	// parameters:
	// - name: id
	//   in: path
	//   description: id of user
	//   type: int
	//   required: true
	// responses:
	//   "200":
	//     "$ref": "#/responses/userResp"
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

	// swagger:operation PATCH /v1/users/{id} users userUpdate
	// ---
	// summary: Updates user's contact information
	// description: Updates user's contact information -> first name, last name, mobile, phone, address.
	// parameters:
	// - name: id
	//   in: path
	//   description: id of user
	//   type: int
	//   required: true
	// - name: request
	//   in: body
	//   description: Request body
	//   required: true
	//   schema:
	//     "$ref": "#/definitions/userUpdate"
	// responses:
	//   "200":
	//     "$ref": "#/responses/userResp"
	//   "400":
	//     "$ref": "#/responses/errMsg"
	//   "401":
	//     "$ref": "#/responses/err"
	//   "403":
	//     "$ref": "#/responses/err"
	//   "500":
	//     "$ref": "#/responses/err"
	ur.PATCH("/:id", h.update)

	// swagger:operation DELETE /v1/users/{id} users userDelete
	// ---
	// summary: Deletes a user
	// description: Deletes a user with requested ID.
	// parameters:
	// - name: id
	//   in: path
	//   description: id of user
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

// User create request
// swagger:model userCreate
type createReq struct {
	FirstName       string `json:"first_name" validate:"required"`
	LastName        string `json:"last_name" validate:"required"`
	Username        string `json:"username" validate:"required,min=3,alphanum"`
	Password        string `json:"password" validate:"required,min=8"`
	PasswordConfirm string `json:"password_confirm" validate:"required"`


	RoleID     go_blog.AccessRole `json:"role_id" validate:"required"`
}

func (h *HTTP) create(c echo.Context) error {
	r := new(createReq)

	if err := c.Bind(r); err != nil {

		return err
	}
	if r.Password != r.PasswordConfirm {
		return go_blog.ErrPasswordsNotMaching
	}

	if r.RoleID < go_blog.SuperAdminRole || r.RoleID > go_blog.UserRole {
		return go_blog.ErrBadRequest
	}

	usr, err := h.svc.Create(c, go_blog.User{
		Username:   r.Username,
		Password:   r.Password,
		FirstName:  r.FirstName,
		LastName:   r.LastName,
		RoleID:     r.RoleID,
	})

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, usr)
}

type listResponse struct {
	Users []go_blog.User `json:"users"`
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

// User update request
// swagger:model userUpdate
type updateReq struct {
	ID        int    `json:"-"`
	FirstName string `json:"first_name,omitempty" validate:"omitempty,min=2"`
	LastName  string `json:"last_name,omitempty" validate:"omitempty,min=2"`

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

	usr, err := h.svc.Update(c, &user.Update{
		ID:        id,
		FirstName: req.FirstName,
		LastName:  req.LastName,
	})

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, usr)
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

