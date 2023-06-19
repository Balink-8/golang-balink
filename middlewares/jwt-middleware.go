package middlewares

import (
	"log"
	"net/http"
	"os"
	"strings"
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

func IsUser(c echo.Context) (int, error) {
	token := c.Request().Header.Get("Authorization")
	if token == "" {
		return 0, echo.NewHTTPError(400, "Missing Token")
	}

	// Extract the token from the "Bearer <token>" format
	tokenParts := strings.Split(token, " ")
	if len(tokenParts) != 2 || strings.ToLower(tokenParts[0]) != "bearer" {
		return 0, echo.NewHTTPError(400, "Invalid Token Format")
	}

	// Parse and validate the JWT token
	jwtToken, err := jwt.Parse(tokenParts[1], func(token *jwt.Token) (interface{}, error) {
		// Replace "your-secret-key" with the actual secret key used to sign the tokens
		// You may need to retrieve the secret key from your configuration or environment variables
		return []byte(os.Getenv("JWT_KEY")), nil
	})
	if err != nil {
		return 0, echo.NewHTTPError(400, "Invalid Token")
	}

	claims, ok := jwtToken.Claims.(jwt.MapClaims)
	if !ok || !jwtToken.Valid {
		return 0, echo.NewHTTPError(401, "Unauthorized")
	}
	// Extract the admin ID from the token's payload
	id := claims["id"].(float64)

	return int(id), nil
}

// butuh func buat reset jwt token
