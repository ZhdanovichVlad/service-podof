package entity

import (
	"time"

	"github.com/ZhdanovichVlad/service-podof/pkg/errorsx"
	"github.com/google/uuid"
)



type Pvz struct {
	Id uuid.UUID `json:"id"`
	RegistrationDate time.Time `json:"registrationDate"`
	City string `json:"city"`
}


type PvzInfo struct {
	Pvz       Pvz                      `json:"pvz"`
	Reception *Reception               `json:"reception,omitempty"`
	Products  []ReceptionWithProducts  `json:"products,omitempty"`
}

type ReceptionWithProducts struct {
    Reception Reception `json:"reception"`
    Products  []Product `json:"products,omitempty"`
}


func (p *Pvz) Validate() error {
	if p.City == "" {
		return errorsx.ErrEmptyField
	}
	if p.RegistrationDate.IsZero() {
		return errorsx.ErrEmptyField
	}
	if p.Id == uuid.Nil {
		return errorsx.ErrEmptyField
	}
	return nil
}
