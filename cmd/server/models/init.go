package models
// import packages
import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"

	"40104/SecretsManager/cmd/server/migrations"
)
// DB conncetion function
func (m *DBModel) ConnectDB(connection_string string) *sql.DB {
	db, err := sql.Open("postgres", connection_string) // Open DB conncetion
    if err != nil {
        log.Fatal(err) // Check error
    }
	// Ping DB
	if err = db.Ping(); err != nil {
		log.Fatal(err) // Check error
	}
	return db 
}
// DB migration function
func (m *DBModel) InitDB() {
	if _, err := m.DB.Query("select * from secrets;"); err != nil { //Check DB folder foes exist 
		if _, err := m.DB.Exec(migrations.Exec()); err != nil { //Start migration
			log.Fatal(err) // Check error
		} else {
			//Init DB by base data
			m.Add_Role(Role{Name:"Administrator"})
			m.Add_Role(Role{Name:"User"})
			m.Add_User(User{UserName:"admin", Password:"admin", Role_ID: m.Get_Role_By_Name("Administrator").ID,})
			m.Add_User(User{UserName:"user", Password:"user", Role_ID: m.Get_Role_By_Name("User").ID,})
			m.Add_Folder(Folder{Name:"root", Owner_ID: m.Get_User_By_UserName("admin").ID, Parrent_Folder_ID: sql.NullInt64{}})
			m.Add_Folder(Folder{Name:"myfolder", Owner_ID: m.Get_User_By_UserName("user").ID, Parrent_Folder_ID: sql.NullInt64{int64(m.Get_Folder_By_Name("root").ID),true}})
			m.Add_Secret(Secret{Name:"secret1",Username:"secret1",Secret:"secret1",Link:"secret1",Description:"secret1", Owner_ID: m.Get_User_By_UserName("user").ID, Folder_ID: m.Get_Folder_By_Name("root").ID,})
			m.Add_Secret(Secret{Name:"secret2",Username:"secret2",Secret:"secret2",Link:"secret2",Description:"secret2", Owner_ID: m.Get_User_By_UserName("user").ID, Folder_ID: m.Get_Folder_By_Name("myfolder").ID,})
		}
    }
}

