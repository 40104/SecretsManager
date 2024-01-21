package models

import (
	"log"
)

func (m *DBModel) Add_Role(role Role) {
	if _, err := m.DB.Exec("insert into roles (name) values ($1)", role.Name); err != nil {
			log.Fatal(err)
    }
}

func (m *DBModel) Get_Role(id int) Role{
	role := Role{}
	row := m.DB.QueryRow("select * from roles where id = $1", id)
	if err := row.Scan(&role.ID, &role.Name); err != nil {
		log.Fatal(err)
	}
	
	return role
}

func (m *DBModel) Get_Role_By_Name(name string) Role{
	role := Role{}
	row := m.DB.QueryRow("select * from roles where name = $1", name)
	if err := row.Scan(&role.ID, &role.Name); err != nil {
		log.Fatal(err)
	}
	
	return role
}

func (m *DBModel) Delete_Role(id int) {
	if _, err := m.DB.Exec("delete from roles where id = $1", id); err != nil {
		log.Fatal(err)
	}
}

func (m *DBModel) Put_Role(role Role) {
	if _, err := m.DB.Exec("update roles set name = $1 where id = $2", role.Name, role.ID); err != nil {
			log.Fatal(err)
    }
}
