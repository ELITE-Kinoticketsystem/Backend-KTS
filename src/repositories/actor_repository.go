package repositories

import (
	"log"

	. "github.com/ELITE-Kinoticketsystem/Backend-KTS/src/.gen/KinoTicketSystem/table"
	kts_errors "github.com/ELITE-Kinoticketsystem/Backend-KTS/src/errors"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/managers"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/models"

	. "github.com/go-jet/jet/v2/mysql"
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

	// for the Movies you have to go through the movie_actors table
	// Please use a union

	stmt := SELECT(
		Actors.AllColumns,
		ActorPictures.AllColumns,
		Movies.AllColumns,
	).
		FROM(
			Actors.
				LEFT_JOIN(ActorPictures, ActorPictures.ActorID.EQ(Actors.ID)).
				LEFT_JOIN(MovieActors, MovieActors.ActorID.EQ(Actors.ID)).
				LEFT_JOIN(Movies, Movies.ID.EQ(MovieActors.MovieID)),
		).
		WHERE(
			Actors.ID.EQ(String(string(binary_id))),
		)

	err := stmt.Query(ar.DatabaseManager.GetDatabaseConnection(), &actor)

	if err != nil {
		log.Println(err)
		return nil, kts_errors.KTS_INTERNAL_ERROR
	}

	log.Println(actor)

	return &actor, nil
}
