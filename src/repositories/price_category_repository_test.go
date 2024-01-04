package repositories

import (
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	kts_errors "github.com/ELITE-Kinoticketsystem/Backend-KTS/src/errors"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/gen/KinoTicketSystem/model"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/managers"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/models"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/samples"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/utils"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestGetPriceCategories(t *testing.T) {
	samplePriceCategories := samples.GetSamplePriceCategories()

	query := "SELECT price_categories.id AS \"price_categories.id\", price_categories.category_name AS \"price_categories.category_name\", price_categories.price AS \"price_categories.price\" FROM `KinoTicketSystem`.price_categories;"

	testCases := []struct {
		name                    string
		setExpectations         func(mock sqlmock.Sqlmock)
		expectedpricecategories *[]model.PriceCategories
		expectedError           *models.KTSError
	}{
		{
			name: "Empty result",
			setExpectations: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery(query).WillReturnRows(
					sqlmock.NewRows([]string{"price_categories.id", "price_categories.category_name", "price_categories.price"}),
				)
			},
			expectedpricecategories: nil,
			expectedError:           kts_errors.KTS_NOT_FOUND,
		},
		{
			name: "Multiple price_categories",
			setExpectations: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery(query).WillReturnRows(
					sqlmock.NewRows(
						[]string{"price_categories.id", "price_categories.category_name", "price_categories.price"},
					).AddRow(
						(*samplePriceCategories)[0].ID, (*samplePriceCategories)[0].CategoryName, (*samplePriceCategories)[0].Price,
					).AddRow(
						(*samplePriceCategories)[1].ID, (*samplePriceCategories)[1].CategoryName, (*samplePriceCategories)[1].Price,
					),
				)
			},
			expectedpricecategories: samplePriceCategories,
			expectedError:           nil,
		},
		{
			name: "Error while querying price_categories",
			setExpectations: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery(query).WillReturnError(sqlmock.ErrCancelled)
			},
			expectedpricecategories: nil,
			expectedError:           kts_errors.KTS_INTERNAL_ERROR,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// GIVEN
			// Create a new mock database connection
			db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
			if err != nil {
				t.Fatalf("Failed to create mock database connection: %v", err)
			}
			defer db.Close()

			// Create a new instance of the PriceCategoryRepository with the mock database connection
			priceCategoryRepo := PriceCategoryRepository{
				DatabaseManager: &managers.DatabaseManager{
					Connection: db,
				},
			}

			tc.setExpectations(mock)

			// WHEN
			// Call the method under test
			priceCategories, kts_err := priceCategoryRepo.GetPriceCategories()

			// THEN
			// Verify the results
			assert.Equal(t, tc.expectedpricecategories, priceCategories)
			assert.Equal(t, tc.expectedError, kts_err)

			// Verify that all expectations were met
			if err = mock.ExpectationsWereMet(); err != nil {
				t.Errorf("There were unfulfilled expectations: %s", err)
			}

		})
	}
}

func TestGetPriceCategoryById(t *testing.T) {
	samplePriceCategory := samples.GetSamplePriceCategory()

	priceCategoryId := samplePriceCategory.ID

	query := "SELECT price_categories.id AS \"price_categories.id\", price_categories.category_name AS \"price_categories.category_name\", price_categories.price AS \"price_categories.price\" FROM `KinoTicketSystem`.price_categories WHERE price_categories.id = ?;"

	testCases := []struct {
		name                    string
		setExpectations         func(mock sqlmock.Sqlmock, id *uuid.UUID)
		expectedpricecategories *model.PriceCategories
		expectedError           *models.KTSError
	}{
		{
			name: "PriceCategory found",
			setExpectations: func(mock sqlmock.Sqlmock, id *uuid.UUID) {
				mock.ExpectQuery(query).WithArgs(utils.EqUUID(priceCategoryId)).WillReturnRows(
					sqlmock.NewRows(
						[]string{"price_categories.id", "price_categories.category_name", "price_categories.price"},
					).AddRow(
						&samplePriceCategory.ID, &samplePriceCategory.CategoryName, &samplePriceCategory.Price,
					),
				)
			},
			expectedpricecategories: samplePriceCategory,
			expectedError:           nil,
		},
		{
			name: "PriceCategory not found",
			setExpectations: func(mock sqlmock.Sqlmock, id *uuid.UUID) {
				mock.ExpectQuery(query).WithArgs(utils.EqUUID(priceCategoryId)).WillReturnRows(
					sqlmock.NewRows([]string{"price_categories.id", "price_categories.category_name", "price_categories.price"}),
				)
			},
			expectedpricecategories: nil,
			expectedError:           kts_errors.KTS_NOT_FOUND,
		},
		{
			name: "Error while querying PriceCategory",
			setExpectations: func(mock sqlmock.Sqlmock, id *uuid.UUID) {
				mock.ExpectQuery(query).WithArgs(utils.EqUUID(priceCategoryId)).WillReturnError(sqlmock.ErrCancelled)
			},
			expectedpricecategories: nil,
			expectedError:           kts_errors.KTS_INTERNAL_ERROR,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// GIVEN
			// Create a new mock database connection
			db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
			if err != nil {
				t.Fatalf("Failed to create mock database connection: %v", err)
			}
			defer db.Close()

			// Create a new instance of the PriceCategoryRepository with the mock database connection
			priceCategoryRepo := PriceCategoryRepository{
				DatabaseManager: &managers.DatabaseManager{
					Connection: db,
				},
			}

			tc.setExpectations(mock, priceCategoryId)

			// WHEN
			// Call the method under test
			priceCategory, kts_err := priceCategoryRepo.GetPriceCategoryById(priceCategoryId)

			// THEN
			// Verify the results
			assert.Equal(t, tc.expectedpricecategories, priceCategory)
			assert.Equal(t, tc.expectedError, kts_err)

			// Verify that all expectations were met
			if err = mock.ExpectationsWereMet(); err != nil {
				t.Errorf("There were unfulfilled expectations: %s", err)
			}

		})
	}
}

