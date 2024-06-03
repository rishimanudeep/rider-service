package migrations

import (
	"gofr.dev/pkg/gofr/migration"
	"log"
)

const create_rider_availabilty = `CREATE TABLE rider_availability (
                                    id SERIAL PRIMARY KEY,
                                    rider_id INT NOT NULL REFERENCES riders(id),
                                    available BOOLEAN NOT NULL,
                                    latitude DECIMAL(9,6) NOT NULL,
                                    longitude DECIMAL(9,6) NOT NULL,
                                    timestamp TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP
);`

func riderAvailableTable() migration.Migrate {
	return migration.Migrate{
		UP: func(d migration.Datasource) error {
			_, err := d.SQL.Exec(create_rider_availabilty)
			if err != nil {
				log.Println(err)
				return err
			}
			return nil
		},
	}
}
