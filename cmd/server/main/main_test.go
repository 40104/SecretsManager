package main
// import packages
import (
	"testing"
	"database/sql"

	"40104/SecretsManager/cmd/server/models"
	//"40104/SecretsManager/cmd/server/controllers"
)

var token string //Global variable

// Create token test
func Test_Create_Token(t *testing.T) {
	// Init User test structure
    type Test struct {
		username string
		role string
	}
	test := Test{"user", "User"} // Init test data
	app := Application{} //Init application class
	app.Init("../configs/app.env") //Use env variables	

	token = app.Controller.CreateToken(test.username, test.role) // Create token
	if token == ""{
		t.Errorf("Error in token generating.")
	}
}

// Verify token test
func Test_Verify_Token(t *testing.T) {
	app := Application{} //Init application class
	app.Init("../configs/app.env") //Use env variables	

	if err, _ := app.Controller.VerifyToken(token); err != nil{ //Verify token	
		t.Errorf("Error in token verifing.")
	}
}

// Role test
func Test_Roles(t *testing.T) {
	// Init Role test structure
    var tests = []struct {
		//Params
        id int
		expected string
    }{
		//Values
        {1, "Administrator"},
        {2, "User"},
    }

	app := Application{} //Init application class
	app.Init("../configs/app.env") //Use env variables	
	//Read test
    for _, test := range tests{
        if output := app.Controller.DBModel.Get_Role(test.id); output.Name != test.expected { // transaction check 
            t.Errorf("Output %q not equal to expected %q", output.Name, test.expected)
        }
    }
	//Add test
	role := &models.Role{
		Name: "test", // New test Role
	}
	app.Controller.DBModel.Add_Role(*role) //Create new role 
	if output := app.Controller.DBModel.Get_Role_By_Name(role.Name); output.Name != role.Name { // transaction check 
		t.Errorf("Output %q not equal to expected %q", output.Name, role.Name ) 
	}
	
	// Put test
	role.ID = app.Controller.DBModel.Get_Role_By_Name(role.Name).ID // Get ID
	role.Name = "test2" // Edit Name
	app.Controller.DBModel.Put_Role(*role) // Edit role in the DB
	if output := app.Controller.DBModel.Get_Role_By_Name(role.Name); output.Name != role.Name { // transaction check 
		t.Errorf("Output %q not equal to expected %q", output.Name, role.Name )
	}
	//Delete test
	app.Controller.DBModel.Delete_Role(role.ID) // Delete role
	
}

// User test
func Test_User(t *testing.T) {
	// Init User test structure
    var tests = []struct {
		//Params
        id int
		expected string
    }{
		//Values
        {1, "admin"},
        {2, "user"},
    }

	app := Application{} //Init application class
	app.Init("../configs/app.env") //Use env variables	
	//Read test
    for _, test := range tests{
        if output := app.Controller.DBModel.Get_User(test.id); output.UserName != test.expected { // transaction check 
            t.Errorf("Output %q not equal to expected %q", output.UserName, test.expected)
        }
    }
	//Add test
	user := &models.User{
		UserName: "test", // New test User
		Password: "test",
		Role_ID: 2,
	}
	app.Controller.DBModel.Add_User(*user) //Create new user 
	if output := app.Controller.DBModel.Get_User_By_UserName(user.UserName); output.UserName != user.UserName { // transaction check 
		t.Errorf("Output %q not equal to expected %q", output.UserName, user.UserName ) 
	}
	
	// Put test
	user.ID = app.Controller.DBModel.Get_User_By_UserName(user.UserName).ID // Get ID
	user.UserName = "test2" // Edit Name
	app.Controller.DBModel.Put_User(*user) // Edit user in the DB
	if output := app.Controller.DBModel.Get_User_By_UserName(user.UserName); output.UserName != user.UserName { // transaction check 
		t.Errorf("Output %q not equal to expected %q", output.UserName, user.UserName )
	}
	//Delete test
	app.Controller.DBModel.Delete_User(user.ID) // Delete user
	
}

// Folder test
func Test_Folder(t *testing.T) {
	// Init Folder test structure
    var tests = []struct {
		//Params
        id int
		expected string
    }{
		//Values
        {1, "root"},
        {2, "myfolder"},
    }

	app := Application{} //Init application class
	app.Init("../configs/app.env") //Use env variables	
	//Read test
    for _, test := range tests{
        if output := app.Controller.DBModel.Get_Folder(test.id); output.Name != test.expected {
            t.Errorf("Output %q not equal to expected %q", output.Name, test.expected)
        }
    }
	//Add test
	folder := &models.Folder{
		Name: "test", // New test Folder
		Owner_ID: 2,
		Parrent_Folder_ID: sql.NullInt64{int64(2),true},
	}
	app.Controller.DBModel.Add_Folder(*folder) //Create new folder 
	if output := app.Controller.DBModel.Get_Folder_By_Name(folder.Name); output.Name != folder.Name { // transaction check 
		t.Errorf("Output %q not equal to expected %q", output.Name, folder.Name ) 
	}
	
	// Put test
	folder.ID = app.Controller.DBModel.Get_Folder_By_Name(folder.Name).ID // Get ID
	folder.Name = "test2" // Edit Name
	app.Controller.DBModel.Put_Folder(*folder) // Edit folder in the DB
	if output := app.Controller.DBModel.Get_Folder_By_Name(folder.Name); output.Name != folder.Name { // transaction check 
		t.Errorf("Output %q not equal to expected %q", output.Name, folder.Name ) 
	}
	//Delete test
	app.Controller.DBModel.Delete_Folder(folder.ID, folder.Owner_ID) // Delete folder
	
}

// Secret test
func Test_Secret(t *testing.T) {
	// Init Secret test structure
    var tests = []struct {
		//Params
        id int
		owner_id int
		expected string
    }{
		//Values
        {1, 2, "secret1"},
        {2, 2, "secret2"},
    }

	app := Application{} //Init application class
	app.Init("../configs/app.env") //Use env variables	
	//Read test
    for _, test := range tests{
        if output := app.Controller.DBModel.Get_Secret(test.id, test.owner_id); output.Name != test.expected { // transaction check 
            t.Errorf("Output %q not equal to expected %q", output.Name, test.expected)
        }
    }
	//Add test
	secret := &models.Secret{
		Name: "test", // New test Secret
		Username: "test",
		Secret: "test",
		Link: "test",
		Description: "test",
		Owner_ID: 2,
		Folder_ID: 2,
	}
	app.Controller.DBModel.Add_Secret(*secret) //Create new secret 

	secrets := app.Controller.DBModel.Get_Secrets_By_Owner_and_Folder(secret.Owner_ID,secret.Folder_ID) // Get list of secrets
	if len(secrets) > 0 {
        secret.ID = secrets[len(secrets)-1].ID // Get ID
    } else {
        t.Errorf("Error in reading list of secrets.") 
    }

	if output := app.Controller.DBModel.Get_Secret(secret.ID, secret.Owner_ID); output.Name != secret.Name { // transaction check 
		t.Errorf("Output %q not equal to expected %q", output.Name, secret.Name ) 
	}
	
	// Put test
	secret.Name = "test2" // Edit Name
	app.Controller.DBModel.Put_Secret(*secret) // Edit secret in the DB
	if output := app.Controller.DBModel.Get_Secret(secret.ID, secret.Owner_ID); output.Name != secret.Name { // transaction check 
		t.Errorf("Output %q not equal to expected %q", output.Name, secret.Name ) 
	}
	//Delete test
	app.Controller.DBModel.Delete_Secret(secret.ID, secret.Owner_ID) // Delete secret
	
}
