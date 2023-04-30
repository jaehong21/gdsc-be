package storage

import (
	"database/sql"

	"github.com/jaehong21/ga-be/entity"
)

func FindOneBuilding(db *sql.DB, id int) (entity.Building, error) {
	var building entity.Building
	row := db.QueryRow("SELECT id, name, location, network_address, subnet_mask FROM \"building\" WHERE id=$1", id)
	err := row.Scan(&building.ID, &building.Name, &building.Location, &building.NetworkAddress, &building.SubnetMask)
	if err != nil {
		return building, err
	}

	return building, nil
}
