package middlewares

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

type JWTS interface {
	CreateJWTToken(id uint, name string) (string, error)
	LogoutJWTToken(c echo.Context) error
}

type jwtS struct {
	issuer    string
	secretKey string
}

func NewJWTS() JWTS {
	err := godotenv.Load()
	if err != nil {
		err = godotenv.Load("../.env")
		if err != nil {
			log.Println(err)
		}

		log.Println(err)
	}
	return &jwtS{
		issuer:    "qwerty",
		secretKey: os.Getenv("JWT_KEY"),
	}
}

func (j *jwtS) CreateJWTToken(id uint, name string) (string, error) {
	claims := jwt.MapClaims{
		"id":   id,
		"name": name,
		"exp":  time.Now().Add(time.Minute * 5).Unix(),
		"iss":  j.issuer,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(j.secretKey))
}

func (j jwtS) LogoutJWTToken(c echo.Context) error {
	claims := jwt.MapClaims{
		"exp": time.Now().Add(-1 * time.Hour).Unix(),
		"iss": j.issuer,
	}

	_ = jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return c.NoContent(http.StatusOK)
}

// butuh func buat reset jwt token
