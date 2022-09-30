package controller

import (
	"jwtAuthentication/initiializers"
	"jwtAuthentication/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Signup(c *gin.Context) {

	//-------------body to strore the value -----------//
	var body struct {
		Name            string `json:"name"`
		Email           string `json:"email"`
		Password        string `json:"password"`
		ConfirmPassword string `json:"confirm_password"`
	}
	c.Bind(&body)

	//--------------hash the store data ----------/

	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "message failer",
		})
		return
	}

	//------create user --------/

	user := models.User{Name: body.Name, Email: body.Email, Password: string(hash), ConfirmPassword: string(hash)}
	result := initiializers.DB.Create(&user)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "failed to connect",
		})
		return

	}

	//------------------if everything is right ------------------//
	c.JSON(http.StatusOK, gin.H{
		"positive": "your account has been created",
	})
}

func Login(c *gin.Context) {
	//---- bodyyy of the database ---\\\

	var body struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	c.Bind(&body)

	/// check the email in database ----///
	var user models.User
	initiializers.DB.First(&user, "email = ?", body.Email)

	if user.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid email",
		})
		return
	}
	/// comparing hashes //

	error := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))

	if error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid passs",
		})
		return
	}

	hsh, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)

	if err != nil {
		log.Println("error while encrypting")
	}

	logger := models.Login{
		Email:    body.Email,
		Password: string(hsh),
	}

	var loger models.Login

	initiializers.DB.First(&logger, "email = ?", body.Email)

	if loger.ID == 0 {
		result2 := initiializers.DB.Create(&logger)
		if result2.Error != nil {
			log.Println("error while editing it into logging databases")
		}
	}
	//---jwt create tokenn ///]]]]

	Toke, err := initiializers.GenerateJwtToken()
	if err != nil {
		log.Println(err)
	}

	initiializers.Tokenn = Toke
	c.JSON(http.StatusOK, gin.H{
		"token": "Your token generatted successfully",
	})

}

func Validate(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "you are loging successfully",
		"token":   "you are a valid token",
	})
}
