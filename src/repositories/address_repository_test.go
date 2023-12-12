package repositories

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/.gen/KinoTicketSystem/model"
	kts_errors "github.com/ELITE-Kinoticketsystem/Backend-KTS/src/errors"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/managers"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/models"

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
