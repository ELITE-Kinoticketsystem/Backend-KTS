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
	UserController      controllers.UserControllerI
	EventController     controllers.EventControllerI
	ActorController     controllers.ActorControllerI
	MovieController     controllers.MovieControllerI
	EventSeatController controllers.EventSeatControllerI
	GenreController     controllers.GenreControllerI
	PriceCategories     controllers.PriceCategoryControllerI
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

	movieActorRepo := &repositories.MovieActorRepository{
		DatabaseManager: databaseManager,
	}

	movieProducerRepo := &repositories.MovieProducerRepository{
		DatabaseManager: databaseManager,
	}

	actorRepo := &repositories.ActorRepository{
		DatabaseManager: databaseManager,
	}

	priceCategoryRepo := &repositories.PriceCategoryRepository{
		DatabaseManager: databaseManager,
	}

	eventSeatRepo := &repositories.EventSeatRepository{
		DatabaseManager: databaseManager,
	}

	// Create controllers
	controller := Controllers{
		UserController: &controllers.UserController{
			UserRepo: userRepo,
		},
		MovieController: &controllers.MovieController{
			MovieRepo:         movieRepo,
			MovieGenreRepo:    movieGenreRepo,
			MovieActorRepo:    movieActorRepo,
			MovieProducerRepo: movieProducerRepo,
		},
		GenreController: &controllers.GenreController{
			GenreRepo: genreRepo,
		},
		ActorController: &controllers.ActorController{
			ActorRepo: actorRepo,
		},
		PriceCategories: &controllers.PriceCategoryController{
			PriceCategoryRepository: priceCategoryRepo,
		},
		EventSeatController: &controllers.EventSeatController{
			EventSeatRepo: eventSeatRepo,
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

	// Will be implemented later
	router.Handle(http.MethodPost, "/movies", handlers.CreateMovie(controller.MovieController))
	// router.Handle(http.MethodPut, "/movies", handlers.UpdateMovie(controller.MovieController))
	// router.Handle(http.MethodDelete, "/movies/:id", handlers.DeleteMovie(controller.MovieController))

	router.Handle(http.MethodGet, "/genres", handlers.GetGenres(controller.GenreController))
	router.Handle(http.MethodGet, "/genres/:name", handlers.GetGenreByName(controller.GenreController))
	router.Handle(http.MethodGet, "/genres/movies", handlers.GetGenresWithMovies(controller.GenreController))
	router.Handle(http.MethodPost, "/genres", handlers.CreateGenre(controller.GenreController))
	router.Handle(http.MethodPut, "/genres", handlers.UpdateGenre(controller.GenreController))
	router.Handle(http.MethodDelete, "/genres/:id", handlers.DeleteGenre(controller.GenreController))

	// Actors
	router.Handle(http.MethodGet, "/actors/:id", handlers.GetActorByIdHandler(controller.ActorController))
	router.Handle(http.MethodGet, "/actors", handlers.GetActorsHandler(controller.ActorController))
	router.Handle(http.MethodPost, "/actors", handlers.CreateActorHandler(controller.ActorController))

	// Price Categories
	router.Handle(http.MethodGet, "/price-categories/:id", handlers.GetPriceCategoryByIdHandler(controller.PriceCategories))
	router.Handle(http.MethodGet, "/price-categories", handlers.GetPriceCategoriesHandler(controller.PriceCategories))
	router.Handle(http.MethodPost, "/price-categories", handlers.CreatePriceCategoryHandler(controller.PriceCategories))
	router.Handle(http.MethodPut, "/price-categories/:id", handlers.UpdatePriceCategoryHandler(controller.PriceCategories))
	router.Handle(http.MethodDelete, "/price-categories/:id", handlers.DeletePriceCategoryHandler(controller.PriceCategories))

	// event seats
	router.Handle(http.MethodGet, "/events/:eventId/seats", handlers.GetEventSeatsHandler(controller.EventSeatController))
	router.Handle(http.MethodPatch, "/events/:eventId/seats/:seatId/block", handlers.BlockEventSeatHandler(controller.EventSeatController))
	router.Handle(http.MethodPatch, "/events/:eventId/seats/:seatId/unblock", handlers.UnblockEventSeatHandler(controller.EventSeatController))
	router.Handle(http.MethodGet, "/events/:eventId/user-seats", handlers.GetSelectedSeatsHandler(controller.EventSeatController))

	// events
	router.Handle(http.MethodPost, "/events", handlers.CreateEventHandler(controller.EventController))
	router.Handle(http.MethodGet, "/movies/:movieId/events", handlers.GetEventsForMovieHandler(controller.EventController))
	router.Handle(http.MethodGet, "/events/special", handlers.GetSpecialEventsHandler(controller.EventController))

	return router
}
