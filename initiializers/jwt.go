package initiializers

import (
	"jwtAuthentication/models"
	"log"
	"time"

	"github.com/golang-jwt/jwt"
)

func GenerateJwtToken() (string, error) {
	mySignedKey := []byte(SECRET_KEY)

	User := models.CustomClaim{
		Token: "creating token",
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().AddDate(2400, 39, 30).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, User)

	ss, err := token.SignedString(mySignedKey)
	if err != nil {
		log.Println(err)
		return "token invalid", err
	}

	log.Println("token generated", ss)

	return ss, nil
}

func CheckJWTToken(token string) string {

	jwtToken, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		return []byte(SECRET_KEY), nil
	})
	if err != nil {
		return "your token is invalid"
	}

	if jwtToken.Valid {
		log.Println("allowww")
		return "allow"
	}

	return "unauthorized"
}
