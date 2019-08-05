package transport

import (
	"net/http"
	"strconv"

	"github.com/vasarostik/go_blog/pkg/api/password"

	"github.com/vasarostik/go_blog/pkg/utl/model"

	"github.com/labstack/echo"
)

// HTTP represents password http transport service
type HTTP struct {
	svc password.Service
}

// NewHTTP creates new password http service
func NewHTTP(svc password.Service, er *echo.Group) {
	h := HTTP{svc}
	pr := er.Group("/password")

	// swagger:operation PATCH /v1/password/{id} password pwChange
	// ---
	// summary: Changes user's password.
	// description: If user's old passowrd is correct, it will be replaced with new password.
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
	//     "$ref": "#/definitions/pwChange"
	// responses:
	//   "200":
	//     "$ref": "#/responses/ok"
	//   "400":
	//     "$ref": "#/responses/errMsg"
	//   "401":
	//     "$ref": "#/responses/err"
	//   "403":
	//     "$ref": "#/responses/err"
	//   "500":
	//     "$ref": "#/responses/err"
	pr.PATCH("/:id", h.change)
}

// Password change request
// swagger:model pwChange
type changeReq struct {
	ID                 int    `json:"-"`
	OldPassword        string `json:"old_password" validate:"required,min=8"`
	NewPassword        string `json:"new_password" validate:"required,min=8"`
	NewPasswordConfirm string `json:"new_password_confirm" validate:"required"`
}

func (h *HTTP) change(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return go_blog.ErrBadRequest
	}

	p := new(changeReq)
	if err := c.Bind(p); err != nil {
		return err
	}

	if p.NewPassword != p.NewPasswordConfirm {
		return go_blog.ErrPasswordsNotMaching
	}

	if err := h.svc.Change(c, id, p.OldPassword, p.NewPassword); err != nil {
		return err
	}

	return c.NoContent(http.StatusOK)
}
