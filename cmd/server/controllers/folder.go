package controllers
// import packages
import (
	"net/http"
	"encoding/json"
	"database/sql"

	"40104/SecretsManager/cmd/server/models"
)
// Get folder controller 
func (c *Controller) Get_Folder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json") // Set Header for the client
	
	req_folder := models.Folder{} 
    json.NewDecoder(r.Body).Decode(&req_folder) // Parse json body to the folder structure

	folder := c.DBModel.Get_Folder(req_folder.ID) // Get folder from DB
	json.NewEncoder(w).Encode(folder) // Send folder data to the client
}
// Get folders controller 
func (c *Controller) Get_Folders(w http.ResponseWriter, r *http.Request) {
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

	req_folder := models.Folder{}
    json.NewDecoder(r.Body).Decode(&req_folder) // Parse json body to the folder structure
	// Get folders from DB
	folders := c.DBModel.Get_Folders_By_Owner_and_Parrent_Folder(c.DBModel.Get_User_By_UserName(claim.UserName).ID, c.DBModel.Get_Folder_By_Name(req_folder.Parrent_Folder.Name).ID)

	json.NewEncoder(w).Encode(folders) //Send folders data to the client
}
// Create new folder controller 
func (c *Controller) Add_Folder(w http.ResponseWriter, r *http.Request) {
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
	
	folder := models.Folder{}
    json.NewDecoder(r.Body).Decode(&folder)  // Parse json body to the folder structure

	folder.Owner_ID = c.DBModel.Get_User_By_UserName(claim.UserName).ID // Add owner data from the request
	parrent_folder_id := sql.NullInt64{}
	if folder.Parrent_Folder.Name != "NULL" {
		parrent_folder_id = sql.NullInt64{int64(c.DBModel.Get_Folder_By_Name(folder.Parrent_Folder.Name).ID), true} // Add parrent folder data from the request
	} 
	folder.Parrent_Folder_ID = parrent_folder_id
	c.DBModel.Add_Folder(folder)  // Add new folder to DB
	
}
// Edit existing folder controller 
func (c *Controller) Put_Folder(w http.ResponseWriter, r *http.Request) {
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
	
    folder := models.Folder{}
    json.NewDecoder(r.Body).Decode(&folder) // Parse json body to the folder structure

	folder.Owner_ID = c.DBModel.Get_User_By_UserName(claim.UserName).ID // Add owner data from the request
	parrent_folder_id := sql.NullInt64{}
	if folder.Parrent_Folder.Name != "NULL" {
		parrent_folder_id = sql.NullInt64{int64(c.DBModel.Get_Folder_By_Name(folder.Parrent_Folder.Name).ID), true} // Add parrent folder data from the request
	} 
	folder.Parrent_Folder_ID = parrent_folder_id

	c.DBModel.Put_Folder(folder) // Edit folder in the DB
}
// Delete folder controller 
func (c *Controller) Delete_Folder(w http.ResponseWriter, r *http.Request) {
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

	folder := models.Folder{}
    json.NewDecoder(r.Body).Decode(&folder) // Parse json body to the folder structure

	c.DBModel.Delete_Folder(folder.ID,  c.DBModel.Get_User_By_UserName(claim.UserName).ID) // Delete folder in the DB
}