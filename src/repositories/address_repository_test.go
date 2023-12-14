package repositories

import (
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/.gen/KinoTicketSystem/model"
	kts_errors "github.com/ELITE-Kinoticketsystem/Backend-KTS/src/errors"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/managers"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/models"
	"github.com/google/uuid"

	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/utils"
	"github.com/stretchr/testify/assert"
)

func TestGetAddresses(t *testing.T) {
	sampleAddresses := utils.GetSampleAddresses()

	query := "SELECT addresses.id AS \"addresses.id\", addresses.street AS \"addresses.street\", addresses.street_nr AS \"addresses.street_nr\", addresses.zipcode AS \"addresses.zipcode\", addresses.city AS \"addresses.city\", addresses.country AS \"addresses.country\" FROM `KinoTicketSystem`.addresses;"

	testCases := []struct {
		name              string
		setExpectations   func(mock sqlmock.Sqlmock)
		expectedaddresses *[]model.Addresses
		expectedError     *models.KTSError
	}{
		{
			name: "Empty result",
			setExpectations: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery(query).WillReturnRows(
					sqlmock.NewRows([]string{"addresses.id", "addresses.street", "addresses.street_nr", "addresses.zipcode", "addresses.city", "addresses.country"}),
				)
			},
			expectedaddresses: nil,
			expectedError:     kts_errors.KTS_NOT_FOUND,
		},
		{
			name: "Multiple addresses",
			setExpectations: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery(query).WillReturnRows(
					sqlmock.NewRows(
						[]string{"addresses.id", "addresses.street", "addresses.street_nr", "addresses.zipcode", "addresses.city", "addresses.country"},
					).AddRow(
						(*sampleAddresses)[0].ID, (*sampleAddresses)[0].Street, (*sampleAddresses)[0].StreetNr, (*sampleAddresses)[0].Zipcode, (*sampleAddresses)[0].City, (*sampleAddresses)[0].Country,
					).AddRow(
						(*sampleAddresses)[1].ID, (*sampleAddresses)[1].Street, (*sampleAddresses)[1].StreetNr, (*sampleAddresses)[1].Zipcode, (*sampleAddresses)[1].City, (*sampleAddresses)[1].Country,
					),
				)
			},
			expectedaddresses: sampleAddresses,
			expectedError:     nil,
		},
		{
			name: "Error while querying addresses",
			setExpectations: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery(query).WillReturnError(sqlmock.ErrCancelled)
			},
			expectedaddresses: nil,
			expectedError:     kts_errors.KTS_INTERNAL_ERROR,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Create a new mock database connection
			db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
			if err != nil {
				t.Fatalf("Failed to create mock database connection: %v", err)
			}
			defer db.Close()

			// Create a new instance of the AddressRepository with the mock database connection
			addressRepo := AddressRepository{
				DatabaseManager: &managers.DatabaseManager{
					Connection: db,
				},
			}

			tc.setExpectations(mock)

			// Call the method under test
			addresses, kts_err := addressRepo.GetAddresses()

			// Verify the results
			assert.Equal(t, tc.expectedaddresses, addresses)
			assert.Equal(t, tc.expectedError, kts_err)

			// Verify that all expectations were met
			if err = mock.ExpectationsWereMet(); err != nil {
				t.Errorf("There were unfulfilled expectations: %s", err)
			}

		})
	}
}

func TestGetAddressById(t *testing.T) {
	sampleAddress := utils.GetSampleAddress()

	addressId := sampleAddress.ID

	query := "SELECT addresses.id AS \"addresses.id\", addresses.street AS \"addresses.street\", addresses.street_nr AS \"addresses.street_nr\", addresses.zipcode AS \"addresses.zipcode\", addresses.city AS \"addresses.city\", addresses.country AS \"addresses.country\" FROM `KinoTicketSystem`.addresses WHERE addresses.id = ?;"

	testCases := []struct {
		name            string
		setExpectations func(mock sqlmock.Sqlmock, id *uuid.UUID)
		expectedAddress *model.Addresses
		expectedError   *models.KTSError
	}{
		{
			name: "Address found",
			setExpectations: func(mock sqlmock.Sqlmock, id *uuid.UUID) {
				mock.ExpectQuery(query).WithArgs(utils.EqUUID(addressId)).WillReturnRows(
					sqlmock.NewRows(
						[]string{"addresses.id", "addresses.street", "addresses.street_nr", "addresses.zipcode", "addresses.city", "addresses.country"},
					).AddRow(
						&sampleAddress.ID, &sampleAddress.Street, &sampleAddress.StreetNr, &sampleAddress.Zipcode, &sampleAddress.City, &sampleAddress.Country,
					),
				)
			},
			expectedAddress: sampleAddress,
			expectedError:   nil,
		},
		{
			name: "Address not found",
			setExpectations: func(mock sqlmock.Sqlmock, id *uuid.UUID) {
				mock.ExpectQuery(query).WithArgs(utils.EqUUID(addressId)).WillReturnRows(
					sqlmock.NewRows([]string{"addresses.id", "addresses.street", "addresses.street_nr", "addresses.zipcode", "addresses.city", "addresses.country"}),
				)
			},
			expectedAddress: nil,
			expectedError:   kts_errors.KTS_NOT_FOUND,
		},
		{
			name: "Error while querying address",
			setExpectations: func(mock sqlmock.Sqlmock, id *uuid.UUID) {
				mock.ExpectQuery(query).WithArgs(utils.EqUUID(addressId)).WillReturnError(sqlmock.ErrCancelled)
			},
			expectedAddress: nil,
			expectedError:   kts_errors.KTS_INTERNAL_ERROR,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Create a new mock database connection
			db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
			if err != nil {
				t.Fatalf("Failed to create mock database connection: %v", err)
			}
			defer db.Close()

			// Create a new instance of the AddressRepository with the mock database connection
			addressRepo := AddressRepository{
				DatabaseManager: &managers.DatabaseManager{
					Connection: db,
				},
			}

			tc.setExpectations(mock, addressId)

			// Call the method under test
			address, kts_err := addressRepo.GetAddressById(addressId)

			// Verify the results
			assert.Equal(t, tc.expectedAddress, address)
			assert.Equal(t, tc.expectedError, kts_err)

			// Verify that all expectations were met
			if err = mock.ExpectationsWereMet(); err != nil {
				t.Errorf("There were unfulfilled expectations: %s", err)
			}

		})
	}
}

