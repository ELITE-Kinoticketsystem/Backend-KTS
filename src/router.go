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
	ActorController controllers.ActorControllerI
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

	actorRepo := &repositories.ActorRepository{
		DatabaseManager: databaseManager,
	}

	// Create controllers
	controller := Controllers{
		UserController: &controllers.UserController{
			UserRepo: userRepo,
		},
		EventController: &controllers.EventController{
			EventRepo:   eventRepo,
			MovieRepo:   movieRepo,
			TheatreRepo: theatreRepo,
		},
		ActorController: &controllers.ActorController{
			ActorRepo: actorRepo,
		},
	}

	// Set routes
	publicRoutes.Handle(http.MethodGet, "/lifecheck", handlers.LifeCheckHandler())

	publicRoutes.Handle(http.MethodPost, "/auth/register", handlers.RegisterUserHandler(controller.UserController))
	publicRoutes.Handle(http.MethodPost, "/auth/login", handlers.LoginUserHandler(controller.UserController))
	publicRoutes.Handle(http.MethodPost, "/auth/check-email", handlers.CheckEmailHandler(controller.UserController))
	publicRoutes.Handle(http.MethodPost, "/auth/check-username", handlers.CheckUsernameHandler(controller.UserController))

	securedRoutes.Handle(http.MethodGet, "/test", handlers.TestJwtToken)

	securedRoutes.Handle(http.MethodPost, "/events", handlers.CreateEventHandler(controller.EventController))

	// Get events for movieId
	publicRoutes.Handle(http.MethodGet, "/events/movies/:id", handlers.GetEventsForMovieHandler(controller.EventController))
	publicRoutes.Handle(http.MethodGet, "/events/special-events", handlers.GetSpecialEventsHandler(controller.EventController))
	// TODO: Do we need to add update event handler because how would we proceed then?

	// Should be only accessible for admins
	securedRoutes.Handle(http.MethodDelete, "/events/:id", handlers.DeleteEventHandler(controller.EventController))

	// Actors
	router.Handle(http.MethodGet, "/actors/:id", handlers.GetActorByIdHandler(controller.ActorController))
	router.Handle(http.MethodGet, "/actors/", handlers.GetActorsHandler(controller.ActorController))

	return router
}
