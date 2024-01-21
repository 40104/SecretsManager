package controllers

import (
	"net/http"
	"encoding/json"

	"40104/SecretsManager/cmd/server/models"
)

func (c *Controller) Get_User(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	req_user := models.User{}
    json.NewDecoder(r.Body).Decode(&req_user)

	user:= c.DBModel.Get_User(req_user.ID)
	json.NewEncoder(w).Encode(user)
}

func (c *Controller) Add_User(w http.ResponseWriter, r *http.Request) {
	user := models.User{}
    json.NewDecoder(r.Body).Decode(&user)

	user.Role_ID = c.DBModel.Get_Role_By_Name(user.Role.Name).ID
	c.DBModel.Add_User(user)
}

func (c *Controller) Put_User(w http.ResponseWriter, r *http.Request) {
	user := models.User{}
    json.NewDecoder(r.Body).Decode(&user)

	user.Role_ID = c.DBModel.Get_Role_By_Name(user.Role.Name).ID
	c.DBModel.Put_User(user)
}

func (c *Controller) Delete_User(w http.ResponseWriter, r *http.Request) {
	user := models.User{}
    json.NewDecoder(r.Body).Decode(&user)

	c.DBModel.Delete_User(user.ID)
}