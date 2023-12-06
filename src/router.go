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
	EventController controllers.EventControllerI
	MovieController controllers.MovieControllerI
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

	eventRepo := &repositories.EventRepository{
		DatabaseManager: databaseManager,
	}

	movieRepo := &repositories.MovieRepository{
		DatabaseManager: databaseManager,
	}

	theatreRepo := &repositories.TheatreRepository{
		DatabaseManager: databaseManager,
	}

	// Create controllers
	controller := Controllers{
		UserController: &controllers.UserController{
			UserRepo: userRepo,
		},
		EventController: &controllers.EventController{
			EventRepo:   eventRepo,
			TheatreRepo: theatreRepo,
		},
		MovieController: &controllers.MovieController{
			MovieRepo: movieRepo,
		},
	}

	// Set routes
	publicRoutes.Handle(http.MethodGet, "/lifecheck", handlers.LifeCheckHandler())

	publicRoutes.Handle(http.MethodPost, "/auth/register", handlers.RegisterUserHandler(controller.UserController))
	publicRoutes.Handle(http.MethodPost, "/auth/login", handlers.LoginUserHandler(controller.UserController))
	publicRoutes.Handle(http.MethodPost, "/auth/check-email", handlers.CheckEmailHandler(controller.UserController))
	publicRoutes.Handle(http.MethodPost, "/auth/check-username", handlers.CheckUsernameHandler(controller.UserController))

	securedRoutes.Handle(http.MethodGet, "/test", handlers.TestJwtToken)

	router.Handle(http.MethodPost, "/events", handlers.CreateEventHandler(controller.EventController))
	router.Handle(http.MethodDelete, "/events/:id", handlers.DeleteEventHandler(controller.EventController))
	// Get events for movieId
	router.Handle(http.MethodGet, "/events/movies/:id", handlers.GetEventsForMovieHandler(controller.EventController))
	router.Handle(http.MethodGet, "/events/special-events", handlers.GetSpecialEventsHandler(controller.EventController))
	// TODO: Do we need to add update event handler because how would we proceed then?

	router.Handle(http.MethodGet, "/movies", handlers.GetMovies(controller.MovieController))
	router.Handle(http.MethodGet, "/movies/:id", handlers.GetMovieById(controller.MovieController))
	router.Handle(http.MethodPost, "/movies", handlers.CreateMovie(controller.MovieController))
	router.Handle(http.MethodPut, "/movies", handlers.UpdateMovie(controller.MovieController))
	router.Handle(http.MethodDelete, "/movies/:id", handlers.DeleteMovie(controller.MovieController))

	router.Handle(http.MethodGet, "/genres", handlers.GetGenres(controller.MovieController))
	router.Handle(http.MethodGet, "/genres/:name", handlers.GetGenreByName(controller.MovieController))
	router.Handle(http.MethodPost, "/genres", handlers.CreateGenre(controller.MovieController))

	router.Handle(http.MethodGet, "/movies/:id/genres", handlers.GetMovieByIdWithGenre(controller.MovieController))
	router.Handle(http.MethodGet, "/genres/:name/movies", handlers.GetGenreByNameWithMovies(controller.MovieController))
	router.Handle(http.MethodGet, "/genres/movies", handlers.GetGenresWithMovies(controller.MovieController))
	router.Handle(http.MethodGet, "/movies/genres", handlers.GetMoviesWithGenres(controller.MovieController))

	return router
}
