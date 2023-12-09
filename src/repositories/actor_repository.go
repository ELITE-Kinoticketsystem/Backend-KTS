package repositories

import (
	"log"

	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/.gen/KinoTicketSystem/model"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/.gen/KinoTicketSystem/table"
	kts_errors "github.com/ELITE-Kinoticketsystem/Backend-KTS/src/errors"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/managers"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/models"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/utils"

	"github.com/go-jet/jet/v2/mysql"
	"github.com/google/uuid"
)

type ActorRepoI interface {
	GetActorById(actorId *uuid.UUID) (*models.ActorDTO, *models.KTSError)
	GetActors() (*[]models.GetActorsDTO, *models.KTSError)
	CreateActor(actor *model.Actors) (*uuid.UUID, *models.KTSError)

	// Actor pictures
	CreateActorPicture(actorPicture *model.ActorPictures) (*uuid.UUID, *models.KTSError)
}

type ActorRepository struct {
	DatabaseManager *managers.DatabaseManager
}

func (ar *ActorRepository) GetActorById(actorId *uuid.UUID) (*models.ActorDTO, *models.KTSError) {
	var actor models.ActorDTO

	mySqlId, err := utils.MysqlUuid(actorId)

	if err != nil {
		return nil, kts_errors.KTS_INTERNAL_ERROR
	}

	stmt := mysql.SELECT(
		table.Actors.AllColumns,
		table.ActorPictures.AllColumns,
		table.Movies.AllColumns,
	).
		FROM(
			table.Actors.
				LEFT_JOIN(table.ActorPictures, table.ActorPictures.ActorID.EQ(table.Actors.ID)).
				LEFT_JOIN(table.MovieActors, table.MovieActors.ActorID.EQ(table.Actors.ID)).
				LEFT_JOIN(table.Movies, table.Movies.ID.EQ(table.MovieActors.MovieID)),
		).
		WHERE(
			table.Actors.ID.EQ(mySqlId),
		)

	err = stmt.Query(ar.DatabaseManager.GetDatabaseConnection(), &actor)

	if err != nil {
		return nil, kts_errors.KTS_NOT_FOUND
	}

	return &actor, nil
}

func (ar *ActorRepository) GetActors() (*[]models.GetActorsDTO, *models.KTSError) {
	var actors []models.GetActorsDTO

	stmt := mysql.SELECT(
		table.Actors.AllColumns,
		table.ActorPictures.AllColumns,
	).
		FROM(
			table.Actors.
				LEFT_JOIN(table.ActorPictures, table.ActorPictures.ActorID.EQ(table.Actors.ID)),
		)

	err := stmt.Query(ar.DatabaseManager.GetDatabaseConnection(), &actors)

	if err != nil {
		return nil, kts_errors.KTS_INTERNAL_ERROR
	}

	if len(actors) == 0 {
		return nil, kts_errors.KTS_NOT_FOUND
	}

	return &actors, nil
}

func (ar *ActorRepository) CreateActor(actor *model.Actors) (*uuid.UUID, *models.KTSError) {
	actor.ID = utils.NewUUID()

	insertStmt := table.Actors.INSERT(table.Actors.AllColumns).MODEL(actor)

	log.Println(insertStmt.DebugSql())

	rows, err := insertStmt.Exec(ar.DatabaseManager.GetDatabaseConnection())

	if err != nil {
		return nil, kts_errors.KTS_INTERNAL_ERROR
	}

	rowsAffected, err := rows.RowsAffected()

	if err != nil {
		return nil, kts_errors.KTS_INTERNAL_ERROR
	}

	if rowsAffected == 0 {
		return nil, kts_errors.KTS_INTERNAL_ERROR
	}

	return actor.ID, nil
}

func (ar *ActorRepository) CreateActorPicture(actorPicture *model.ActorPictures) (*uuid.UUID, *models.KTSError) {
	actorPicture.ID = utils.NewUUID()

	insertStmt := table.ActorPictures.INSERT(table.ActorPictures.AllColumns).MODEL(actorPicture)

	rows, err := insertStmt.Exec(ar.DatabaseManager.GetDatabaseConnection())

	if err != nil {
		return nil, kts_errors.KTS_INTERNAL_ERROR
	}

	rowsAffected, err := rows.RowsAffected()

	if err != nil {
		return nil, kts_errors.KTS_INTERNAL_ERROR
	}

	if rowsAffected == 0 {
		return nil, kts_errors.KTS_INTERNAL_ERROR
	}

	return actorPicture.ID, nil
}
