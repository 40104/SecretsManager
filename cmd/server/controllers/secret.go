package controllers

import (
	"net/http"
	"log"
	"encoding/json"

	"40104/SecretsManager/cmd/server/models"
)


func (c *Controller) Get_Secret(w http.ResponseWriter, r *http.Request) {
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

	req_secret := models.Secret{}
    json.NewDecoder(r.Body).Decode(&req_secret)
	
	secret := c.DBModel.Get_Secret(req_secret.ID, c.DBModel.Get_User_By_UserName(claim.UserName).ID)
	json.NewEncoder(w).Encode(secret)
}

func (c *Controller) Get_Secrets(w http.ResponseWriter, r *http.Request) {
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

	secret := models.Secret{}
    json.NewDecoder(r.Body).Decode(&secret)
	secrets := c.DBModel.Get_Secrets_By_Owner_and_Folder(
		c.DBModel.Get_User_By_UserName(claim.UserName).ID,
		c.DBModel.Get_Folder_By_Name(secret.Folder.Name).ID)
	
	json.NewEncoder(w).Encode(secrets)
}

func (c *Controller) Add_Secret(w http.ResponseWriter, r *http.Request) {
	token_string := r.Header.Get("Authorization")
    if token_string == "" {
		w.WriteHeader(http.StatusUnauthorized)
    }
    token_string = token_string[len("Bearer "):]
    err, claim:= c.VerifyToken(token_string)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
	}

	secret := models.Secret{}
    json.NewDecoder(r.Body).Decode(&secret)

	secret.Owner_ID = c.DBModel.Get_User_By_UserName(claim.UserName).ID
	secret.Folder_ID = c.DBModel.Get_Folder_By_Name(secret.Folder.Name).ID
	
	c.DBModel.Add_Secret(secret)
}

func (c *Controller) Generate_Secret(w http.ResponseWriter, r *http.Request) {
	token_string := r.Header.Get("Authorization")
    if token_string == "" {
		w.WriteHeader(http.StatusUnauthorized)
    }
    token_string = token_string[len("Bearer "):]
    err, claim:= c.VerifyToken(token_string)
	if err != nil {
    	log.Fatal(err)
    }
	
	secret := models.Secret{}
    json.NewDecoder(r.Body).Decode(&secret)

	secret.Owner_ID = c.DBModel.Get_User_By_UserName(claim.UserName).ID
	secret.Folder_ID = c.DBModel.Get_Folder_By_Name(secret.Folder.Name).ID

	c.DBModel.Add_Secret(secret)
}

func (c *Controller) Put_Secret(w http.ResponseWriter, r *http.Request) {
	token_string := r.Header.Get("Authorization")
    if token_string == "" {
		w.WriteHeader(http.StatusUnauthorized)
    }
    token_string = token_string[len("Bearer "):]
    err, claim:= c.VerifyToken(token_string)
	if err != nil {
    	log.Fatal(err)
    }

	secret := models.Secret{}
    json.NewDecoder(r.Body).Decode(&secret)

	secret.Owner_ID = c.DBModel.Get_User_By_UserName(claim.UserName).ID
	secret.Folder_ID = c.DBModel.Get_Folder_By_Name(secret.Folder.Name).ID

	c.DBModel.Put_Secret(secret)
}

func (c *Controller) Delete_Secret(w http.ResponseWriter, r *http.Request) {
	token_string := r.Header.Get("Authorization")
    if token_string == "" {
		w.WriteHeader(http.StatusUnauthorized)
    }
    token_string = token_string[len("Bearer "):]
    err, claim:= c.VerifyToken(token_string)
	if err != nil {
    	log.Fatal(err)
    }

	secret := models.Secret{}
    json.NewDecoder(r.Body).Decode(&secret)

	c.DBModel.Delete_Secret(secret.ID, c.DBModel.Get_User_By_UserName(claim.UserName).ID)
}
