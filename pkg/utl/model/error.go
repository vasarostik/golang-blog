package go_blog

import (
	"errors"
	"net/http"

	"github.com/labstack/echo"
)

var (
	// ErrGeneric is used for testing purposes and for errors handled later in the callstack
	ErrGeneric = errors.New("generic error")

	// ErrBadRequest (400) is returned for bad request (validation)
	ErrBadRequest = echo.NewHTTPError(400)

	// ErrUnauthorized (401) is returned when user is not authorized
	ErrUnauthorized = echo.ErrUnauthorized

	ErrAlreadyExists = echo.NewHTTPError(http.StatusInternalServerError, "Title already exists.")

	ErrInvalidCredentials = echo.NewHTTPError(http.StatusUnauthorized, "Username or password does not exist")

	ErrIncorrectPassword = echo.NewHTTPError(http.StatusBadRequest, "incorrect old password")

	ErrInsecurePassword  = echo.NewHTTPError(http.StatusBadRequest, "insecure password")

	ErrUsernameAlreadyExists = echo.NewHTTPError(http.StatusInternalServerError, "Username already exists.")

	ErrPasswordsNotMaching = echo.NewHTTPError(http.StatusBadRequest, "passwords do not match")


)
