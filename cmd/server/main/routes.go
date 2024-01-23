package main
// Import packages
import (
	"net/http" 
)
// Setup routes
func (app *Application) Routes() *http.ServeMux {
    // Set new server
    mux := http.NewServeMux()
    // Set url paths
    mux.HandleFunc("/", app.Controller.Home)
    mux.HandleFunc("/login", app.Controller.Login)
    
    mux.Handle("/secret/get", app.Controller.UserAuthMiddleWare(http.HandlerFunc(app.Controller.Get_Secret)))
    mux.Handle("/secret/getall", app.Controller.UserAuthMiddleWare(http.HandlerFunc(app.Controller.Get_Secrets)))
    mux.Handle("/secret/put", app.Controller.UserAuthMiddleWare(http.HandlerFunc(app.Controller.Put_Secret)))
    mux.Handle("/secret/add", app.Controller.UserAuthMiddleWare(http.HandlerFunc(app.Controller.Add_Secret)))
    mux.Handle("/secret/delete", app.Controller.UserAuthMiddleWare(http.HandlerFunc(app.Controller.Delete_Secret)))
    mux.Handle("/secret/generate", app.Controller.UserAuthMiddleWare(http.HandlerFunc(app.Controller.Generate_Secret)))

    mux.Handle("/folder/get", app.Controller.UserAuthMiddleWare(http.HandlerFunc(app.Controller.Get_Folder)))
    mux.Handle("/folder/put", app.Controller.UserAuthMiddleWare(http.HandlerFunc(app.Controller.Put_Folder)))
    mux.Handle("/folder/add", app.Controller.UserAuthMiddleWare(http.HandlerFunc(app.Controller.Add_Folder)))
    mux.Handle("/folder/delete", app.Controller.UserAuthMiddleWare(http.HandlerFunc(app.Controller.Delete_Folder)))
    mux.Handle("/folder/list", app.Controller.UserAuthMiddleWare(http.HandlerFunc(app.Controller.Get_Folders)))

    mux.Handle("/role/get", app.Controller.AdminAuthMiddleWare(http.HandlerFunc(app.Controller.Get_Role)))
    mux.Handle("/role/put", app.Controller.AdminAuthMiddleWare(http.HandlerFunc(app.Controller.Put_Role)))
    mux.Handle("/role/add", app.Controller.AdminAuthMiddleWare(http.HandlerFunc(app.Controller.Add_Role)))
    mux.Handle("/role/delete", app.Controller.AdminAuthMiddleWare(http.HandlerFunc(app.Controller.Delete_Role)))

    mux.Handle("/user/get", app.Controller.AdminAuthMiddleWare(http.HandlerFunc(app.Controller.Get_User)))
    mux.Handle("/user/put", app.Controller.AdminAuthMiddleWare(http.HandlerFunc(app.Controller.Put_User)))
    mux.Handle("/user/add", app.Controller.AdminAuthMiddleWare(http.HandlerFunc(app.Controller.Add_User)))
    mux.Handle("/user/delete", app.Controller.AdminAuthMiddleWare(http.HandlerFunc(app.Controller.Delete_User)))

    // Set staticv folder path
    fileServer := http.FileServer(http.Dir("./views/static/"))
    mux.Handle("/static/", app.Controller.UserAuthMiddleWare(http.StripPrefix("/static", fileServer)))
    // Return new server
    return mux
}