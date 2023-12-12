package repositories

import (
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/.gen/KinoTicketSystem/table"
	kts_errors "github.com/ELITE-Kinoticketsystem/Backend-KTS/src/errors"

	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/.gen/KinoTicketSystem/model"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/managers"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/models"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/utils"

	jet_mysql "github.com/go-jet/jet/v2/mysql"
	"github.com/google/uuid"
)

type AddressRepositoryI interface {
	GetAddresses() (*[]model.Addresses, *models.KTSError)
	GetAddressById(id *uuid.UUID) (*model.Addresses, *models.KTSError)
	CreateAddress(address *model.Addresses) *models.KTSError
	UpdateAddress(address *model.Addresses) *models.KTSError
	DeleteAddress(id *uuid.UUID) *models.KTSError
}

type AddressRepository struct {
	DatabaseManager managers.DatabaseManagerI
}

func (ar *AddressRepository) GetAddresses() (*[]model.Addresses, *models.KTSError) {
	var addresses []model.Addresses

	// Create the query
	stmt := jet_mysql.SELECT(
		table.Addresses.AllColumns,
	).FROM(
		table.Addresses,
	)

	// Execute the query
	err := stmt.Query(ar.DatabaseManager.GetDatabaseConnection(), &addresses)
	if err != nil {
		return nil, kts_errors.KTS_INTERNAL_ERROR
	}

	if len(addresses) == 0 {
		return nil, kts_errors.KTS_NOT_FOUND
	}

	return &addresses, nil
}

func (ar *AddressRepository) GetAddressById(id *uuid.UUID) (*model.Addresses, *models.KTSError) {
	var address model.Addresses
	// Create the query
	stmt := jet_mysql.SELECT(
		table.Addresses.AllColumns,
	).FROM(
		table.Addresses,
	).WHERE(
		table.Addresses.ID.EQ(utils.MysqlUuid(id)),
	)

	// Execute the query
	err := stmt.Query(ar.DatabaseManager.GetDatabaseConnection(), &address)
	if err != nil {
		if err.Error() == "qrm: no rows in result set" {
			return nil, kts_errors.KTS_NOT_FOUND
		}
		return nil, kts_errors.KTS_INTERNAL_ERROR
	}

	return &address, nil
}

func (ar *AddressRepository) CreateAddress(address *model.Addresses) *models.KTSError {
	// Create the query
	stmt := table.Addresses.INSERT(
		table.Addresses.AllColumns,
	).VALUES(
		table.Addresses.Street.SET(jet_mysql.String(address.Street)),
		table.Addresses.StreetNr.SET(jet_mysql.String(address.StreetNr)),
		table.Addresses.Zipcode.SET(jet_mysql.String(address.Zipcode)),
		table.Addresses.City.SET(jet_mysql.String(address.City)),
		table.Addresses.Country.SET(jet_mysql.String(address.Country)),
	)

	// Execute the query
	rows, err := stmt.Exec(ar.DatabaseManager.GetDatabaseConnection())
	if err != nil {
		return kts_errors.KTS_INTERNAL_ERROR
	}

	rowsAffected, err := rows.RowsAffected()
	if err != nil {
		return kts_errors.KTS_INTERNAL_ERROR
	}

	if rowsAffected == 0 {
		return kts_errors.KTS_NOT_FOUND
	}

	return nil
}

func (ar *AddressRepository) UpdateAddress(address *model.Addresses) *models.KTSError {
	// Create the query
	stmt := table.Addresses.UPDATE(
		table.Addresses.AllColumns,
	).SET(
		table.Addresses.ID.SET(jet_mysql.String(address.ID.String())),
		table.Addresses.Street.SET(jet_mysql.String(address.Street)),
		table.Addresses.StreetNr.SET(jet_mysql.String(address.StreetNr)),
		table.Addresses.Zipcode.SET(jet_mysql.String(address.Zipcode)),
		table.Addresses.City.SET(jet_mysql.String(address.City)),
		table.Addresses.Country.SET(jet_mysql.String(address.Country)),
	).WHERE(
		table.Addresses.ID.EQ(utils.MysqlUuid(address.ID)),
	)

	// Execute the query
	rows, err := stmt.Exec(ar.DatabaseManager.GetDatabaseConnection())
	if err != nil {
		return kts_errors.KTS_INTERNAL_ERROR
	}

	rowsAffected, err := rows.RowsAffected()
	if err != nil {
		return kts_errors.KTS_INTERNAL_ERROR
	}

	if rowsAffected == 0 {
		return kts_errors.KTS_NOT_FOUND
	}

	return nil
}

func (ar *AddressRepository) DeleteAddress(id *uuid.UUID) *models.KTSError {
	// Create the query
	stmt := table.Addresses.DELETE().WHERE(table.Addresses.ID.EQ(utils.MysqlUuid(id)))

	// Execute the query
	rows, err := stmt.Exec(ar.DatabaseManager.GetDatabaseConnection())
	if err != nil {
		return kts_errors.KTS_INTERNAL_ERROR
	}

	rowsAffected, err := rows.RowsAffected()
	if err != nil {
		return kts_errors.KTS_INTERNAL_ERROR
	}

	if rowsAffected == 0 {
		return kts_errors.KTS_NOT_FOUND
	}

	return nil
}
