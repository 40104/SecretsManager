package controllers
// import packages
import (
	"net/http"
	"encoding/json"

	"40104/SecretsManager/cmd/server/models"
)
// Get user controller 
func (c *Controller) Get_User(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json") //Set head for client

	req_user := models.User{}
    json.NewDecoder(r.Body).Decode(&req_user) // Parse json body to the user structure

	user:= c.DBModel.Get_User(req_user.ID)  // Get user from DB
	json.NewEncoder(w).Encode(user) // Send role data to the client
}
// Create new user controller 
func (c *Controller) Add_User(w http.ResponseWriter, r *http.Request) {
	user := models.User{}
    json.NewDecoder(r.Body).Decode(&user) // Parse json body to the user structure

	user.Role_ID = c.DBModel.Get_Role_By_Name(user.Role.Name).ID // Get role id by name 
	c.DBModel.Add_User(user) // Create new user
}
// Edit existing user controller 
func (c *Controller) Put_User(w http.ResponseWriter, r *http.Request) {
	user := models.User{}
    json.NewDecoder(r.Body).Decode(&user) // Parse json body to the user structure

	user.Role_ID = c.DBModel.Get_Role_By_Name(user.Role.Name).ID // Get role id by name 
	c.DBModel.Put_User(user) // Edit existing user 
}
// Delete user controller
func (c *Controller) Delete_User(w http.ResponseWriter, r *http.Request) {
	user := models.User{}
    json.NewDecoder(r.Body).Decode(&user) // Parse json body to the user structure

	c.DBModel.Delete_User(user.ID) // Delete user 
}