package middleware

import (
	"fmt"
	"jwtAuthentication/initiializers"
	"jwtAuthentication/models"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)


func AuthRequire(c *gin.Context) {

	///get the cookies
  tokenString, err :=c.Cookie("authorize")

  if err != nil {
	c.JSON(http.StatusUnauthorized , gin.H{
		"error": "Failed to validsate",
	})
  }
	//decode and alidate 

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {

		if err !=nil {
			log.Println(err)
		}

		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, 
			fmt.Errorf("unexpected signing method: %v", token.Header["a lg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(os.Getenv("SECRET")), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			//cheeck the user with token sub 
      if float64(time.Now().Unix()) > claims["exp"].(float64){
	  			c.AbortWithStatus(http.StatusUnauthorized)
	  }

	  // find the user 

	  var user models.User

		  initiializers.DB.First(&user , claims ["sub"])

		  if user.ID == 0 {
			c.AbortWithStatus(http.StatusUnauthorized)
		  }

		  c.Set("user", user)
	       //attach the req 
    	 c.Next()

		} else {
		c.JSON(http.StatusUnauthorized , gin.H{
			"error": "Failed to validsate",
		})
	}



}