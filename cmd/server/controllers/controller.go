package controllers

import (
	"net/http"
	"fmt"
    "encoding/json"

    "github.com/golang-jwt/jwt/v5"
    "golang.org/x/crypto/bcrypt"

    "40104/SecretsManager/cmd/server/models"
)

type Controller struct {
	DBModel *models.DBModel
    JWT_secret []byte
}

type Claims struct {
	UserName string `json:"username"`
    Role string `json:"role"`
	jwt.RegisteredClaims
}

func (_ *Controller) Home(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w,"Welcome to Secrets Manager application.")
}

func (c *Controller) Login(w http.ResponseWriter, r *http.Request) {
    auth_user := models.User{}
    json.NewDecoder(r.Body).Decode(&auth_user)
    db_user :=c.DBModel.Get_User_By_UserName(auth_user.UserName)
    fmt.Println(auth_user.UserName)
    if err := bcrypt.CompareHashAndPassword([]byte(db_user.Password),[]byte(auth_user.Password)); err == nil {
        token_string := c.CreateToken(auth_user.UserName, c.DBModel.Get_Role(db_user.Role_ID).Name)
        w.WriteHeader(http.StatusOK)
        http.SetCookie(w, &http.Cookie{
            Name:    "token",
            Value:   token_string,
        })
        fmt.Fprint(w, token_string)
    } else {
        w.WriteHeader(http.StatusUnauthorized)
    }
}

