package models

import (
	"log"
	"golang.org/x/crypto/bcrypt"
)

func (m *DBModel) HashPassword(password string) (string) {
    bytes, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    return string(bytes)
}

func (m *DBModel) Add_User(user User) {
	bytes:= m.HashPassword(user.Password)
	if _, err := m.DB.Exec("insert into users (username, password, role_id) values ($1, $2, $3)", 
		user.UserName, bytes, user.Role_ID); err != nil {
			log.Fatal(err)
    }
}

func (m *DBModel) Get_User(id int) (User){
	user := User{}
	row := m.DB.QueryRow("select * from users where id = $1", id)
	if err := row.Scan(&user.ID, &user.UserName, &user.Password, &user.Role_ID); err != nil {
		log.Fatal(err)
	}
	user.Role = &Role{Name: m.Get_Role(user.Role_ID).Name}

	return user
}

func (m *DBModel) Get_User_By_UserName(username string) (User){
	user := User{}
	row := m.DB.QueryRow("select * from users where username = $1", username)
	if err := row.Scan(&user.ID, &user.UserName, &user.Password, &user.Role_ID); err != nil {
		log.Fatal(err)
	}
	user.Role = &Role{Name: m.Get_Role(user.Role_ID).Name}
	
	return user
}

func (m *DBModel) Delete_User(id int) {
	if _, err := m.DB.Exec("delete from users where id = $1", id); err != nil {
		log.Fatal(err)
	}
}

func (m *DBModel) Put_User(user User) {
	if _, err := m.DB.Exec("update users set username = $1, password = $2, role_id = $3  where id = $4",
		 user.UserName, m.Encrypt(user.Password), user.Role_ID, user.ID); err != nil {
			log.Fatal(err)
    }
}