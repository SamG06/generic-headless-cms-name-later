package database

import (
	"database/sql"
	"fmt"

	env "github.com/SamG06/generic-headless-cms-name-later/env_variables"
)

func Connection() *sql.DB {
	config := env.ProjectEnvs()

	psqlAuth := fmt.Sprintf("host=%s port=%s user=%s " + 
	"password=%s dbname=%s", config.Host, config.Port,
	 config.DBuser, config.Password, config.DBname)
	
	 db, err := sql.Open("postgres", psqlAuth)

	 if err != nil {
		print(err)
	 }

	 err = db.Ping()
	 
	 if err != nil {
	   panic(err)
	 }

	 fmt.Println("Successfully connected to database!")

	 return db
}