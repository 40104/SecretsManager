package models

import (
	"log"
)

func (m *DBModel) Add_Folder(folder Folder) {
	if _, err := m.DB.Exec("insert into folders (name, owner_id, parrent_folder_id) values ($1, $2, $3)", 
		folder.Name, folder.Owner_ID, folder.Parrent_Folder_ID); err != nil {
			log.Fatal(err)
	}	
}

func (m *DBModel) Get_Folder(id int) (Folder){
	folder := Folder{}
	row := m.DB.QueryRow("select * from folders where id = $1 ;", id)
	if err := row.Scan(&folder.ID, &folder.Name, &folder.Owner_ID, &folder.Parrent_Folder_ID); err != nil {
		log.Fatal(err)
	}
	
	folder.Owner = &User{UserName: m.Get_User(folder.Owner_ID).UserName}
	if folder.Parrent_Folder_ID.Valid {
		folder.Parrent_Folder = &Folder{Name: m.Get_Folder(int(folder.Parrent_Folder_ID.Int64)).Name}
	}

	return folder
}

func (m *DBModel) Get_Folder_By_Name(name string) (Folder){
	folder := Folder{}
	row := m.DB.QueryRow("select * from folders where name = $1", name)
	if err := row.Scan(&folder.ID, &folder.Name, &folder.Owner_ID, &folder.Parrent_Folder_ID); err != nil {
		log.Fatal(err)
	}
	folder.Owner = &User{UserName: m.Get_User(folder.Owner_ID).UserName}
	if folder.Parrent_Folder_ID.Valid {
		folder.Parrent_Folder = &Folder{Name: m.Get_Folder(int(folder.Parrent_Folder_ID.Int64)).Name}
	}

	return folder
}

func (m *DBModel) Get_Folders_By_Owner_and_Parrent_Folder(owner_id int, parrent_folder_id int) ([]Folder){
	rows, err := m.DB.Query("select * from folders where owner_id = $1 and parrent_folder_id = $2", owner_id, parrent_folder_id)
	
	if err != nil {
        log.Fatal(err)
    }

	folders:= []Folder{}

	for rows.Next(){
		folder:= Folder{}

		if err = rows.Scan(&folder.ID, &folder.Name, &folder.Owner_ID, &folder.Parrent_Folder_ID); err != nil{
				log.Fatal(err)
				continue
		}
		
		folder.Owner = &User{UserName: m.Get_User(folder.Owner_ID).UserName}
		if folder.Parrent_Folder_ID.Valid {
			folder.Parrent_Folder = &Folder{Name: m.Get_Folder(int(folder.Parrent_Folder_ID.Int64)).Name}
		}

		folders = append(folders, folder)
	}

	return folders
}

func (m *DBModel) Delete_Folder(id int, owner_id int) {
	if _, err := m.DB.Exec("delete from folders where id = $1 and owner_id = $2;", id, owner_id); err != nil {
		log.Fatal(err)
	}
}

func (m *DBModel) Put_Folder(folder Folder) {
	if _, err := m.DB.Exec("update folders set name = $1, owner_id = $2, parrent_folder_id = $3  where id = $4",
		 folder.Name, folder.Owner_ID, folder.Parrent_Folder_ID, folder.ID); err != nil {
			log.Fatal(err)
    }
}