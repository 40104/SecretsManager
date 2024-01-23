package models

import (
	"fmt"
	"encoding/json"
)

func (env *Env) Get_User(user_id int) {	
	request_url := fmt.Sprintf("http://%s/user/get", env.Host)
	json_string := fmt.Sprintf(`{"id": %d}`, user_id)
	user := User{}
	json.Unmarshal(env.Connect(request_url,json_string), &user)
	fmt.Printf("| ID: %d | UserName: %s | Password: %s | Role: %s |\n", user.ID, user.UserName, user.Password, user.Role.Name)
}

func (env *Env) Add_User(username string, password string, role_name string) {	
	request_url := fmt.Sprintf("http://%s/user/add", env.Host)
	user := &User{
		UserName: username,
		Password: password,
		Role: &Role{
			Name: role_name,
		},
	}
	json_string, _ := json.Marshal(user)
	env.Connect(request_url,string(json_string))
	fmt.Printf("New user successfully created. \n")
}

func (env *Env) Put_User(user_id int,username string, password string, role_name string) {	
	request_url := fmt.Sprintf("http://%s/user/put", env.Host)
	user := &User{
		ID: user_id,
		UserName: username,
		Password: password,
		Role: &Role{
			Name: role_name,
		},
	}
	json_string, _ := json.Marshal(user)
	env.Connect(request_url,string(json_string))
	fmt.Printf("User successfully changed. \n")
}

func (env *Env) Delete_User(user_id int) {	
	request_url := fmt.Sprintf("http://%s/user/delete", env.Host)
	json_string := fmt.Sprintf(`{"id": %d}`, user_id)
	env.Connect(request_url,string(json_string))
	fmt.Printf("User with id: %d successfully deleted. \n", user_id)
}