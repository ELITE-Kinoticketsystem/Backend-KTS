package repositories

import (
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
