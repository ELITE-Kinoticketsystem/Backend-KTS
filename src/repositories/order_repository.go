package repositories

import (
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/.gen/KinoTicketSystem/model"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/.gen/KinoTicketSystem/table"
	kts_errors "github.com/ELITE-Kinoticketsystem/Backend-KTS/src/errors"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/managers"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/models"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/utils"
	"github.com/google/uuid"
)

type OrderRepoI interface {
	CreateOrder(order *model.Orders) (*uuid.UUID, *models.KTSError)
}

type OrderRepository struct {
	DatabaseManager managers.DatabaseManagerI
}

func (or *OrderRepository) CreateOrder(order *model.Orders) (*uuid.UUID, *models.KTSError) {

	insertStmt := table.Orders.INSERT(table.Orders.AllColumns).
		VALUES(
			utils.MysqlUuid(order.ID),
			order.Totalprice,
			order.IsPaid,
			utils.MysqlUuid(order.PaymentMethodID),
			utils.MysqlUuid(order.UserID),
		)

	rows, err := insertStmt.Exec(or.DatabaseManager.GetDatabaseConnection())

	if err != nil {
		return nil, kts_errors.KTS_INTERNAL_ERROR
	}

	affectedRows, err := rows.RowsAffected()

	if err != nil {
		return nil, kts_errors.KTS_INTERNAL_ERROR
	}

	if affectedRows == 0 {
		return nil, kts_errors.KTS_INTERNAL_ERROR
	}

	return order.ID, nil
}
