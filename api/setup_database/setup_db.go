package setupDB

import (
	"fmt"

	database "github.com/SamG06/generic-headless-cms-name-later/database"
	_ "github.com/lib/pq"
)

func SetupDatabase(){
	db := database.Connection()
	
	tableQueries := TablesSetup()

	for i := 0; i < len(tableQueries); i++ {

		_, err := db.Exec(tableQueries[i])

		if err != nil {
			panic(err)
		}

		fmt.Println("Created table with query:", tableQueries[i])
	}
}