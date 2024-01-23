package controllers
// import packages
import (
	"net/http"
	"log"
	"encoding/json"

	"40104/SecretsManager/cmd/server/models"
)
// Get secret controller 
func (c *Controller) Get_Secret(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json") // Set Header for the client

	// Check permissions
	token_string := r.Header.Get("Authorization") //Get auth header
    if token_string == "" {
		w.WriteHeader(http.StatusUnauthorized) // Set unauthorized status
    }
    token_string = token_string[len("Bearer "):] 
    err, claim:= c.VerifyToken(token_string) // Verify user token
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized) // Set unauthorized status
	}

	req_secret := models.Secret{}
    json.NewDecoder(r.Body).Decode(&req_secret) // Parse json body to the secret structure
	
	secret := c.DBModel.Get_Secret(req_secret.ID, c.DBModel.Get_User_By_UserName(claim.UserName).ID) // Get secret from the DB
	json.NewEncoder(w).Encode(secret) //Send secret data to the client
}
// Get secrets controller 
func (c *Controller) Get_Secrets(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json") // Set Header for the client
	// Check permissions
	token_string := r.Header.Get("Authorization") //Get auth header
    if token_string == "" {
		w.WriteHeader(http.StatusUnauthorized) // Set unauthorized status
    }
    token_string = token_string[len("Bearer "):] 
    err, claim:= c.VerifyToken(token_string) // Verify user token
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized) // Set unauthorized status
	}

	secret := models.Secret{}
    json.NewDecoder(r.Body).Decode(&secret) // Parse json body to the secret structure
	secrets := c.DBModel.Get_Secrets_By_Owner_and_Folder(
		c.DBModel.Get_User_By_UserName(claim.UserName).ID, // Get owner id
		c.DBModel.Get_Folder_By_Name(secret.Folder.Name).ID) // Get folder id
	
	json.NewEncoder(w).Encode(secrets) //Send secrets data to the client
}
// Create new secret controller 
func (c *Controller) Add_Secret(w http.ResponseWriter, r *http.Request) {
    // Check permissions
	token_string := r.Header.Get("Authorization") //Get auth header
    if token_string == "" {
		w.WriteHeader(http.StatusUnauthorized) // Set unauthorized status
    }
    token_string = token_string[len("Bearer "):] 
    err, claim:= c.VerifyToken(token_string) // Verify user token
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized) // Set unauthorized status
	}

	secret := models.Secret{}
    json.NewDecoder(r.Body).Decode(&secret) // Parse json body to the secret structure

	secret.Owner_ID = c.DBModel.Get_User_By_UserName(claim.UserName).ID // Get owner id
	secret.Folder_ID = c.DBModel.Get_Folder_By_Name(secret.Folder.Name).ID // Get folder id
	
	c.DBModel.Add_Secret(secret) // Create new secret 
}
// Generate new secret controller 
func (c *Controller) Generate_Secret(w http.ResponseWriter, r *http.Request) {
    // Check permissions
	token_string := r.Header.Get("Authorization") //Get auth header
    if token_string == "" {
		w.WriteHeader(http.StatusUnauthorized) // Set unauthorized status
    }
    token_string = token_string[len("Bearer "):] 
    err, claim:= c.VerifyToken(token_string) // Verify user token
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized) // Set unauthorized status
	}
	
	secret := models.Secret{}
    json.NewDecoder(r.Body).Decode(&secret) // Parse json body to the secret structure

	secret.Owner_ID = c.DBModel.Get_User_By_UserName(claim.UserName).ID // Get owner id
	secret.Folder_ID = c.DBModel.Get_Folder_By_Name(secret.Folder.Name).ID // Get folder id
	
	c.DBModel.Add_Secret(secret) // Create new secret 
}
// Edit existing secret controller 
func (c *Controller) Put_Secret(w http.ResponseWriter, r *http.Request) {
    // Check permissions
	token_string := r.Header.Get("Authorization") //Get auth header
    if token_string == "" {
		w.WriteHeader(http.StatusUnauthorized) // Set unauthorized status
    }
    token_string = token_string[len("Bearer "):] 
    err, claim:= c.VerifyToken(token_string) // Verify user token
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized) // Set unauthorized status
	}

	secret := models.Secret{}
    json.NewDecoder(r.Body).Decode(&secret) // Parse json body to the secret structure

	secret.Owner_ID = c.DBModel.Get_User_By_UserName(claim.UserName).ID // Get owner id
	secret.Folder_ID = c.DBModel.Get_Folder_By_Name(secret.Folder.Name).ID // Get folder id

	c.DBModel.Put_Secret(secret)
}
// Delete secret controller 
func (c *Controller) Delete_Secret(w http.ResponseWriter, r *http.Request) {
	// Check permissions
	token_string := r.Header.Get("Authorization") //Get auth header
    if token_string == "" {
		w.WriteHeader(http.StatusUnauthorized) // Set unauthorized status
    }
    token_string = token_string[len("Bearer "):] 
    err, claim:= c.VerifyToken(token_string) // Verify user token
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized) // Set unauthorized status
	}

	secret := models.Secret{}
    json.NewDecoder(r.Body).Decode(&secret) // Parse json body to the secret structure

	c.DBModel.Delete_Secret(secret.ID, c.DBModel.Get_User_By_UserName(claim.UserName).ID) // Delete secret 
}
