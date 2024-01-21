package models

import (
	"fmt"
	"encoding/json"
)

func (env *Env) Get_Secret(secret_id int) {	
	request_url := fmt.Sprintf("http://%s/secret/get", env.Host)
	json_string := fmt.Sprintf(`{"id": %d}`, secret_id)
	secret := Secret{}
	json.Unmarshal(env.Connect(request_url,json_string), &secret)
	fmt.Printf("| ID: %d |  Name: %s | Username: %s | Secret: %s | Link: %s | Description: %s | Owner: %s | Folder: %s | \n",
	 secret.ID, secret.Name, secret.Username, secret.Secret, secret.Link, secret.Description, secret.Owner.UserName, secret.Folder.Name )
}

func (env *Env) Add_Secret(name string, username string, password string, link string, description string, folder_name string) {	
	request_url := fmt.Sprintf("http://%s/secret/add", env.Host)
	secret := &Secret{
		Name: name,
		Username: username,
		Secret: password,
		Link: link,
		Description: description,
		Folder: &Folder{
			Name: folder_name,
		},
	}
	json_string, _ := json.Marshal(secret)
	env.Connect(request_url,string(json_string))
	fmt.Printf("New secret successfully created. \n")
	
}

func (env *Env) Put_Secret(secret_id int, name string, username string, password string, link string, description string, folder_name string) {	
	request_url := fmt.Sprintf("http://%s/secret/put", env.Host)
	secret := &Secret{
		ID: secret_id,
		Name: name,
		Username: username,
		Secret: password,
		Link: link,
		Description: description,
		Folder: &Folder{
			Name: folder_name,
		},
	}
	json_string, _ := json.Marshal(secret)
	env.Connect(request_url,string(json_string))
	fmt.Printf("Secret successfully changed. \n")
}

func (env *Env) Delete_Secret(secret_id int) {	
	request_url := fmt.Sprintf("http://%s/secret/delete", env.Host)
	json_string := fmt.Sprintf(`{"id": %d}`, secret_id)
	env.Connect(request_url,json_string)
	fmt.Printf("Secret with id: %d successfully deleted. \n", secret_id)
}