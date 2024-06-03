package migrations

import (
	"gofr.dev/pkg/gofr/migration"
	"log"
)

const create_riders_table = `CREATE TABLE riders (
                        id SERIAL PRIMARY KEY,
                        name VARCHAR(100) NOT NULL,
                        email VARCHAR(100) UNIQUE NOT NULL,
                        vehicle_details VARCHAR(100),
                        created TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
                        updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);`

func createRiderTable() migration.Migrate {
	return migration.Migrate{
		UP: func(d migration.Datasource) error {
			_, err := d.SQL.Exec(create_riders_table)
			if err != nil {
				log.Println(err)
				return err
			}
			return nil
		},
	}
}
