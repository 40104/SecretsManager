package controllers
// import packages
import (
	"net/http"
	"fmt"
    "encoding/json"

    "github.com/golang-jwt/jwt/v5"
    "golang.org/x/crypto/bcrypt"

    "40104/SecretsManager/cmd/server/models"
)
// Controller class
type Controller struct {
	DBModel *models.DBModel
    JWT_secret []byte
}
// Claims struct
type Claims struct {
	UserName string `json:"username"`
    Role string `json:"role"`
	jwt.RegisteredClaims
}
//Base home controller
func (_ *Controller) Home(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w,"Welcome to Secrets Manager application.")
}
//Login controller
func (c *Controller) Login(w http.ResponseWriter, r *http.Request) {
    auth_user := models.User{}
    json.NewDecoder(r.Body).Decode(&auth_user) // Parse json body to the user structure
    db_user :=c.DBModel.Get_User_By_UserName(auth_user.UserName) //Get user id
    if err := bcrypt.CompareHashAndPassword([]byte(db_user.Password),[]byte(auth_user.Password)); err == nil { //Compare hash from auth and DB
        token_string := c.CreateToken(auth_user.UserName, c.DBModel.Get_Role(db_user.Role_ID).Name) //Create token
        w.WriteHeader(http.StatusOK)
        http.SetCookie(w, &http.Cookie{ //Set token cookie
            Name:    "token",
            Value:   token_string,
        })
        fmt.Fprint(w, token_string) //Send token
    } else {
        w.WriteHeader(http.StatusUnauthorized)
    }
}

