package controllers

import (
	"net/http"
	"encoding/json"

	"40104/SecretsManager/cmd/server/models"
)

func (c *Controller) Get_Role(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	req_role := models.Role{}
    json.NewDecoder(r.Body).Decode(&req_role)

	role := c.DBModel.Get_Role(req_role.ID)
	json.NewEncoder(w).Encode(role)
}

func (c *Controller) Add_Role(w http.ResponseWriter, r *http.Request) {
	role := models.Role{}
    json.NewDecoder(r.Body).Decode(&role)
	
	c.DBModel.Add_Role(role)
}

func (c *Controller) Put_Role(w http.ResponseWriter, r *http.Request) {
	role := models.Role{}
    json.NewDecoder(r.Body).Decode(&role)

	c.DBModel.Put_Role(role)
}

func (c *Controller) Delete_Role(w http.ResponseWriter, r *http.Request) {
	role := models.Role{}
    json.NewDecoder(r.Body).Decode(&role)

	c.DBModel.Delete_Role(role.ID)
}