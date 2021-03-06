package jwt

import (
"net/http"
"strings"
"time"

"github.com/vasarostik/go_blog/pkg/utl/model"

"github.com/labstack/echo"

jwt "github.com/dgrijalva/jwt-go"
)

// New generates new JWT service necessery for auth middleware
func New(secret, algo string, d int) *Service {
	signingMethod := jwt.GetSigningMethod(algo)
	if signingMethod == nil {
		panic("invalid jwt signing method")
	}
	return &Service{
		key:      []byte(secret),
		algo:     signingMethod,
		duration: time.Duration(d) * time.Minute,
	}
}

// Service provides a Json-Web-Token authentication implementation
type Service struct {
	// Secret key used for signing.
	key []byte

	// Duration for which the jwt token is valid.
	duration time.Duration

	// JWT signing algorithm
	algo jwt.SigningMethod
}

// MWFunc makes JWT implement the Middleware interface.
func (j *Service) MWFunc() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			token, err := j.ParseToken(c)
			if err != nil || !token.Valid {
				return c.NoContent(http.StatusUnauthorized)
			}

			claims := token.Claims.(jwt.MapClaims)

			id := int(claims["id"].(float64))
			username := claims["u"].(string)
			role := go_blog.AccessRole(claims["r"].(float64))

			c.Set("id", id)
			c.Set("username", username)
			c.Set("role", role)

			return next(c)
		}
	}
}


func (j *Service) MWFuncURL() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			token, err := j.ParseTokenURL(c)
			if err != nil || !token.Valid {
				return c.NoContent(http.StatusUnauthorized)
			}

			claims := token.Claims.(jwt.MapClaims)

			id := int(claims["id"].(float64))
			username := claims["u"].(string)
			role := go_blog.AccessRole(claims["r"].(float64))

			c.Set("id", id)
			c.Set("username", username)
			c.Set("role", role)

			return next(c)
		}
	}
}

// ParseToken parses token from Authorization header
func (j *Service) ParseToken(c echo.Context) (*jwt.Token, error) {

	token := c.Request().Header.Get("Authorization")

	if token == "" {
		return nil, go_blog.ErrGeneric
	}
	parts := strings.SplitN(token, " ", 2)
	if !(len(parts) == 2 && parts[0] == "Bearer") {
		return nil, go_blog.ErrGeneric
	}

	return jwt.Parse(parts[1], func(token *jwt.Token) (interface{}, error) {
		if j.algo != token.Method {
			return nil, go_blog.ErrGeneric
		}
		return j.key, nil
	})

}

// ParseTokenURL parses token from URL
func (j *Service) ParseTokenURL(c echo.Context) (*jwt.Token, error) {

	token := c.QueryParam("token")

	if token == "" {
		return nil, go_blog.ErrGeneric
	}

	return jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {

		if j.algo != token.Method {
			return nil, go_blog.ErrGeneric
		}
		return j.key, nil
	})

}

// GenerateToken generates new JWT token and populates it with user data
func (j *Service) GenerateToken(u *go_blog.User) (string, string, error) {
	expire := time.Now().Add(j.duration)

	token := jwt.NewWithClaims((j.algo), jwt.MapClaims{
		"id":  u.ID,
		"u":   u.Username,
		"r":   u.Role.AccessLevel,
		"exp": expire.Unix(),
	})

	tokenString, err := token.SignedString(j.key)

	return tokenString, expire.Format(time.RFC3339), err
}
