package models

import (
	"fmt"
	"encoding/json"
)

func (env *Env) Get_Role(role_id int) {	
	request_url := fmt.Sprintf("http://%s/role/get", env.Host)
	json_string := fmt.Sprintf(`{"id": %d}`, role_id)
	role := Role{}
	json.Unmarshal(env.Connect(request_url,json_string), &role)
	fmt.Printf("| ID: %d | Name: %s |\n", role.ID, role.Name)
}

func (env *Env) Add_Role(role_name string) {	
	request_url := fmt.Sprintf("http://%s/role/add", env.Host)
	json_string := fmt.Sprintf(`{"name": %s}`, role_name)
	env.Connect(request_url,string(json_string))
	fmt.Printf("New role successfully created. \n")
}

func (env *Env) Put_Role(role_id int, role_name string) {	
	request_url := fmt.Sprintf("http://%s/role/put", env.Host)
	json_string := fmt.Sprintf(`{"id": %d, "name": %s}`, role_id, role_name)
	env.Connect(request_url,string(json_string))
	fmt.Printf("Role successfully changed. \n")
}

func (env *Env) Delete_Role(role_id int) {	
	request_url := fmt.Sprintf("http://%s/role/delete", env.Host)
	json_string := fmt.Sprintf(`{"id": %d}`, role_id)
	env.Connect(request_url,string(json_string))
	fmt.Printf("Role with id: %d successfully deleted. \n", role_id)
}
