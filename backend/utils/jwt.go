package utils

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

// TODO: use env variables for secret key later
// TODO: Add roles/claims maybe
// TODO: Add validate methods

// jwtKey: Secret key for signing the JWT
var jwtKey = []byte("someSecret")

// tokens: Slice to store one or multiple JWTs (Optional, not in use atm) TODO: Remove if not using later
var tokens []string

// Claims : Custom claim to store JWT payload(claims)
type Claims struct {
	Username             string `json:"username"`
	jwt.RegisteredClaims        //Embeds standard JWT claims like ExpiresAt
}

// GenerateJWT : Generates a signed JWT for the user
// Will consist of: header, payload and signature
func GenerateJWT(username string) (string, error) {
	expirationTime := time.Now().Add(5 * time.Minute) // Expiration time for token (5 min from now)  TODO: Test if expiration time works

	// The payload
	claims := &Claims{
		Username: username,
		//Role     string `json:"role"`, // If needed later. Also pass it as param.
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime), //Converts time.Time to the format used in JWT claims (JWT numeric format)
		},
	}

	// creates new JWT with signing method HS256 and specified claims (header is created based on signing method used)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Signs the token with the provided key
	// The token will be sent to the frontend. Users can now use token to authenticate themselves instead of querying Database on each request.
	// The server knows the jwtKey, so it will know when someone tries to use a fake one or if the key has been tampered with.
	return token.SignedString(jwtKey)
}

// getSecretKey : Validates the signing method and returns the secret key for verifying a token's signature
func getSecretKey(token *jwt.Token) (interface{}, error) {
	// Check if the signing method is HMAC (Currently using HS256)
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
	}
	// Return the secret key used for signature verification //TODO: Fix secret key here too
	return []byte("someSecret"), nil
}

// VerifyJWT Checks if JWT is valid
func VerifyJWT(tokenString string) (*Claims, error) {
	// Decode tokenString, and verifies signature using returned key from the Keyfunc (getSecretKey)
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, getSecretKey)
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(*Claims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid token")
	}
	//TODO: Remove the print
	fmt.Println(claims.Username, claims.RegisteredClaims.Issuer)
	return claims, nil
}
