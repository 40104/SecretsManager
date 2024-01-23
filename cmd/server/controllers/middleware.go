package controllers
// import packages
import (
	"net/http"
	"log"
	"time"
    "strings"
    
    "github.com/golang-jwt/jwt/v5"
)
// Create token function
func (c *Controller) CreateToken(username string, role string) (string) {
    //Creare new claim
	claims := &Claims{
		UserName: username,
        Role:   role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(90 * time.Minute)), // Date expiration time 
		},
	}
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims) //Init token
    
	token_string, err := token.SignedString(c.JWT_secret) //Singin token with JWT secret key
	if err != nil {
    	log.Fatal(err) //check error
    }
	return token_string //return token
}
// Verify token function
func (c *Controller) VerifyToken(token_string string ) (error, *Claims) {
    claims := &Claims{}
	token, err := jwt.ParseWithClaims(token_string, claims, func(token *jwt.Token) (interface{}, error) { //Parse token to the claim struct
		return c.JWT_secret, nil 
	})
	if err != nil {
		return err, claims //chech error
	}
	if !token.Valid {
		return err, claims //chech validation
	}
	return err, claims
}
//Base user middleware
func (c *Controller) UserAuthMiddleWare(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        token_string := r.Header.Get("Authorization") //Get Header
        if token_string == "" {
            w.WriteHeader(http.StatusUnauthorized) // Set unauthorized status 
            return
        }
        token_string = strings.TrimPrefix(token_string, "Bearer ")
        err, claim:= c.VerifyToken(token_string) //Verify token
        if err != nil {
            w.WriteHeader(http.StatusUnauthorized) // Set unauthorized status 
            return
        }
        
        if claim.Role == "Administrator" || claim.Role == "User"{ //check role
            next.ServeHTTP(w, r)
            return
        }
        
        w.WriteHeader(http.StatusUnauthorized) // Set unauthorized status 
	})
}
//Admin middleware
func (c *Controller) AdminAuthMiddleWare(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        token_string := r.Header.Get("Authorization") //Get Header
        if token_string == "" {
            w.WriteHeader(http.StatusUnauthorized) // Set unauthorized status 
            return
        }
        token_string = strings.TrimPrefix(token_string, "Bearer ")
        err, claim:= c.VerifyToken(token_string) //Verify token
        if err != nil {
            w.WriteHeader(http.StatusUnauthorized) // Set unauthorized status 
            return
        }
        if claim.Role == "Administrator" { //check role
            next.ServeHTTP(w, r)
        }
        
        w.WriteHeader(http.StatusForbidden) // Set forbidden status 
	})
}





