package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"syscall"

	"github.com/urfave/cli/v2"
	"github.com/joho/godotenv"
	"github.com/dchest/uniuri"
	
	"40104/SecretsManager/cmd/cli_client/models"
	
)

type Application struct{
	CLI *cli.App
	Env *models.Env
	Param *models.Params
}

// Import env function
func (app *Application) EnvVariable(file string) {
    if err := godotenv.Load(file); err != nil {
    	log.Fatal(err)
    }
	Base_Length, err := strconv.Atoi(os.Getenv("BASE_LENGTH"));
	if  err != nil {
        log.Fatal(err)
    
	}

	app.Env.Base_Length = Base_Length
	app.Env.Host = os.Getenv("HOST")
	app.Env.Username = os.Getenv("USERNAME")
	app.Env.Password = os.Getenv("PASSWORD")
}

func (app *Application) Setup() *cli.App{
	CLI := &cli.App{
        Commands: []*cli.Command{
            {
                Name:    "init",
                Aliases: []string{"i"},
                Usage:   "Init connection to the server",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:  "host",
						Usage: "Set host",
						Value: app.Env.Host,
						Destination: &app.Env.Host,
					},
					&cli.StringFlag{
						Name:  "username",
						Usage: "Set username",
						Value: app.Env.Username,
						Destination: &app.Env.Username,
					},
					&cli.StringFlag{
						Name:  "password",
						Usage: "Set password",
						Value: app.Env.Password,
						Destination: &app.Env.Password,
					},
				},
                Action: func(c *cli.Context) error {
					app.Init()
					return nil
				},
            },
			{
                Name:    "folder",
                Aliases: []string{"f"},
                Usage:   "Actions with folders",
                Subcommands: []*cli.Command{
					{
                        Name:  "read",
                        Usage: "Read folder data.",
						Aliases: []string{"r"},
						Flags: []cli.Flag{
							&cli.StringFlag{
								Name:  "folder",
								Usage: "Folder name",
								Aliases: []string{"f"},
								Value: "root",
								Destination: &app.Param.Folder,
								DefaultText: "root",
							},
						},
                        Action: func(cCtx *cli.Context) error {
                            app.Env.Read_Folder(app.Param.Folder)
                            return nil
                        },
                    },
					{
                        Name:  "get",
                        Usage: "Get folder",
						Aliases: []string{"g"},
						Flags: []cli.Flag{
							&cli.IntFlag{
								Name:  "id",
								Usage: "Folder id",
								Destination: &app.Param.ID,
								Required: true,
							},
						},
                        Action: func(cCtx *cli.Context) error {
                            app.Env.Get_Folder(app.Param.ID)
                            return nil
                        },
                    },
                    {
                        Name:  "add",
                        Usage: "Create new folder",
						Aliases: []string{"a"},
						Flags: []cli.Flag{
							&cli.StringFlag{
								Name:  "name",
								Usage: "Folder name",
								Aliases: []string{"n"},
								Destination: &app.Param.Folder,
								Required: true,
							},
							&cli.StringFlag{
								Name:  "parrent_folder",
								Usage: "Parrent folder name",
								Aliases: []string{"p"},
								Value: "root",
								Destination: &app.Param.Parrent_Folder,
								DefaultText: "root",
							},
						},
                        Action: func(cCtx *cli.Context) error {
                            app.Env.Add_Folder(app.Param.Folder, app.Param.Parrent_Folder)
                            return nil
                        },
                    },
                    {
                        Name:  "put",
                        Usage: "Update folder",
						Aliases: []string{"p"},
						Flags: []cli.Flag{
							&cli.IntFlag{
								Name:  "id",
								Usage: "Folder id",
								Destination: &app.Param.ID,
								Required: true,
							},
							&cli.StringFlag{
								Name:  "name",
								Usage: "Folder name",
								Aliases: []string{"n"},
								Destination: &app.Param.Folder,
								Required: true,
							},
							&cli.StringFlag{
								Name:  "parrent_folder",
								Usage: "Parrent folder name",
								Aliases: []string{"p"},
								Value: "root",
								Destination: &app.Param.Parrent_Folder,
								DefaultText: "root",
							},
						},
                        Action: func(cCtx *cli.Context) error {
                            app.Env.Put_Folder(app.Param.ID,app.Param.Folder, app.Param.Parrent_Folder)
                            return nil
                        },
                    },
					{
                        Name:  "delete",
                        Usage: "Delete folder",
						Aliases: []string{"d"},
						Flags: []cli.Flag{
							&cli.IntFlag{
								Name:  "id",
								Usage: "Folder id",
								Destination: &app.Param.ID,
								Required: true,
							},
						},
                        Action: func(cCtx *cli.Context) error {
                            app.Env.Delete_Folder(app.Param.ID)
                            return nil
                        },
                    },
                },
            },
			{
                Name:    "secret",
                Aliases: []string{"s"},
                Usage:   "Actions with secrets",
                Subcommands: []*cli.Command{
					{
                        Name:  "get",
                        Usage: "Get secret",
						Aliases: []string{"g"},
						Flags: []cli.Flag{
							&cli.IntFlag{
								Name:  "id",
								Usage: "Secret id",
								Destination: &app.Param.ID,
								Required: true,
							},
						},
                        Action: func(cCtx *cli.Context) error {
                            app.Env.Get_Secret(app.Param.ID)
                            return nil
                        },
                    },
                    {
                        Name:  "add",
                        Usage: "Create new secret",
						Aliases: []string{"a"},
						Flags: []cli.Flag{
							&cli.StringFlag{
								Name:  "name",
								Usage: "Set name",
								Aliases: []string{"n"},
								Destination: &app.Param.Name,
								Required: true,
							},
							&cli.StringFlag{
								Name:  "username",
								Usage: "Set Username",
								Aliases: []string{"u"},
								Destination: &app.Param.UserName,
								Required: true,
							},
							&cli.StringFlag{
								Name:  "secret",
								Usage: "Set Secret",
								Aliases: []string{"s"},
								Destination: &app.Param.Password,
								Required: true,
							},
							&cli.StringFlag{
								Name:  "link",
								Usage: "Set link",
								Aliases: []string{"l"},
								Destination: &app.Param.Link,
							},
							&cli.StringFlag{
								Name:  "description",
								Usage: "Set description",
								Aliases: []string{"d"},
								Destination: &app.Param.Description,
							},
							&cli.StringFlag{
								Name:  "folder",
								Usage: "Set folder",
								Aliases: []string{"f"},
								Destination: &app.Param.Folder,
								Value: "root",
								DefaultText: "root",
							},
						},
                        Action: func(cCtx *cli.Context) error {
                            app.Env.Add_Secret(app.Param.Name,app.Param.UserName, app.Param.Password, app.Param.Link, app.Param.Description, app.Param.Folder)
                            return nil
                        },
                    },
					{
                        Name:  "generate",
                        Usage: "Generate new secret",
						Aliases: []string{"gen"},
						Flags: []cli.Flag{
							&cli.StringFlag{
								Name:  "name",
								Usage: "Set name",
								Aliases: []string{"n"},
								Destination: &app.Param.Name,
								Required: true,
							},
							&cli.StringFlag{
								Name:  "username",
								Usage: "Set Username",
								Aliases: []string{"u"},
								Destination: &app.Param.UserName,
								Required: true,
							},
							&cli.StringFlag{
								Name:  "link",
								Usage: "Set link",
								Aliases: []string{"l"},
								Destination: &app.Param.Link,
							},
							&cli.StringFlag{
								Name:  "description",
								Usage: "Set description",
								Aliases: []string{"d"},
								Destination: &app.Param.Description,
							},
							&cli.StringFlag{
								Name:  "folder",
								Usage: "Set folder",
								Aliases: []string{"f"},
								Destination: &app.Param.Folder,
								Value: "root",
								DefaultText: "root",
							},
						},
                        Action: func(cCtx *cli.Context) error {
                            app.Env.Add_Secret(app.Param.Name,app.Param.UserName, uniuri.NewLen(app.Env.Base_Length), app.Param.Link, app.Param.Description, app.Param.Folder)
                            return nil
                        },
                    },
                    {
                        Name:  "put",
                        Usage: "Update secret",
						Aliases: []string{"p"},
						Flags: []cli.Flag{
							&cli.IntFlag{
								Name:  "id",
								Usage: "Set id",
								Destination: &app.Param.ID,
								Required: true,
							},
							&cli.StringFlag{
								Name:  "name",
								Usage: "Set name",
								Aliases: []string{"n"},
								Destination: &app.Param.Name,
								Required: true,
							},
							&cli.StringFlag{
								Name:  "username",
								Usage: "Set Username",
								Aliases: []string{"u"},
								Destination: &app.Param.UserName,
								Required: true,
							},
							&cli.StringFlag{
								Name:  "secret",
								Usage: "Set Secret",
								Aliases: []string{"s"},
								Destination: &app.Param.Password,
								Required: true,
							},
							&cli.StringFlag{
								Name:  "link",
								Usage: "Set link",
								Aliases: []string{"l"},
								Destination: &app.Param.Link,
							},
							&cli.StringFlag{
								Name:  "description",
								Usage: "Set description",
								Aliases: []string{"d"},
								Destination: &app.Param.Description,
							},
							&cli.StringFlag{
								Name:  "folder",
								Usage: "Set folder",
								Aliases: []string{"f"},
								Destination: &app.Param.Folder,
								Value: "root",
								DefaultText: "root",
							},
						},
                        Action: func(cCtx *cli.Context) error {
                            app.Env.Put_Secret(app.Param.ID, app.Param.Name,app.Param.UserName, app.Param.Password, app.Param.Link, app.Param.Description, app.Param.Folder)
                            return nil
                        },
                    },
					{
                        Name:  "delete",
                        Usage: "Delete secret",
						Aliases: []string{"d"},
						Flags: []cli.Flag{
							&cli.IntFlag{
								Name:  "id",
								Usage: "Secret id",
								Destination: &app.Param.ID,
								Required: true,
							},
						},
                        Action: func(cCtx *cli.Context) error {
                            app.Env.Delete_Secret(app.Param.ID)
                            return nil
                        },
                    },
                },
            },
			{
                Name:    "user",
                Aliases: []string{"u"},
                Usage:   "Actions with users",
                Subcommands: []*cli.Command{
					{
                        Name:  "get",
                        Usage: "Get user",
						Aliases: []string{"g"},
						Flags: []cli.Flag{
							&cli.IntFlag{
								Name:  "id",
								Usage: "User id",
								Destination: &app.Param.ID,
								Required: true,
							},
						},
                        Action: func(cCtx *cli.Context) error {
                            app.Env.Get_User(app.Param.ID)
                            return nil
                        },
                    },
                    {
                        Name:  "add",
                        Usage: "Create new user",
						Aliases: []string{"a"},
						Flags: []cli.Flag{
							&cli.StringFlag{
								Name:  "username",
								Usage: "Set Username",
								Aliases: []string{"u"},
								Destination: &app.Param.UserName,
								Required: true,
							},
							&cli.StringFlag{
								Name:  "password",
								Usage: "Set password",
								Aliases: []string{"p"},
								Destination: &app.Param.Password,
								Required: true,
							},
							&cli.StringFlag{
								Name:  "role",
								Usage: "Role name",
								Aliases: []string{"r"},
								Destination: &app.Param.Role_Name,
								Required: true,
								Value: "User",
								DefaultText: "User",
							},
						},
                        Action: func(cCtx *cli.Context) error {
                            app.Env.Add_User(app.Param.UserName, app.Param.Password, app.Param.Role_Name)
                            return nil
                        },
                    },
                    {
                        Name:  "put",
                        Usage: "Update user",
						Aliases: []string{"p"},
						Flags: []cli.Flag{
							&cli.IntFlag{
								Name:  "id",
								Usage: "User id",
								Destination: &app.Param.ID,
								Required: true,
							},
							&cli.StringFlag{
								Name:  "username",
								Usage: "Set Username",
								Aliases: []string{"u"},
								Destination: &app.Param.UserName,
								Required: true,
							},
							&cli.StringFlag{
								Name:  "password",
								Usage: "Set password",
								Aliases: []string{"p"},
								Destination: &app.Param.Password,
								Required: true,
							},
							&cli.StringFlag{
								Name:  "role",
								Usage: "Role name",
								Aliases: []string{"r"},
								Destination: &app.Param.Role_Name,
								Required: true,
								Value: "User",
								DefaultText: "User",
							},
						},
                        Action: func(cCtx *cli.Context) error {
                            app.Env.Put_User(app.Param.ID, app.Param.UserName, app.Param.Password, app.Param.Role_Name)
                            return nil
                        },
                    },
					{
                        Name:  "delete",
                        Usage: "Delete user",
						Aliases: []string{"d"},
						Flags: []cli.Flag{
							&cli.IntFlag{
								Name:  "id",
								Usage: "User id",
								Destination: &app.Param.ID,
								Required: true,
							},
						},
                        Action: func(cCtx *cli.Context) error {
                            app.Env.Delete_User(app.Param.ID)
                            return nil
                        },
                    },
                },
            },
			{
                Name:    "role",
                Aliases: []string{"r"},
                Usage:   "Actions with roles",
                Subcommands: []*cli.Command{
					{
                        Name:  "get",
                        Usage: "Get role",
						Aliases: []string{"g"},
						Flags: []cli.Flag{
							&cli.IntFlag{
								Name:  "id",
								Usage: "Role id",
								Destination: &app.Param.ID,
								Required: true,
							},
						},
                        Action: func(cCtx *cli.Context) error {
                            app.Env.Get_Role(app.Param.ID)
                            return nil
                        },
                    },
                    {
                        Name:  "add",
                        Usage: "Create new role",
						Aliases: []string{"a"},
						Flags: []cli.Flag{
							&cli.StringFlag{
								Name:  "name",
								Usage: "Role name",
								Aliases: []string{"n"},
								Destination: &app.Param.Name,
								Required: true,
							},
						},
                        Action: func(cCtx *cli.Context) error {
                            app.Env.Add_Role(app.Param.Name)
                            return nil
                        },
                    },
                    {
                        Name:  "put",
                        Usage: "Update role",
						Aliases: []string{"p"},
						Flags: []cli.Flag{
							&cli.IntFlag{
								Name:  "id",
								Usage: "Role id",
								Destination: &app.Param.ID,
								Required: true,
							},
							&cli.StringFlag{
								Name:  "name",
								Usage: "Role name",
								Aliases: []string{"n"},
								Destination: &app.Param.Folder,
								Required: true,
							},
						},
                        Action: func(cCtx *cli.Context) error {
                            app.Env.Put_Role(app.Param.ID,app.Param.Name)
                            return nil
                        },
                    },
					{
                        Name:  "delete",
                        Usage: "Delete role",
						Aliases: []string{"d"},
						Flags: []cli.Flag{
							&cli.IntFlag{
								Name:  "id",
								Usage: "Role id",
								Destination: &app.Param.ID,
								Required: true,
							},
						},
                        Action: func(cCtx *cli.Context) error {
                            app.Env.Delete_Role(app.Param.ID)
                            return nil
                        },
                    },
                },
            },
        },
    }
	return CLI
}

func (app *Application) Init(){
	request_url := fmt.Sprintf("http://%s/login", app.Env.Host)
	json_string := fmt.Sprintf(`{"username": "%s", "password": "%s"}`, app.Env.Username, app.Env.Password)
	jwt_key := app.Env.Connect(request_url,json_string)
	
	fmt.Printf("Your JWT key: %s\n", jwt_key)
	os.Setenv("JWT_KEY", string(jwt_key))
    syscall.Exec(os.Getenv("SHELL"), []string{os.Getenv("SHELL")}, syscall.Environ())
}