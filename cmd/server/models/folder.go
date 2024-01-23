package models
// import packages
import (
	"log"
)
// Create new Folder function
func (m *DBModel) Add_Folder(folder Folder) {
	if _, err := m.DB.Exec("insert into folders (name, owner_id, parrent_folder_id) values ($1, $2, $3)", 
		folder.Name, folder.Owner_ID, folder.Parrent_Folder_ID); err != nil { // Create new Folder 
			log.Fatal(err) //check error
	}	
}
// Get Folder by id function
func (m *DBModel) Get_Folder(id int) (Folder){
	folder := Folder{}
	row := m.DB.QueryRow("select * from folders where id = $1 ;", id) // Get Folder by id 
	if err := row.Scan(&folder.ID, &folder.Name, &folder.Owner_ID, &folder.Parrent_Folder_ID); err != nil { // Parse data to the Folder structure
		log.Fatal(err) //check error
	}
	
	folder.Owner = &User{UserName: m.Get_User(folder.Owner_ID).UserName} //Get owner username
	if folder.Parrent_Folder_ID.Valid {
		folder.Parrent_Folder = &Folder{Name: m.Get_Folder(int(folder.Parrent_Folder_ID.Int64)).Name} //Get parrent folder name if it exist
	}

	return folder
}
// Get Folder by name function
func (m *DBModel) Get_Folder_By_Name(name string) (Folder){
	folder := Folder{}
	row := m.DB.QueryRow("select * from folders where name = $1", name) // Get Folder by name
	if err := row.Scan(&folder.ID, &folder.Name, &folder.Owner_ID, &folder.Parrent_Folder_ID); err != nil { // Parse data to the Folder structure
		log.Fatal(err) //check error
	}
	folder.Owner = &User{UserName: m.Get_User(folder.Owner_ID).UserName} //Get owner username
	if folder.Parrent_Folder_ID.Valid {
		folder.Parrent_Folder = &Folder{Name: m.Get_Folder(int(folder.Parrent_Folder_ID.Int64)).Name}  //Get parrent folder name if it exist
	}

	return folder
}
// Get Folders by owner and parrent folder function
func (m *DBModel) Get_Folders_By_Owner_and_Parrent_Folder(owner_id int, parrent_folder_id int) ([]Folder){
	rows, err := m.DB.Query("select * from folders where owner_id = $1 and parrent_folder_id = $2", owner_id, parrent_folder_id) // Get Folders by owner and parrent folder 
	
	if err != nil {
        log.Fatal(err) //check error
    }

	folders:= []Folder{} //Init new slice

	for rows.Next(){
		folder:= Folder{}

		if err = rows.Scan(&folder.ID, &folder.Name, &folder.Owner_ID, &folder.Parrent_Folder_ID); err != nil{ // Parse data to the Folder structure
				log.Fatal(err) //check error
				continue
		}
		
		folder.Owner = &User{UserName: m.Get_User(folder.Owner_ID).UserName} //Get owner username
		if folder.Parrent_Folder_ID.Valid {
			folder.Parrent_Folder = &Folder{Name: m.Get_Folder(int(folder.Parrent_Folder_ID.Int64)).Name} //Get parrent folder name if it exist
		}

		folders = append(folders, folder) //Append folder to the slice
	}

	return folders
}
// Delete Folder function
func (m *DBModel) Delete_Folder(id int, owner_id int) {
	if _, err := m.DB.Exec("delete from folders where id = $1 and owner_id = $2;", id, owner_id); err != nil { // Delete Folder
		log.Fatal(err) //check error
	}
}
// Edit existing Folder function
func (m *DBModel) Put_Folder(folder Folder) {
	if _, err := m.DB.Exec("update folders set name = $1, owner_id = $2, parrent_folder_id = $3  where id = $4", // Edit existing Folder 
		 folder.Name, folder.Owner_ID, folder.Parrent_Folder_ID, folder.ID); err != nil {
			log.Fatal(err) //check error
    }
}