func TestCreateAddress(t *testing.T) {
	sampleAddress := utils.GetSampleAddress()

	query := "INSERT INTO `KinoTicketSystem`.addresses (id, street, street_nr, zipcode, city, country) VALUES (id = ?, street = ?, street_nr = ?, zipcode = ?, city = ?, country = ?);"

	testCases := []struct {
		name              string
		setExpectations   func(mock sqlmock.Sqlmock, address *model.Addresses)
		expectedAddressId bool
		expectedError     *models.KTSError
	}{
		{
			name: "Address created",
			setExpectations: func(mock sqlmock.Sqlmock, address *model.Addresses) {
				mock.ExpectExec(query).WithArgs(sqlmock.AnyArg(), address.Street, address.StreetNr, address.Zipcode, address.City, address.Country).WillReturnResult(sqlmock.NewResult(1, 1))
			},
			expectedAddressId: true,
			expectedError:     nil,
		},
		{
			name: "Error while creating address",
			setExpectations: func(mock sqlmock.Sqlmock, address *model.Addresses) {
				mock.ExpectExec(query).WithArgs(sqlmock.AnyArg(), address.Street, address.StreetNr, address.Zipcode, address.City, address.Country).WillReturnError(sqlmock.ErrCancelled)
			},
			expectedAddressId: false,
			expectedError:     kts_errors.KTS_INTERNAL_ERROR,
		},

		{
			name: "Error while converting rows affected",
			setExpectations: func(mock sqlmock.Sqlmock, address *model.Addresses) {
				mock.ExpectExec(query).WithArgs(sqlmock.AnyArg(), address.Street, address.StreetNr, address.Zipcode, address.City, address.Country).WillReturnResult(
					sqlmock.NewErrorResult(errors.New("rows affected conversion did not work")),
				)
			},
			expectedAddressId: false,
			expectedError:     kts_errors.KTS_INTERNAL_ERROR,
		},
		{
			name: "Movie not found",
			setExpectations: func(mock sqlmock.Sqlmock, address *model.Addresses) {
				mock.ExpectExec(query).WithArgs(sqlmock.AnyArg(), address.Street, address.StreetNr, address.Zipcode, address.City, address.Country).WillReturnResult(sqlmock.NewResult(1, 0))
			},
			expectedAddressId: false,
			expectedError:     kts_errors.KTS_NOT_FOUND,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Create a new mock database connection
			db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
			if err != nil {
				t.Fatalf("Failed to create mock database connection: %v", err)
			}
			defer db.Close()

			// Create a new instance of the AddressRepository with the mock database connection
			addressRepo := AddressRepository{
				DatabaseManager: &managers.DatabaseManager{
					Connection: db,
				},
			}

			tc.setExpectations(mock, sampleAddress)

			// Call the method under test
			addressId, kts_err := addressRepo.CreateAddress(sampleAddress)

			// Verify the results
			assert.Equal(t, tc.expectedError, kts_err)

			if tc.expectedAddressId && addressId == nil {
				t.Error("Expected actor ID, got nil")
			}

			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("There were unfulfilled expectations: %s", err)
			}

		})
	}
}

