package models

import "github.com/ELITE-Kinoticketsystem/Backend-KTS/src/myid"

type IdResponse struct {
	Id *myid.UUID `json:"id"`
}
