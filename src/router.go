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
	ReviewController    controllers.ReviewControllerI
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

	reviewsRepo := &repositories.ReviewRepository{
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
		ReviewController: &controllers.ReviewController{
			ReviewRepo: reviewsRepo,
		},
	}

	// Set routes
	publicRoutes.Handle(http.MethodGet, "/lifecheck", handlers.LifeCheckHandler())

	publicRoutes.Handle(http.MethodPost, "/auth/register", handlers.RegisterUserHandler(controller.UserController))
	publicRoutes.Handle(http.MethodPost, "/auth/login", handlers.LoginUserHandler(controller.UserController))
	publicRoutes.Handle(http.MethodPost, "/auth/check-email", handlers.CheckEmailHandler(controller.UserController))
	publicRoutes.Handle(http.MethodPost, "/auth/check-username", handlers.CheckUsernameHandler(controller.UserController))
	publicRoutes.Handle(http.MethodGet, "/auth/logged-in", handlers.LoggedInHandler)

	securedRoutes.Handle(http.MethodGet, "/test", handlers.TestJwtToken)

	publicRoutes.Handle(http.MethodGet, "/movies", handlers.GetMovies(controller.MovieController))
	publicRoutes.Handle(http.MethodGet, "/movies/genres", handlers.GetMoviesWithGenres(controller.MovieController))
	publicRoutes.Handle(http.MethodGet, "/movies/:id", handlers.GetMovieById(controller.MovieController))

	securedRoutes.Handle(http.MethodPost, "/movies", handlers.CreateMovie(controller.MovieController))

	publicRoutes.Handle(http.MethodGet, "/genres", handlers.GetGenres(controller.GenreController))
	publicRoutes.Handle(http.MethodGet, "/genres/:name", handlers.GetGenreByName(controller.GenreController))
	publicRoutes.Handle(http.MethodGet, "/genres/movies", handlers.GetGenresWithMovies(controller.GenreController))
	publicRoutes.Handle(http.MethodPost, "/genres", handlers.CreateGenre(controller.GenreController))
	publicRoutes.Handle(http.MethodPut, "/genres", handlers.UpdateGenre(controller.GenreController))
	publicRoutes.Handle(http.MethodDelete, "/genres/:id", handlers.DeleteGenre(controller.GenreController))

	// Actors
	publicRoutes.Handle(http.MethodGet, "/actors/:id", handlers.GetActorByIdHandler(controller.ActorController))
	publicRoutes.Handle(http.MethodGet, "/actors", handlers.GetActorsHandler(controller.ActorController))
	securedRoutes.Handle(http.MethodPost, "/actors", handlers.CreateActorHandler(controller.ActorController))

	// Price Categories
	publicRoutes.Handle(http.MethodGet, "/price-categories/:id", handlers.GetPriceCategoryByIdHandler(controller.PriceCategories))
	publicRoutes.Handle(http.MethodGet, "/price-categories", handlers.GetPriceCategoriesHandler(controller.PriceCategories))
	securedRoutes.Handle(http.MethodPost, "/price-categories", handlers.CreatePriceCategoryHandler(controller.PriceCategories))
	securedRoutes.Handle(http.MethodPut, "/price-categories/:id", handlers.UpdatePriceCategoryHandler(controller.PriceCategories))
	securedRoutes.Handle(http.MethodDelete, "/price-categories/:id", handlers.DeletePriceCategoryHandler(controller.PriceCategories))

	// event seats
	securedRoutes.Handle(http.MethodGet, "/events/:eventId/seats", handlers.GetEventSeatsHandler(controller.EventSeatController))
	securedRoutes.Handle(http.MethodPatch, "/events/:eventId/seats/:seatId/block", handlers.BlockEventSeatHandler(controller.EventSeatController))
	securedRoutes.Handle(http.MethodPatch, "/events/:eventId/seats/:seatId/unblock", handlers.UnblockEventSeatHandler(controller.EventSeatController))
	securedRoutes.Handle(http.MethodGet, "/events/:eventId/user-seats", handlers.GetSelectedSeatsHandler(controller.EventSeatController))

	// events
	securedRoutes.Handle(http.MethodPost, "/events", handlers.CreateEventHandler(controller.EventController))
	publicRoutes.Handle(http.MethodGet, "/movies/:id/events", handlers.GetEventsForMovieHandler(controller.EventController))
	publicRoutes.Handle(http.MethodGet, "/events/special", handlers.GetSpecialEventsHandler(controller.EventController))

	publicRoutes.Handle(http.MethodPost, "/reviews", handlers.CreateReviewHandler(controller.ReviewController))
	publicRoutes.Handle(http.MethodDelete, "/reviews/:id", handlers.DeleteReviewHandler(controller.ReviewController))

	return router
}
