package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/Tympanix/artoodetoo/config"
	jwt "github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

type credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// Authenticate is a middleware used to authentication requests
func auth(fn func(w http.ResponseWriter, r *http.Request)) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		authentication := r.Header.Get("Authentication")

		token, err := jwt.Parse(authentication, getSecret)

		if err != nil || !token.Valid {
			http.Error(w, "Wrong token", http.StatusUnauthorized)
			return
		}

		fn(w, r)
	})
}

func getSecret(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
	}
	return []byte(config.Secret), nil
}

func init() {
	API.HandleFunc("/login", login).Methods("POST")
}

func login(w http.ResponseWriter, r *http.Request) {
	cred := new(credentials)

	dec := json.NewDecoder(r.Body)
	if err := dec.Decode(cred); err != nil {
		http.Error(w, "Missing credentials", http.StatusInternalServerError)
		return
	}

	hash, ok := config.Passwords[cred.Username]

	if !ok {
		unauthorize(w)
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(cred.Password)); err != nil {
		unauthorize(w)
		return
	}

	claims := &jwt.StandardClaims{
		ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(config.Secret))

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write([]byte(tokenString))
}

func unauthorize(w http.ResponseWriter) {
	http.Error(w, "Wrong credentials", http.StatusUnauthorized)
}
