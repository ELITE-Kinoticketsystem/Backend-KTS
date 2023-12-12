package main

import (
	"database/sql"
	"net/http"

	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/controllers"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/handlers"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/managers"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/middlewares"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/repositories"
	"github.com/gin-gonic/gin"
)

type Controllers struct {
	UserController  controllers.UserControllerI
	MovieController controllers.MovieControllerI
	GenreController controllers.GenreControllerI
}

func createRouter(dbConnection *sql.DB) *gin.Engine {
	router := gin.Default()

	// Attach Middleware
	router.Use(middlewares.CorsMiddleware())

	// Create api groups, with special middleware
	publicRoutes := router.Group("/")
	securedRoutes := router.Group("/", middlewares.JwtAuthMiddleware())

	// Create managers and repositories
	databaseManager := &managers.DatabaseManager{
		Connection: dbConnection,
	}

	userRepo := &repositories.UserRepository{
		DatabaseManager: databaseManager,
	}

	movieRepo := &repositories.MovieRepository{
		DatabaseManager: databaseManager,
	}

	genreRepo := &repositories.GenreRepository{
		DatabaseManager: databaseManager,
	}

	movieGenreRepo := &repositories.MovieGenreRepository{
		DatabaseManager: databaseManager,
	}

	// theatreRepo := &repositories.TheatreRepository{
	// 	DatabaseManager: databaseManager,
	// }

	// Create controllers
	controller := Controllers{
		UserController: &controllers.UserController{
			UserRepo: userRepo,
		},
		MovieController: &controllers.MovieController{
			MovieRepo:      movieRepo,
			MovieGenreRepo: movieGenreRepo,
		},
		GenreController: &controllers.GenreController{
			GenreRepo: genreRepo,
		},
	}

	// Set routes
	publicRoutes.Handle(http.MethodGet, "/lifecheck", handlers.LifeCheckHandler())

	publicRoutes.Handle(http.MethodPost, "/auth/register", handlers.RegisterUserHandler(controller.UserController))
	publicRoutes.Handle(http.MethodPost, "/auth/login", handlers.LoginUserHandler(controller.UserController))
	publicRoutes.Handle(http.MethodPost, "/auth/check-email", handlers.CheckEmailHandler(controller.UserController))
	publicRoutes.Handle(http.MethodPost, "/auth/check-username", handlers.CheckUsernameHandler(controller.UserController))

	securedRoutes.Handle(http.MethodGet, "/test", handlers.TestJwtToken)

	router.Handle(http.MethodGet, "/movies", handlers.GetMovies(controller.MovieController))
	router.Handle(http.MethodGet, "/movies/genres", handlers.GetMoviesWithGenres(controller.MovieController))
	router.Handle(http.MethodGet, "/movies/:id", handlers.GetMovieById(controller.MovieController))
	router.Handle(http.MethodGet, "/moviewitheverything/:id", handlers.GetMovieByIdWithEverything(controller.MovieController))
	router.Handle(http.MethodGet, "/movies/:id/genres", handlers.GetMovieByIdWithGenre(controller.MovieController)) // Might not be necessary
	// Will be implemented later
	// router.Handle(http.MethodPost, "/movies", handlers.CreateMovie(controller.MovieController))
	// router.Handle(http.MethodPut, "/movies", handlers.UpdateMovie(controller.MovieController))
	// router.Handle(http.MethodDelete, "/movies/:id", handlers.DeleteMovie(controller.MovieController))
	
	router.Handle(http.MethodGet, "/genres", handlers.GetGenres(controller.GenreController))
	router.Handle(http.MethodGet, "/genres/:name", handlers.GetGenreByName(controller.GenreController))
	router.Handle(http.MethodPost, "/genres/:name", handlers.CreateGenre(controller.GenreController))
	router.Handle(http.MethodGet, "/genres/movies", handlers.GetGenresWithMovies(controller.GenreController))
	router.Handle(http.MethodGet, "/genres/:name/movies", handlers.GetGenreByNameWithMovies(controller.GenreController))


	return router
}