func TestCreatePriceCategory(t *testing.T) {
	samplePriceCategory := samples.GetSamplePriceCategory()

	query := "INSERT INTO `KinoTicketSystem`.price_categories (id, category_name, price) VALUES (?, ?, ?);"

	testCases := []struct {
		name                    string
		setExpectations         func(mock sqlmock.Sqlmock, priceCategory *model.PriceCategories)
		expectedPriceCategoryID bool
		expectedError           *models.KTSError
	}{
		{
			name: "PriceCategory created",
			setExpectations: func(mock sqlmock.Sqlmock, priceCategory *model.PriceCategories) {
				mock.ExpectExec(query).WithArgs(sqlmock.AnyArg(), priceCategory.CategoryName, priceCategory.Price).WillReturnResult(sqlmock.NewResult(1, 1))
			},
			expectedPriceCategoryID: true,
			expectedError:           nil,
		},
		{
			name: "Error while creating PriceCategory",
			setExpectations: func(mock sqlmock.Sqlmock, priceCategory *model.PriceCategories) {
				mock.ExpectExec(query).WithArgs(sqlmock.AnyArg(), priceCategory.CategoryName, priceCategory.Price).WillReturnError(sqlmock.ErrCancelled)
			},
			expectedPriceCategoryID: false,
			expectedError:           kts_errors.KTS_INTERNAL_ERROR,
		},
		{
			name: "Error while converting rows affected",
			setExpectations: func(mock sqlmock.Sqlmock, priceCategory *model.PriceCategories) {
				mock.ExpectExec(query).WithArgs(sqlmock.AnyArg(), priceCategory.CategoryName, priceCategory.Price).WillReturnResult(
					sqlmock.NewErrorResult(errors.New("rows affected conversion did not work")),
				)
			},
			expectedPriceCategoryID: false,
			expectedError:           kts_errors.KTS_INTERNAL_ERROR,
		},
		{
			name: "PriceCategory not found",
			setExpectations: func(mock sqlmock.Sqlmock, priceCategory *model.PriceCategories) {
				mock.ExpectExec(query).WithArgs(sqlmock.AnyArg(), priceCategory.CategoryName, priceCategory.Price).WillReturnResult(sqlmock.NewResult(1, 0))
			},
			expectedPriceCategoryID: false,
			expectedError:           kts_errors.KTS_NOT_FOUND,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// GIVEN
			db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
			if err != nil {
				t.Fatalf("Failed to create mock database connection: %v", err)
			}
			defer db.Close()

			// Create a new instance of the PriceCategoryRepository with the mock database connection
			priceCategoryRepo := PriceCategoryRepository{
				DatabaseManager: &managers.DatabaseManager{
					Connection: db,
				},
			}

			tc.setExpectations(mock, samplePriceCategory)

			// WHEN
			priceCategoryID, kts_err := priceCategoryRepo.CreatePriceCategory(samplePriceCategory)

			// THEN
			assert.Equal(t, tc.expectedError, kts_err)
			if tc.expectedPriceCategoryID && priceCategoryID == nil {
				t.Error("Expected priceCategory ID, got nil")
			}

			// Verify that all expectations were met
			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("There were unfulfilled expectations: %s", err)
			}

		})
	}
}

