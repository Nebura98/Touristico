package usecases

import (
	"crypto"
	"net/http"
	"os"
	"time"

	repo "main/internal/adapters/repositories"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type UserClaims struct {
	Id    string `json:"id"`
	First string `json:"first"`
	Last  string `json:"last"`
	jwt.ClaimStrings
}

func SignUp() {

}

func SignIn(c gin.Context) {

	user := repo.GetUser("")

	token, err := generateToken()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create token",
		})
		return
	}

	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", token, 3600*24*30, "", "", false, true)
	c.JSON(http.StatusOK, gin.H{})
}

func generateToken() (string, error) {
	t := jwt.NewWithClaims(jwt.SigningMethodES256, jwt.MapClaims{
		"sub": crypto.MD5,
		"exp": time.Now().Add(time.Hour * 25 * 30).Unix(),
	})

	token, err := t.SignedString([]byte(os.Getenv("JWT_SECRET")))

	return token, err
}

func NewRefreshToken(claims jwt.Claims) (string, error) {
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return refreshToken.SignedString([]byte(os.Getenv("TOKEN_SECRET")))
}
