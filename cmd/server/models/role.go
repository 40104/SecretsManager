package models
// import packages
import (
	"log"
)
// Create new Role function
func (m *DBModel) Add_Role(role Role) {
	if _, err := m.DB.Exec("insert into roles (name) values ($1)", role.Name); err != nil { //Add new role
			log.Fatal(err) // Check error
    }
}
// Get Role by id function
func (m *DBModel) Get_Role(id int) Role{
	role := Role{}
	row := m.DB.QueryRow("select * from roles where id = $1", id) // Get Role by id 
	if err := row.Scan(&role.ID, &role.Name); err != nil { // Parse request to the Role structure
		log.Fatal(err) // Check error
	}
	
	return role
}
// Get Role by name function
func (m *DBModel) Get_Role_By_Name(name string) Role{
	role := Role{}
	row := m.DB.QueryRow("select * from roles where name = $1", name) // Get Role by name 
	if err := row.Scan(&role.ID, &role.Name); err != nil { // Parse request to the Role structure
		log.Fatal(err) // Check error
	}
	
	return role
}
// Delete Role function
func (m *DBModel) Delete_Role(id int) {
	if _, err := m.DB.Exec("delete from roles where id = $1", id); err != nil {  // Delete Role by id
		log.Fatal(err) // Check error
	}
}
// Edit existing Role function
func (m *DBModel) Put_Role(role Role) {
	if _, err := m.DB.Exec("update roles set name = $1 where id = $2", role.Name, role.ID); err != nil { // Edit existing Role 
			log.Fatal(err) // Check error
    }
}
