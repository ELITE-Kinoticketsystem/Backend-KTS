package repositories

import (
	kts_errors "github.com/ELITE-Kinoticketsystem/Backend-KTS/src/errors"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/gen/KinoTicketSystem/table"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/managers"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/models"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/myid"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/utils"
)

type UserMovieRepositoryI interface {
	RemoveAllUserMovieCombinationWithMovie(movieId *myid.UUID) *models.KTSError
}

type UserMovieRepository struct {
	DatabaseManager managers.DatabaseManagerI
}

func (umr *UserMovieRepository) RemoveAllUserMovieCombinationWithMovie(movieId *myid.UUID) *models.KTSError {
	deleteQuery := table.UserMovies.DELETE().WHERE(
		table.UserMovies.MovieID.EQ(utils.MysqlUuid(*movieId)),
	)

	// Execute the query
	_, err := deleteQuery.Exec(umr.DatabaseManager.GetDatabaseConnection())
	if err != nil {
		return kts_errors.KTS_INTERNAL_ERROR
	}

	return nil

}
