package models

import (
	"log"
)

func (m *DBModel) Add_Secret(secret Secret) {
	if _, err := m.DB.Exec("insert into secrets (name, username, secret, link, description, owner_id, folder_id) values ($1, $2, $3, $4, $5, $6, $7);", 
		m.Encrypt(secret.Name), m.Encrypt(secret.Username), m.Encrypt(secret.Secret), 
		m.Encrypt(secret.Link), m.Encrypt(secret.Description), secret.Owner_ID, secret.Folder_ID); err != nil {
			log.Fatal(err)
    }
}

func (m *DBModel) Get_Secret(id int, owner_id int) (Secret) {
	secret := Secret{}
	row := m.DB.QueryRow("select * from secrets where id = $1 and owner_id = $2;", id, owner_id)
	
	row.Scan(&secret.ID, &secret.Name, &secret.Username, &secret.Secret, &secret.Link, &secret.Description, &secret.Owner_ID, &secret.Folder_ID)

	secret.Name=m.Decrypt(secret.Name)
	secret.Username=m.Decrypt(secret.Username)
	secret.Secret=m.Decrypt(secret.Secret)
	secret.Link=m.Decrypt(secret.Link)
	secret.Description=m.Decrypt(secret.Description)

	secret.Owner = &User{UserName: m.Get_User(secret.Owner_ID).UserName}
	secret.Folder = &Folder{Name: m.Get_Folder(secret.Folder_ID).Name}

	return secret
}


func (m *DBModel) Get_Secrets_By_Owner_and_Folder(owner_id int, folder_id int) []Secret {
	rows, err := m.DB.Query("select * from secrets where owner_id = $1 and folder_id = $2;", owner_id, folder_id)
	if err != nil {
        log.Fatal(err)
    }
	secrets:= []Secret{}

	for rows.Next(){
		secret:= Secret{}

		if err = rows.Scan(&secret.ID, &secret.Name, &secret.Username, &secret.Secret, &secret.Link, &secret.Description, &secret.Owner_ID, &secret.Folder_ID); err != nil{
				log.Fatal(err)
				continue
		}

		secret.Name=m.Decrypt(secret.Name)
		secret.Username=m.Decrypt(secret.Username)
		secret.Link=m.Decrypt(secret.Link)
		secret.Description=m.Decrypt(secret.Description)

		secret.Owner = &User{UserName: m.Get_User(secret.Owner_ID).UserName}
		secret.Folder = &Folder{Name: m.Get_Folder(secret.Folder_ID).Name}
		
		secrets = append(secrets, secret)
	}

	return secrets
}

func (m *DBModel) Delete_Secret(id int, owner_id int) {
	if _, err := m.DB.Exec("delete from secrets where id = $1 and owner_id = $2;", id, owner_id); err != nil {
		log.Fatal(err)
	}
}

func (m *DBModel) Put_Secret(secret Secret) {
	if _, err := m.DB.Exec("update secrets set name = $1, username = $2, secret = $3, link = $4, description = $5, owner_id = $6, folder_id = $7 where id = $8;", 
		m.Encrypt(secret.Name), m.Encrypt(secret.Username), m.Encrypt(secret.Secret),
		 m.Encrypt(secret.Link), m.Encrypt(secret.Description),secret.Owner_ID, secret.Folder_ID, secret.ID); err != nil {
			log.Fatal(err)
    }
}


