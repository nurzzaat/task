package pkg

import (
	"database/sql"

)

type Application struct {
	Env *Env
	Sql *sql.DB
}

func App() (Application, error) {
	app := &Application{}
	app.Env = NewEnv()
	conn, err := NewSQLConn(app.Env)
	if err != nil {
		return Application{}, err
	}
	app.Sql = conn

	return *app, nil
}
