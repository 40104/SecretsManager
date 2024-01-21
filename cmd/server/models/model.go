package models

import (
	"database/sql"
)

type DBModel struct {
	DB *sql.DB
	Key string
}

type Role struct {
	ID int `json:"id"`
	Name string `json:"name"`
}
type User struct {
	ID int `json:"id"`
	UserName string `json:"username"`
	Password string `json:"password"`
	Role_ID int `json:"role_id"`
	Role *Role `json:"role"`
}
type Folder struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Owner_ID int `json:"owner_id"`
	Owner *User `json:"owner"`
	Parrent_Folder_ID sql.NullInt64 `json:"parrent_folder_id"`
	Parrent_Folder *Folder `json:"parrent_folder"`
}
type Secret struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Username string `json:"username"`
	Secret string `json:"secret"`
	Link string `json:"link"`
	Description string `json:"description"`
	Owner_ID int `json:"owner_id"`
	Owner *User `json:"owner"`
	Folder_ID int `json:"folder_id"`
	Folder *Folder `json:"folder"`
}


