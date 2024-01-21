package migrations

func Exec() string{
	exec:= `
		CREATE TABLE IF NOT EXISTS roles (
			id SERIAL PRIMARY KEY,
			name VARCHAR(50) NOT NULL UNIQUE
		);
		CREATE TABLE IF NOT EXISTS users (
			id SERIAL NOT NULL PRIMARY KEY,
			username VARCHAR(50) NOT NULL UNIQUE,
			password TEXT NOT NULL,
			role_id INTEGER NOT NULL REFERENCES roles(id) ON DELETE SET DEFAULT
		);
		CREATE TABLE IF NOT EXISTS folders (
			id SERIAL NOT NULL PRIMARY KEY,
			name VARCHAR(50) NOT NULL,
			owner_id INTEGER NOT NULL REFERENCES users(id) ON DELETE SET DEFAULT,
			parrent_folder_id INTEGER REFERENCES folders(id) ON DELETE CASCADE
		);
		CREATE TABLE IF NOT EXISTS secrets (
			id SERIAL NOT NULL PRIMARY KEY,
			name TEXT ,
			username TEXT ,
			secret TEXT,
			link TEXT,
			description TEXT,
			owner_id INTEGER NOT NULL REFERENCES users(id) ON DELETE SET DEFAULT,
			folder_id INTEGER NOT NULL REFERENCES folders(id) ON DELETE CASCADE
		);
	`
	return exec
}

/*
"
SELECT s.name
FROM secrets AS s
WHERE s.owner_id  = (SELECT id FROM users WHERE username='user') AND s.folder_id  = (SELECT id FROM folders WHERE name='root');

SELECT f.name
FROM folders AS f
WHERE f.owner_id  = (SELECT id FROM users WHERE username='user') AND f.parrent_folder_id  = (SELECT id FROM folders WHERE name='root');

"
*/
	