func TestUpdateAddress(t *testing.T) {
	sampleAddress := utils.GetSampleAddress()

	query := "UPDATE `KinoTicketSystem`.addresses SET id = ?, street = ?, street_nr = ?, zipcode = ?, city = ?, country = ? WHERE addresses.id = ?;"

	testCases := []struct {
		name            string
		setExpectations func(mock sqlmock.Sqlmock, address *model.Addresses)
		expectedError   *models.KTSError
	}{
		{
			name: "Address updated",
			setExpectations: func(mock sqlmock.Sqlmock, address *model.Addresses) {
				mock.ExpectExec(query).WithArgs(sqlmock.AnyArg(), address.Street, address.StreetNr, address.Zipcode, address.City, address.Country, utils.EqUUID(address.ID)).WillReturnResult(sqlmock.NewResult(1, 1))
			},
			expectedError: nil,
		},
		{
			name: "Error while updating address",
			setExpectations: func(mock sqlmock.Sqlmock, address *model.Addresses) {
				mock.ExpectExec(query).WithArgs(sqlmock.AnyArg(), address.Street, address.StreetNr, address.Zipcode, address.City, address.Country, utils.EqUUID(address.ID)).WillReturnError(sqlmock.ErrCancelled)
			},
			expectedError: kts_errors.KTS_INTERNAL_ERROR,
		},

		{
			name: "Error while converting rows affected",
			setExpectations: func(mock sqlmock.Sqlmock, address *model.Addresses) {
				mock.ExpectExec(query).WithArgs(sqlmock.AnyArg(), address.Street, address.StreetNr, address.Zipcode, address.City, address.Country, utils.EqUUID(address.ID)).WillReturnResult(
					sqlmock.NewErrorResult(errors.New("rows affected conversion did not work")),
				)
			},
			expectedError: kts_errors.KTS_INTERNAL_ERROR,
		},
		{
			name: "Movie not found",
			setExpectations: func(mock sqlmock.Sqlmock, address *model.Addresses) {
				mock.ExpectExec(query).WithArgs(sqlmock.AnyArg(), address.Street, address.StreetNr, address.Zipcode, address.City, address.Country, utils.EqUUID(address.ID)).WillReturnResult(sqlmock.NewResult(1, 0))
			},
			expectedError: kts_errors.KTS_NOT_FOUND,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Create a new mock database connection
			db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
			if err != nil {
				t.Fatalf("Failed to create mock database connection: %v", err)
			}
			defer db.Close()

			// Create a new instance of the AddressRepository with the mock database connection
			addressRepo := AddressRepository{
				DatabaseManager: &managers.DatabaseManager{
					Connection: db,
				},
			}

			tc.setExpectations(mock, sampleAddress)

			// Call the method under test
			kts_err := addressRepo.UpdateAddress(sampleAddress)

			// Verify the results
			assert.Equal(t, tc.expectedError, kts_err)

			// Verify that all expectations were met
			if err = mock.ExpectationsWereMet(); err != nil {
				t.Errorf("There were unfulfilled expectations: %s", err)
			}

		})
	}
}

func TestDeleteAddress(t *testing.T) {
	sampleAddress := utils.GetSampleAddress()

	id := sampleAddress.ID

	query := "DELETE FROM `KinoTicketSystem`.addresses WHERE addresses.id = ?;"

	testCases := []struct {
		name            string
		setExpectations func(mock sqlmock.Sqlmock, id *uuid.UUID)
		expectedError   *models.KTSError
	}{
		{
			name: "Address deleted",
			setExpectations: func(mock sqlmock.Sqlmock, id *uuid.UUID) {
				mock.ExpectExec(query).WithArgs(utils.EqUUID(id)).WillReturnResult(sqlmock.NewResult(1, 1))
			},
			expectedError: nil,
		},
		{
			name: "Error while deleting address",
			setExpectations: func(mock sqlmock.Sqlmock, id *uuid.UUID) {
				mock.ExpectExec(query).WithArgs(utils.EqUUID(id)).WillReturnError(sqlmock.ErrCancelled)
			},
			expectedError: kts_errors.KTS_INTERNAL_ERROR,
		},

		{
			name: "Error while converting rows affected",
			setExpectations: func(mock sqlmock.Sqlmock, id *uuid.UUID) {
				mock.ExpectExec(query).WithArgs(utils.EqUUID(id)).WillReturnResult(
					sqlmock.NewErrorResult(errors.New("rows affected conversion did not work")),
				)
			},
			expectedError: kts_errors.KTS_INTERNAL_ERROR,
		},
		{
			name: "Movie not found",
			setExpectations: func(mock sqlmock.Sqlmock, id *uuid.UUID) {
				mock.ExpectExec(query).WithArgs(utils.EqUUID(id)).WillReturnResult(sqlmock.NewResult(1, 0))
			},
			expectedError: kts_errors.KTS_NOT_FOUND,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Create a new mock database connection
			db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
			if err != nil {
				t.Fatalf("Failed to create mock database connection: %v", err)
			}
			defer db.Close()

			// Create a new instance of the AddressRepository with the mock database connection
			addressRepo := AddressRepository{
				DatabaseManager: &managers.DatabaseManager{
					Connection: db,
				},
			}

			tc.setExpectations(mock, id)

			// Call the method under test
			kts_err := addressRepo.DeleteAddress(id)

			// Verify the results
			assert.Equal(t, tc.expectedError, kts_err)

			// Verify that all expectations were met
			if err = mock.ExpectationsWereMet(); err != nil {
				t.Errorf("There were unfulfilled expectations: %s", err)
			}

		})
	}
}
