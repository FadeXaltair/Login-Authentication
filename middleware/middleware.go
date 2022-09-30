package middleware

import (
	"jwtAuthentication/initiializers"

	"github.com/gin-gonic/gin"
)


// custom token checker ...
func AuthRequire(c *gin.Context) {

	yourToken := initiializers.Tokenn

	initiializers.CheckJWTToken(yourToken)
}