func TestUpdatePriceCategory(t *testing.T) {
	samplePriceCategory := samples.GetSamplePriceCategory()

	query := "UPDATE `KinoTicketSystem`.price_categories SET id = ?, category_name = ?, price = ? WHERE price_categories.id = ?;"

	testCases := []struct {
		name                    string
		setExpectations         func(mock sqlmock.Sqlmock, priceCategory *model.PriceCategories)
		expectedPriceCategoryID bool
		expectedError           *models.KTSError
	}{
		{
			name: "PriceCategory updated",
			setExpectations: func(mock sqlmock.Sqlmock, priceCategory *model.PriceCategories) {
				mock.ExpectExec(query).WithArgs(sqlmock.AnyArg(), priceCategory.CategoryName, priceCategory.Price, utils.EqUUID(priceCategory.ID)).WillReturnResult(sqlmock.NewResult(1, 1))
			},
			expectedPriceCategoryID: true,
			expectedError:           nil,
		},
		{
			name: "Error while updating PriceCategory",
			setExpectations: func(mock sqlmock.Sqlmock, priceCategory *model.PriceCategories) {
				mock.ExpectExec(query).WithArgs(sqlmock.AnyArg(), priceCategory.CategoryName, priceCategory.Price, utils.EqUUID(priceCategory.ID)).WillReturnError(sqlmock.ErrCancelled)
			},
			expectedPriceCategoryID: false,
			expectedError:           kts_errors.KTS_INTERNAL_ERROR,
		},
		{
			name: "Error while converting rows affected",
			setExpectations: func(mock sqlmock.Sqlmock, priceCategory *model.PriceCategories) {
				mock.ExpectExec(query).WithArgs(sqlmock.AnyArg(), priceCategory.CategoryName, priceCategory.Price, utils.EqUUID(priceCategory.ID)).WillReturnResult(
					sqlmock.NewErrorResult(errors.New("rows affected conversion did not work")),
				)
			},
			expectedPriceCategoryID: false,
			expectedError:           kts_errors.KTS_INTERNAL_ERROR,
		},
		{
			name: "PriceCategory not found",
			setExpectations: func(mock sqlmock.Sqlmock, priceCategory *model.PriceCategories) {
				mock.ExpectExec(query).WithArgs(sqlmock.AnyArg(), priceCategory.CategoryName, priceCategory.Price, utils.EqUUID(priceCategory.ID)).WillReturnResult(sqlmock.NewResult(1, 0))
			},
			expectedPriceCategoryID: false,
			expectedError:           kts_errors.KTS_NOT_FOUND,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// GIVEN
			db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
			if err != nil {
				t.Fatalf("Failed to create mock database connection: %v", err)
			}
			defer db.Close()

			// Create a new instance of the PriceCategoryRepository with the mock database connection
			priceCategoryRepo := PriceCategoryRepository{
				DatabaseManager: &managers.DatabaseManager{
					Connection: db,
				},
			}

			tc.setExpectations(mock, samplePriceCategory)

			// WHEN
			priceCategoryID, kts_err := priceCategoryRepo.UpdatePriceCategory(samplePriceCategory)

			// THEN
			assert.Equal(t, tc.expectedError, kts_err)
			if tc.expectedPriceCategoryID && priceCategoryID == nil {
				t.Error("Expected actor ID, got nil")
			}

			// Verify that all expectations were met
			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("There were unfulfilled expectations: %s", err)
			}
		})
	}
}

func TestDeletePriceCategory(t *testing.T) {
	samplePriceCategory := samples.GetSamplePriceCategory()

	priceCategoryId := samplePriceCategory.ID

	query := "DELETE FROM `KinoTicketSystem`.price_categories WHERE price_categories.id = ?;"

	testCases := []struct {
		name            string
		setExpectations func(mock sqlmock.Sqlmock, id *uuid.UUID)
		expectedError   *models.KTSError
	}{
		{
			name: "PriceCategory deleted",
			setExpectations: func(mock sqlmock.Sqlmock, id *uuid.UUID) {
				mock.ExpectExec(query).WithArgs(utils.EqUUID(id)).WillReturnResult(sqlmock.NewResult(1, 1))
			},
			expectedError: nil,
		},
		{
			name: "Error while deleting PriceCategory",
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
			name: "PriceCategory not found",
			setExpectations: func(mock sqlmock.Sqlmock, id *uuid.UUID) {
				mock.ExpectExec(query).WithArgs(utils.EqUUID(id)).WillReturnResult(sqlmock.NewResult(1, 0))
			},
			expectedError: kts_errors.KTS_NOT_FOUND,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// GIVEN
			// Create a new mock database connection
			db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
			if err != nil {
				t.Fatalf("Failed to create mock database connection: %v", err)
			}
			defer db.Close()

			// Create a new instance of the PriceCategoryRepository with the mock database connection
			priceCategoryRepo := PriceCategoryRepository{
				DatabaseManager: &managers.DatabaseManager{
					Connection: db,
				},
			}

			tc.setExpectations(mock, priceCategoryId)

			// WHEN
			// Call the method under test
			kts_err := priceCategoryRepo.DeletePriceCategory(priceCategoryId)

			// THEN
			// Verify the results
			assert.Equal(t, tc.expectedError, kts_err)

			// Verify that all expectations were met
			if err = mock.ExpectationsWereMet(); err != nil {
				t.Errorf("There were unfulfilled expectations: %s", err)
			}

		})
	}
}
