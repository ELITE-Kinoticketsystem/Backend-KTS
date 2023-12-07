package repositories

import (
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/.gen/KinoTicketSystem/table"
	kts_errors "github.com/ELITE-Kinoticketsystem/Backend-KTS/src/errors"

	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/.gen/KinoTicketSystem/model"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/managers"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/models"

	jet_mysql "github.com/go-jet/jet/v2/mysql"
	"github.com/google/uuid"
)

type AddressRepositoryI interface {
	GetAddresses() ([]*model.Addresses, *models.KTSError)
	GetAddressById(id *uuid.UUID) (*model.Addresses, *models.KTSError)
	CreateAddress(address *model.Addresses) *models.KTSError
	UpdateAddress(address *model.Addresses) *models.KTSError
	DeleteAddress(id *uuid.UUID) *models.KTSError
}

type AddressRepository struct {
	DatabaseManager managers.DatabaseManagerI
}

func (mr *MovieRepository) GetAddresses() (*[]model.Addresses, *models.KTSError) {
	var addresses []model.Addresses

	// Create the query
	stmt := jet_mysql.SELECT(
		table.Addresses.AllColumns,
	).FROM(
		table.Addresses,
	)

	// Execute the query
	err := stmt.Query(mr.DatabaseManager.GetDatabaseConnection(), &addresses)
	if err != nil {
		return nil, kts_errors.KTS_INTERNAL_ERROR
	}

	if len(addresses) == 0 {
		return nil, kts_errors.KTS_NOT_FOUND
	}

	return &addresses, nil
}

func (mr *MovieRepository) GetAddressById(id *uuid.UUID) (*model.Addresses, *models.KTSError) {
	var address model.Addresses

	binary_id, _ := id.MarshalBinary()

	// Create the query
	stmt := jet_mysql.SELECT(
		table.Addresses.AllColumns,
	).FROM(
		table.Addresses,
	).WHERE(
		table.Addresses.ID.EQ(jet_mysql.String(string(binary_id))),
	)

	// Execute the query
	err := stmt.Query(mr.DatabaseManager.GetDatabaseConnection(), &address)
	if err != nil {
		return nil, kts_errors.KTS_INTERNAL_ERROR
	}

	return &address, nil
}

// func (mr *MovieRepository) CreateAddress(address *model.Addresses) *models.KTSError {
// 	// Create the query
// 	stmt := jet_mysql.INSERT(
// 		table.Addresses,
// 	).VALUES(
// 		table.Addresses.ID.SET(address.ID),
// 		table.Addresses.Street.SET(address.Street),
// 		table.Addresses.HouseNumber.SET(address.HouseNumber),
// 		table.Addresses.ZipCode.SET(address.ZipCode),
// 		table.Addresses.City.SET(address.City),
// 		table.Addresses.Country.SET(address.Country),
// 	)

// 	// Execute the query
// 	_, err := stmt.Exec(mr.DatabaseManager.GetDatabaseConnection())
// 	if err != nil {
// 		return kts_errors.KTS_INTERNAL_ERROR
// 	}

// 	return nil
// }
