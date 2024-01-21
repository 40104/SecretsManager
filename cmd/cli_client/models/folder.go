package models

import (
	"fmt"
	"encoding/json"
)

func (env *Env) Get_Folder(folder_id int) {	
	request_url := fmt.Sprintf("http://%s/folder/get", env.Host)
	json_string := fmt.Sprintf(`{"id": %d}`, folder_id)
	folder := Folder{}
	json.Unmarshal(env.Connect(request_url,json_string), &folder)
	fmt.Printf("| ID: %d | Name: %s | Owner: %s | Parrent folder: %s |\n", folder.ID, folder.Name, folder.Owner.UserName, folder.Parrent_Folder.Name)
}

func (env *Env) Add_Folder(name string, parrent_folder string) {	
	request_url := fmt.Sprintf("http://%s/folder/add", env.Host)
	folder := &Folder{
		Name: name,
		Parrent_Folder: &Folder{
			Name: parrent_folder,
		},
	}
	json_string, _ := json.Marshal(folder)
	env.Connect(request_url,string(json_string))
	fmt.Printf("New folder successfully created. \n")
	
}

func (env *Env) Put_Folder(folder_id int, name string, parrent_folder string) {	
	request_url := fmt.Sprintf("http://%s/folder/put", env.Host)
	folder := &Folder{
		ID: folder_id,
		Name: name,
		Parrent_Folder: &Folder{
			Name: parrent_folder,
		},
	}
	json_string, _ := json.Marshal(folder)
	env.Connect(request_url,string(json_string))
	fmt.Printf("Folder successfully changed. \n")
}

func (env *Env) Delete_Folder(folder_id int) {	
	request_url := fmt.Sprintf("http://%s/folder/delete", env.Host)
	json_string := fmt.Sprintf(`{"id": %d}`, folder_id)
	env.Connect(request_url,json_string)
	fmt.Printf("Folder with id: %d successfully deleted. \n", folder_id)
}

func (env *Env) Read_Folder(parrent_folder string) {	
	request_url := fmt.Sprintf("http://%s/secret/getall", env.Host)
	secret := &Secret{
		Folder: &Folder{
			Name: parrent_folder,
		},
	}
	json_string, _ := json.Marshal(secret)

	secrets := []Secret{}
	json.Unmarshal(env.Connect(request_url,string(json_string)), &secrets)
	fmt.Printf(" Secrets: \n")
	for _, val := range secrets {
        fmt.Printf(" ID: %d  Name: %s Username: %s Secret: %s Link: %s Description: %s Owner: %s  Parrent folder: %s \n \n", 
			val.ID, val.Name, val.Username, val.Secret, val.Link, val.Description, val.Owner.UserName, val.Folder.Name )
    }
	
	request_url = fmt.Sprintf("http://%s/folder/list", env.Host)
	folder := &Folder{
		Parrent_Folder: &Folder{
			Name: parrent_folder,
		},
	}
	json_string, _ = json.Marshal(folder)
	
	folders := []Folder{}
	json.Unmarshal(env.Connect(request_url,string(json_string)), &folders)
	fmt.Printf("Folders: \n")
	for _, val := range folders {
        fmt.Printf(" ID: %d  Name: %s  Owner: %s  Parrent folder: %s \n", val.ID, val.Name, val.Owner.UserName, val.Parrent_Folder.Name)
    }
}


