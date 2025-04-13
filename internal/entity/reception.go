package entity

import (
	"time"

	"github.com/google/uuid"
	"github.com/ZhdanovichVlad/service-podof/pkg/errorsx"

)


const (
	 ReceptionStatusInProgress = "in_progress"
	 ReceptionStatusCompleted  = "closed"
)


type Reception struct {
	Id            uuid.UUID       `json:"id"`
	DateTime      time.Time       `json:"dateTime"`
	PvzID         uuid.UUID       `json:"pvzId,omitempty"`
	Status        string 	 	  `json:"status"`
}


func (r *Reception) Validate() error {
	if r.PvzID == uuid.Nil {
		return errorsx.ErrEmptyField
	}
	return nil
}





