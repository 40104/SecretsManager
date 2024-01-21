package controllers

import (
	"net/http"
	"log"
	"time"
    "strings"
    
    "github.com/golang-jwt/jwt/v5"
)

func (c *Controller) CreateToken(username string, role string) (string) {
	claims := &Claims{
		UserName: username,
        Role:   role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(90 * time.Minute)),
		},
	}
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    
	token_string, err := token.SignedString(c.JWT_secret)
	if err != nil {
    	log.Fatal(err)
    }
	return token_string
}

func (c *Controller) VerifyToken(token_string string ) (error, *Claims) {
    claims := &Claims{}
	token, err := jwt.ParseWithClaims(token_string, claims, func(token *jwt.Token) (interface{}, error) {
		return c.JWT_secret, nil
	})
	if err != nil {
		return err, claims
	}
	if !token.Valid {
		return err, claims
	}
	return err, claims
}

func (c *Controller) UserAuthMiddleWare(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        token_string := r.Header.Get("Authorization")
        if token_string == "" {
            w.WriteHeader(http.StatusUnauthorized)
            return
        }
        token_string = strings.TrimPrefix(token_string, "Bearer ")
        err, claim:= c.VerifyToken(token_string)
        if err != nil {
            w.WriteHeader(http.StatusUnauthorized)
            return
        }
        
        if claim.Role == "Administrator" || claim.Role == "User"{
            next.ServeHTTP(w, r)
            return
        }
        
        w.WriteHeader(http.StatusUnauthorized)
	})
}

func (c *Controller) AdminAuthMiddleWare(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        token_string := r.Header.Get("Authorization")
        if token_string == "" {
            w.WriteHeader(http.StatusUnauthorized)
            return
        }
        token_string = strings.TrimPrefix(token_string, "Bearer ")
        err, claim:= c.VerifyToken(token_string)
        if err != nil {
            w.WriteHeader(http.StatusUnauthorized)
            return
        }
        if claim.Role == "Administrator" {
            next.ServeHTTP(w, r)
        }
        
        w.WriteHeader(http.StatusForbidden)
	})
}




