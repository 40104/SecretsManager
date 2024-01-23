package controllers
// import packages
import (
	"net/http"
	"encoding/json"

	"40104/SecretsManager/cmd/server/models"
)
// Get role controller 
func (c *Controller) Get_Role(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json") // Set Head for client

	req_role := models.Role{}
    json.NewDecoder(r.Body).Decode(&req_role) // Parse json body to the role structure

	role := c.DBModel.Get_Role(req_role.ID) // Get role from DB
	json.NewEncoder(w).Encode(role) // Send role data to the client
}
// Create new role controller 
func (c *Controller) Add_Role(w http.ResponseWriter, r *http.Request) {
	role := models.Role{}
    json.NewDecoder(r.Body).Decode(&role) // Parse json body to the role structure
	
	c.DBModel.Add_Role(role) // Create new role
}
// Edit existing role controller 
func (c *Controller) Put_Role(w http.ResponseWriter, r *http.Request) {
	role := models.Role{}
    json.NewDecoder(r.Body).Decode(&role) // Parse json body to the role structure

	c.DBModel.Put_Role(role) // Edit existing role
}
// Delete role controller 
func (c *Controller) Delete_Role(w http.ResponseWriter, r *http.Request) {
	role := models.Role{}
    json.NewDecoder(r.Body).Decode(&role) // Parse json body to the role structure

	c.DBModel.Delete_Role(role.ID) // Delete role
}