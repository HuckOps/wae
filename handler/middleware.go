package handler

import (
	"net/http"
	"strings"
	"wae/pkg/oidc"
	"wae/pkg/restful"

	"github.com/gin-gonic/gin"
)

func OIDCMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authorization := c.Request.Header.Get("Authorization")
		if authorization == "" {
			c.JSON(http.StatusUnauthorized, restful.Restful[any]{
				Code:    restful.PermissionDenied,
				Message: "Can't get token in Headers",
			})
			c.Abort()
			return
		}

		tokenSplit := strings.Split(authorization, " ")
		if len(tokenSplit) != 2 {
			c.JSON(http.StatusBadRequest, restful.Restful[any]{
				Code:    restful.PermissionDenied,
				Message: "Can't get token in Headers",
			})
			c.Abort()
			return
		}
		token := tokenSplit[1]

		idToken, err := oidc.Verifier.Verify(c.Request.Context(), token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, restful.Restful[any]{
				Code:    restful.PermissionDenied,
				Message: "Invalid token",
			})
			c.Abort()
			return
		}

		var claims oidc.Claims
		if err := idToken.Claims(&claims); err != nil {
			c.JSON(http.StatusUnauthorized, restful.Restful[any]{
				Code:    restful.PermissionDenied,
				Message: "Invalid token claims",
			})
			c.Abort()
			return
		}

		ctx := oidc.NewContext(c.Request.Context(), &claims)
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}
