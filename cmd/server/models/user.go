package models
// import packages
import (
	"log"
	"golang.org/x/crypto/bcrypt"
)
// Hashing function
func (m *DBModel) HashPassword(password string) (string) {
    bytes, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost) // Hashing 
    return string(bytes)
}
// Create new User function
func (m *DBModel) Add_User(user User) {
	bytes:= m.HashPassword(user.Password) // Hashing password
	if _, err := m.DB.Exec("insert into users (username, password, role_id) values ($1, $2, $3)", // Create new User
		user.UserName, bytes, user.Role_ID); err != nil {
			log.Fatal(err) //check error
    }
}
// Get User by id function
func (m *DBModel) Get_User(id int) (User){
	user := User{}
	row := m.DB.QueryRow("select * from users where id = $1", id) // Get User by id
	if err := row.Scan(&user.ID, &user.UserName, &user.Password, &user.Role_ID); err != nil { // Parse user to the structure
		log.Fatal(err) //check error
	}
	user.Role = &Role{Name: m.Get_Role(user.Role_ID).Name} // Get role name

	return user
}
// Get User by username function
func (m *DBModel) Get_User_By_UserName(username string) (User){
	user := User{}
	row := m.DB.QueryRow("select * from users where username = $1", username) // Get User by username 
	if err := row.Scan(&user.ID, &user.UserName, &user.Password, &user.Role_ID); err != nil { // Parse user to the structure
		log.Fatal(err) //check error
	}
	user.Role = &Role{Name: m.Get_Role(user.Role_ID).Name}  // Get role name
	
	return user
}
// Delete User function
func (m *DBModel) Delete_User(id int) {
	if _, err := m.DB.Exec("delete from users where id = $1", id); err != nil { // Delete User
		log.Fatal(err) //check error
	}
}
// Edit User function
func (m *DBModel) Put_User(user User) {
	if _, err := m.DB.Exec("update users set username = $1, password = $2, role_id = $3  where id = $4",
		 user.UserName, m.Encrypt(user.Password), user.Role_ID, user.ID); err != nil { // Edit User
			log.Fatal(err) //check error
    }
}