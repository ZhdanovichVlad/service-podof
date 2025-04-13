package entity

import (
	"time"

	"github.com/ZhdanovichVlad/service-podof/pkg/errorsx"
	"github.com/google/uuid"
)



type Product struct {
	ID          uuid.UUID `json:"id"`
	Type        string    `json:"type"`
	DateTime    time.Time `json:"dateTime"`
	ReceptionID uuid.UUID `json:"receptionId"`
}


func (p *Product) Validate() error {
	if p.Type == "" {
		return errorsx.ErrEmptyField
	}
	
	return nil
}

