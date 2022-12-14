package setupDB

func TablesSetup() []string {
	
	Posts := `CREATE TABLE IF NOT EXISTS posts(
		id SERIAL PRIMARY KEY,
		title VARCHAR (100),
		summary VARCHAR (300),
		content TEXT,
		author_id INT,
		FOREIGN KEY (author_id) REFERENCES authors(id)
	)`;

	Authors := `CREATE TABLE IF NOT EXISTS authors(
		id SERIAL PRIMARY KEY,
		name VARCHAR(100),
		bio VARCHAR(250)
	)`;

	return []string{Authors, Posts}

}