package security

import (
	"github.com/dgrijalva/jwt-go"
	"log"
	"net/http"
	"time"
)

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

var jwtKey = []byte("Demo-Server-Key")

const TimeOut = 30

func AddCookie(w http.ResponseWriter, name string, value string) {
	expire := time.Now().Add(time.Minute * TimeOut)
	cookie := http.Cookie{
		Name:     name,
		Value:    value,
		Expires:  expire,
		Secure:   true,
		HttpOnly: true,
		Path:     "/",
		Domain:   "localhost",
	}
	http.SetCookie(w, &cookie)
}

func GenerateToken(username string) string {
	var tokenString string
	var err error

	expirationTime := time.Now().Add(TimeOut * time.Minute)
	claims := &Claims{Username: username, StandardClaims: jwt.StandardClaims{ExpiresAt: expirationTime.Unix()}}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	if tokenString, err = token.SignedString(jwtKey); err != nil {
		log.Panic(err)
		return ""
	}
	return tokenString
}

func ValidateToken(token string) (bool, string) {
	var err error
	var tkn *jwt.Token

	claims := &Claims{}
	if tkn, err = jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) { return jwtKey, nil }); err != nil {
		log.Panic(err)
		return false, ""
	}
	if tkn.Valid {
		return true, claims.Username
	} else if err == jwt.ErrSignatureInvalid {
		return false, ""
	} else {
		return false, ""
	}
}
