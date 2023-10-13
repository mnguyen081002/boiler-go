package middlewares

import (
	"erp/api/response"
	"erp/api_errors"
	"erp/domain"
	"fmt"
	"net/http"

	"github.com/pkg/errors"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func (e *GinMiddleware) Auth(authorization bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		auth := c.Request.Header.Get("Authorization")

		if auth == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, response.ResponseError{
				Message: "Unauthorized",
				Code:    api_errors.ErrUnauthorizedAccess,
			})
			return
		}
		c.Request.Header.Set("Authorization", auth)

		fmt.Println("XXXXXXXXXXXXX", c.Writer.Header())

		if !authorization {
			c.Next()
			return
		}

		c.Next()
	}
}

func parseToken(jwtToken string, secret string) (*domain.JwtClaims, error) {
	token, err := jwt.ParseWithClaims(jwtToken, &domain.JwtClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	if err != nil {
		if (err.(*jwt.ValidationError)).Errors == jwt.ValidationErrorExpired {
			return nil, errors.New(api_errors.ErrTokenExpired)
		}
		return nil, errors.Wrap(err, "cannot parse token")
	}

	if claims, OK := token.Claims.(*domain.JwtClaims); OK && token.Valid {
		return claims, nil
	}

	return nil, errors.New(api_errors.ErrTokenInvalid)
}
