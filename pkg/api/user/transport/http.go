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

	e.POST("/users", h.create)

	ur := er.Group("/users")

	ur.GET("", h.list)

	ur.GET("/:id", h.view)

	ur.PATCH("/:id", h.update)

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

