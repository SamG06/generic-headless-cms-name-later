package setupDB

import (
	"database/sql"
	"fmt"

	env "github.com/SamG06/generic-headless-cms-name-later/env_variables"
	_ "github.com/lib/pq"
)

func SetupDatabase(){
	config := env.ProjectEnvs()

	psqlAuth := fmt.Sprintf("host=%s port=%s user=%s " + 
	"password=%s dbname=%s", config.Host, config.Port,
	 config.DBuser, config.Password, config.DBname)
	
	 db, err := sql.Open("postgres", psqlAuth)

	 defer db.Close()

	 err = db.Ping()
	 if err != nil {
	   panic(err)
	 }
   
	 fmt.Println("Successfully connected!")

	 tableQueries := TablesSetup()

	 for i := 0; i < len(tableQueries); i++ {
		res, err := db.Exec(tableQueries[i])

		if err != nil {
			panic(err)
		}

		fmt.Println(res)
	 }
}