package samples

import (
	"time"

	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/.gen/KinoTicketSystem/model"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/models"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/utils"
	"github.com/google/uuid"
)

func GetGetSpecialEventsDTO(eventId *uuid.UUID) *models.GetSpecialEventsDTO {
	rating := 4.5
	return &models.GetSpecialEventsDTO{
		Events: model.Events{
			ID:           eventId,
			Title:        "My Events",
			Start:        time.Now(),
			End:          time.Now().Add(time.Hour),
			Description:  utils.GetStringPointer("MyDescription"),
			EventType:    "special event",
			CinemaHallID: utils.NewUUID(),
		},
		Movies: []*model.Movies{
			{
				ID:           utils.NewUUID(),
				Title:        "The Godfather",
				Description:  "The Godfather \"Don\" Vito Corleone is the head of the Corleone mafia family in New York. He is at the event of his daughter's wedding. Michael, Vito's youngest son and a decorated WW II Marine is also present at the wedding. Michael seems to be uninterested in being a part of the family business. Vito is a powerful man, and is kind to all those who give him respect but is ruthless against those who do not. But when a powerful and treacherous rival wants to sell drugs and needs the Don's influence for the same, Vito refuses to do it. What follows is a clash between Vito's fading old values and the new ways which may cause Michael to do the thing he was most reluctant in doing and wage a mob war against all the other mafia families which could tear the Corleone family apart.",
				BannerPicURL: utils.GetStringPointer("https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcREKr_g8JfVvWEiTS_17e4p8o-0RQHc9ydl6xlhSM9-vCBUXumK-b3bNPv0LRM_aTGacQ&usqp=CAU"),
				CoverPicURL:  utils.GetStringPointer("https://upload.wikimedia.org/wikipedia/en/1/1c/Godfather_ver1.jpg"),
				TrailerURL:   utils.GetStringPointer("UaVTIH8mujA"),
				Rating:       &rating,
				ReleaseDate:  time.Now(),
				TimeInMin:    175,
				Fsk:          16,
			},
			{
				ID:           utils.NewUUID(),
				Title:        "The Dark Knight",
				Description:  "When the menace known as the Joker wreaks havoc and chaos on the people of Gotham, Batman must accept one of the greatest psychological and physical tests of his ability to fight injustice.",
				BannerPicURL: utils.GetStringPointer("https://images.moviesanywhere.com/9d3294ee14822bbb81e24bb0a9610b98/d030d491-7a4a-42aa-832f-b2fd25d5b1f8.jpg?r=3x1&w=2400"),
				CoverPicURL:  utils.GetStringPointer("https://images.moviesanywhere.com/bd47f9b7d090170d79b3085804075d41/c6140695-a35f-46e2-adb7-45ed829fc0c0.webp?h=375&resize=fit&w=250"),
				TrailerURL:   utils.GetStringPointer("_64S_ixM5Ng"),
				Rating:       &rating,
				ReleaseDate:  time.Now(),
				TimeInMin:    152,
				Fsk:          12,
			},
		},
	}
}

func GetModelEvents() []*model.Events {
	return []*model.Events{{
		ID:           utils.NewUUID(),
		Title:        "Test Event 1",
		Start:        time.Now(),
		End:          time.Now().Add(time.Hour),
		Description:  nil,
		EventType:    "Test event type 1",
		CinemaHallID: utils.NewUUID(),
	},
		{
			ID:           utils.NewUUID(),
			Title:        "Test Event 2",
			Start:        time.Now(),
			End:          time.Now().Add(time.Hour),
			Description:  nil,
			EventType:    "Test event type 2",
			CinemaHallID: utils.NewUUID(),
		},
	}
}
