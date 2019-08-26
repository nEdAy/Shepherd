package jwt

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/nEdAy/Shepherd/pkg/config"
	"github.com/pkg/errors"
	"time"
)

var (
	hmacSecret       = []byte(config.App.HmacSecret)
	TokenExpired     = errors.New("Token is expired")
	TokenNotValidYet = errors.New("Token not active yet")
	TokenMalformed   = errors.New("That's not even a token")
	TokenInvalid     = errors.New("Couldn't handle this token")
)

const KeyUserId = "USER_ID"

type CustomClaims struct {
	UserId uint
	jwt.StandardClaims
}

// creating, signing, and encoding a JWT token using the HMAC signing
func CreateToken(userId uint) (string, error) {
	// Create a new token object, specifying signing method and the claims
	// you would like it to contain.
	claims := CustomClaims{
		userId,
		jwt.StandardClaims{
			// 过期时间.通常与Unix UTC时间做对比过期后token无效
			ExpiresAt: time.Now().Add(3 * time.Hour).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString(hmacSecret)
	return tokenString, err
}

// parsing and validating a token using the HMAC signing
func ParseToken(tokenString string) (claims *CustomClaims, err error) {
	// Parse takes the token string and a function for looking up the key. The latter is especially
	// useful if you use multiple keys for your application.  The standard is to use 'kid' in the
	// head of the token to identify which key to use, but the parsed token (head and claims) is provided
	// to the callback, providing flexibility.
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return hmacSecret, nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, TokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				return nil, TokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, TokenNotValidYet
			} else {
				return nil, TokenInvalid
			}
		}
	}
	if token != nil {
		if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
			return claims, nil
		}
	}
	return nil, TokenInvalid
}
