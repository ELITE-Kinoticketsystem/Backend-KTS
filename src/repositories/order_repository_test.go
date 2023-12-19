package repositories

import (
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/.gen/KinoTicketSystem/model"
	kts_errors "github.com/ELITE-Kinoticketsystem/Backend-KTS/src/errors"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/managers"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/models"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/utils"
)

func TestCreateOrder(t *testing.T) {
	order := &model.Orders{
		ID:              utils.NewUUID(),
		Totalprice:      100,
		IsPaid:          false,
		PaymentMethodID: utils.NewUUID(),
		UserID:          utils.NewUUID(),
	}

	query := "INSERT INTO `KinoTicketSystem`.orders .*"

	testCases := []struct {
		name            string
		setExpectations func(mock sqlmock.Sqlmock)
		expectOrderID   bool
		expectedError   *models.KTSError
	}{
		{
			name: "Create order",
			setExpectations: func(mock sqlmock.Sqlmock) {
				mock.ExpectExec(query).WithArgs(sqlmock.AnyArg(), order.Totalprice, order.IsPaid, sqlmock.AnyArg(), sqlmock.AnyArg()).WillReturnResult(sqlmock.NewResult(1, 1))
			},
			expectOrderID: true,
			expectedError: nil,
		},
		{
			name: "Create order - error",
			setExpectations: func(mock sqlmock.Sqlmock) {
				mock.ExpectExec(query).WithArgs(sqlmock.AnyArg(), order.Totalprice, order.IsPaid, sqlmock.AnyArg(), sqlmock.AnyArg()).WillReturnError(errors.New("error"))
			},
			expectOrderID: false,
			expectedError: kts_errors.KTS_INTERNAL_ERROR,
		},
		{
			name: "Create order - no rows affected",
			setExpectations: func(mock sqlmock.Sqlmock) {
				mock.ExpectExec(query).WithArgs(sqlmock.AnyArg(), order.Totalprice, order.IsPaid, sqlmock.AnyArg(), sqlmock.AnyArg()).WillReturnResult(sqlmock.NewResult(1, 0))
			},
			expectOrderID: false,
			expectedError: kts_errors.KTS_INTERNAL_ERROR,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherRegexp))
			if err != nil {
				t.Fatalf("Failed to create mock database connection: %v", err)
			}
			defer db.Close()

			orderRepo := &OrderRepository{
				DatabaseManager: &managers.DatabaseManager{
					Connection: db,
				},
			}

			tc.setExpectations(mock)

			orderID, ktsErr := orderRepo.CreateOrder(order)

			if ktsErr != tc.expectedError {
				t.Errorf("Unexpected error: %v", ktsErr)
			}

			if tc.expectOrderID && orderID == nil {
				t.Error("Expected order ID, got nil")
			}

			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("There were unfulfilled expectations: %s", err)
			}
		})
	}
}
