package repositories

import (
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/.gen/KinoTicketSystem/model"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/.gen/KinoTicketSystem/table"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/managers"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/models"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/utils"
	"github.com/google/uuid"

	kts_errors "github.com/ELITE-Kinoticketsystem/Backend-KTS/src/errors"

	jet_mysql "github.com/go-jet/jet/v2/mysql"
)

type PriceCategoryRepositoryI interface {
	GetPriceCategories() (*[]model.PriceCategories, *models.KTSError)
	GetPriceCategoryById(id *uuid.UUID) (*model.PriceCategories, *models.KTSError)
	CreatePriceCategory(priceCategory *model.PriceCategories) *models.KTSError
	UpdatePriceCategory(priceCategory *model.PriceCategories) *models.KTSError
	DeletePriceCategory(id *uuid.UUID) *models.KTSError
}

type PriceCategoryRepository struct {
	DatabaseManager managers.DatabaseManagerI
}

func (pcr *PriceCategoryRepository) GetPriceCategories() (*[]model.PriceCategories, *models.KTSError) {
	var priceCategory []model.PriceCategories

	// Create the query
	stmt := jet_mysql.SELECT(
		table.PriceCategories.AllColumns,
	).FROM(
		table.PriceCategories,
	)

	// Execute the query
	err := stmt.Query(pcr.DatabaseManager.GetDatabaseConnection(), &priceCategory)
	if err != nil {
		return nil, kts_errors.KTS_INTERNAL_ERROR
	}

	if len(priceCategory) == 0 {
		return nil, kts_errors.KTS_NOT_FOUND
	}

	return &priceCategory, nil
}

func (pcr *PriceCategoryRepository) GetPriceCategoryById(id *uuid.UUID) (*model.PriceCategories, *models.KTSError) {
	var priceCategory model.PriceCategories

	// Create the query
	stmt := jet_mysql.SELECT(
		table.PriceCategories.AllColumns,
	).FROM(
		table.PriceCategories,
	).WHERE(
		table.PriceCategories.ID.EQ(utils.MysqlUuid(id)),
	)

	// Execute the query
	err := stmt.Query(pcr.DatabaseManager.GetDatabaseConnection(), &priceCategory)
	if err != nil {
		if err.Error() == "qrm: no rows in result set" {
			return nil, kts_errors.KTS_NOT_FOUND
		}
		return nil, kts_errors.KTS_INTERNAL_ERROR
	}

	return &priceCategory, nil
}

func (pcr *PriceCategoryRepository) CreatePriceCategory(priceCategory *model.PriceCategories) *models.KTSError {
	// Create the query
	stmt := table.PriceCategories.INSERT(
		table.PriceCategories.AllColumns,
	).VALUES(
		table.PriceCategories.CategoryName.SET(jet_mysql.String(priceCategory.CategoryName)),
		table.PriceCategories.Price.SET(jet_mysql.Int32(priceCategory.Price)),
	)

	// Execute the query
	rows, err := stmt.Exec(pcr.DatabaseManager.GetDatabaseConnection())
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

func (pcr *PriceCategoryRepository) UpdatePriceCategory(priceCategory *model.PriceCategories) *models.KTSError {
	// Create the query
	stmt := table.PriceCategories.UPDATE(
		table.PriceCategories.AllColumns,
	).SET(
		table.PriceCategories.ID.SET(jet_mysql.String(priceCategory.ID.String())),
		table.PriceCategories.CategoryName.SET(jet_mysql.String(priceCategory.CategoryName)),
		table.PriceCategories.Price.SET(jet_mysql.Int32(priceCategory.Price)),
	).WHERE(
		table.PriceCategories.ID.EQ(utils.MysqlUuid(priceCategory.ID)),
	)

	// Execute the query
	rows, err := stmt.Exec(pcr.DatabaseManager.GetDatabaseConnection())
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

func (pcr *PriceCategoryRepository) DeletePriceCategory(id *uuid.UUID) *models.KTSError {
	// Create the query
	stmt := table.PriceCategories.DELETE().WHERE(table.PriceCategories.ID.EQ(utils.MysqlUuid(id)))

	// Execute the query
	rows, err := stmt.Exec(pcr.DatabaseManager.GetDatabaseConnection())
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
