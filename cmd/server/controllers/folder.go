package controllers

import (
	"net/http"
	"encoding/json"
	"database/sql"

	"40104/SecretsManager/cmd/server/models"
)

func (c *Controller) Get_Folder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	

	req_folder := models.Folder{}
    json.NewDecoder(r.Body).Decode(&req_folder)

	folder := c.DBModel.Get_Folder(req_folder.ID)
	json.NewEncoder(w).Encode(folder)
}

func (c *Controller) Get_Folders(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	token_string := r.Header.Get("Authorization")
    if token_string == "" {
		w.WriteHeader(http.StatusUnauthorized)
    }
    token_string = token_string[len("Bearer "):]
    err, claim:= c.VerifyToken(token_string)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
	}

	req_folder := models.Folder{}
    json.NewDecoder(r.Body).Decode(&req_folder)
	folders := c.DBModel.Get_Folders_By_Owner_and_Parrent_Folder(c.DBModel.Get_User_By_UserName(claim.UserName).ID, c.DBModel.Get_Folder_By_Name(req_folder.Parrent_Folder.Name).ID)

	json.NewEncoder(w).Encode(folders)
}

func (c *Controller) Add_Folder(w http.ResponseWriter, r *http.Request) {
	token_string := r.Header.Get("Authorization")
    if token_string == "" {
		w.WriteHeader(http.StatusUnauthorized)
    }
    token_string = token_string[len("Bearer "):]
    err, claim:= c.VerifyToken(token_string)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
	}
	
	folder := models.Folder{}
    json.NewDecoder(r.Body).Decode(&folder)

	folder.Owner_ID = c.DBModel.Get_User_By_UserName(claim.UserName).ID
	parrent_folder_id := sql.NullInt64{}
	if folder.Parrent_Folder.Name != "NULL" {
		parrent_folder_id = sql.NullInt64{int64(c.DBModel.Get_Folder_By_Name(folder.Parrent_Folder.Name).ID), true}
	} 
	folder.Parrent_Folder_ID = parrent_folder_id
	c.DBModel.Add_Folder(folder)
	
}

func (c *Controller) Put_Folder(w http.ResponseWriter, r *http.Request) {
	token_string := r.Header.Get("Authorization")
    if token_string == "" {
		w.WriteHeader(http.StatusUnauthorized)
    }
    token_string = token_string[len("Bearer "):]
    err, claim:= c.VerifyToken(token_string)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
	}
	
    folder := models.Folder{}
    json.NewDecoder(r.Body).Decode(&folder)

	folder.Owner_ID = c.DBModel.Get_User_By_UserName(claim.UserName).ID
	parrent_folder_id := sql.NullInt64{}
	if folder.Parrent_Folder.Name != "NULL" {
		parrent_folder_id = sql.NullInt64{int64(c.DBModel.Get_Folder_By_Name(folder.Parrent_Folder.Name).ID), true}
	} 
	folder.Parrent_Folder_ID = parrent_folder_id

	c.DBModel.Put_Folder(folder)
}

func (c *Controller) Delete_Folder(w http.ResponseWriter, r *http.Request) {
	token_string := r.Header.Get("Authorization")
    if token_string == "" {
		w.WriteHeader(http.StatusUnauthorized)
    }
    token_string = token_string[len("Bearer "):]
    err, claim:= c.VerifyToken(token_string)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
	}

	folder := models.Folder{}
    json.NewDecoder(r.Body).Decode(&folder)

	c.DBModel.Delete_Folder(folder.ID,  c.DBModel.Get_User_By_UserName(claim.UserName).ID)
}