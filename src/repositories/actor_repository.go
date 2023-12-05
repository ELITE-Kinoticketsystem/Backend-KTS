package repositories

import (
	"log"

	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/.gen/KinoTicketSystem/table"
	kts_errors "github.com/ELITE-Kinoticketsystem/Backend-KTS/src/errors"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/managers"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/models"

	"github.com/go-jet/jet/v2/mysql"
	"github.com/google/uuid"
)

type ActorRepoI interface {
	GetActorById(actorId *uuid.UUID) (*models.ActorDTO, *models.KTSError)
}

type ActorRepository struct {
	DatabaseManager *managers.DatabaseManager
}

func (ar *ActorRepository) GetActorById(actorId *uuid.UUID) (*models.ActorDTO, *models.KTSError) {
	var actor models.ActorDTO

	binary_id, _ := actorId.MarshalBinary()

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
			table.Actors.ID.EQ(mysql.String(string(binary_id))),
		)

	err := stmt.Query(ar.DatabaseManager.GetDatabaseConnection(), &actor)

	if err != nil {
		return nil, kts_errors.KTS_INTERNAL_ERROR
	}

	log.Println(actor)

	return &actor, nil
}
