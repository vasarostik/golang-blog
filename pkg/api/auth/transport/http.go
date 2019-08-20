package transport

import (
	"net/http"

	"github.com/vasarostik/go_blog/pkg/api/auth"

	"github.com/labstack/echo"
)

// HTTP represents auth http service
type HTTP struct {
	svc auth.Service
}

// NewHTTP creates new auth http service
func NewHTTP(svc auth.Service, e *echo.Echo, mw echo.MiddlewareFunc) {
	h := HTTP{svc}

	e.POST("/login", h.login)

	e.GET("/refresh/:token", h.refresh)

	e.GET("/me", h.me, mw)
}

type credentials struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

func (h *HTTP) login(c echo.Context) error {
	cred := new(credentials)
	if err := c.Bind(cred); err != nil {
		return err
	}
	r, err := h.svc.Authenticate(c, cred.Username, cred.Password)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, r)
}

func (h *HTTP) refresh(c echo.Context) error {
	r, err := h.svc.Refresh(c, c.Param("token"))
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, r)
}

func (h *HTTP) me(c echo.Context) error {
	user, err := h.svc.Me(c)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, user)
}
