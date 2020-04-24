package v1

import (
	"github.com/dgrijalva/jwt-go"
)

var (
	key = []byte("HappySailaSalt420")
)

// CustomClaims is our custom metadata, which will be hashed
// and sent as the second segment in our JWT
type CustomClaims struct {
	jwt.StandardClaims
}

type AuthToken struct {}

// Decode a token string into a token object
func (srv *AuthToken) Decode(tokenString string) (*CustomClaims, error) {

	// Parse the token
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return key, nil
	})

	if err!=nil {
		return nil, err
	}

	// Validate the token and return the custom claims
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, err
	}
}

// Encode a claim into a JWT
func (srv *AuthToken) Encode() (string, error) {

	// Create the Claims
	claims := CustomClaims{
		jwt.StandardClaims{
			Issuer:    "HappySaila",
		},
	}

	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign token and return
	return token.SignedString(key)
}
