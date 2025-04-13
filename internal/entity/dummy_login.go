package entity

import "github.com/ZhdanovichVlad/service-podof/pkg/errorsx"

type DummyLogin struct {
	Role string `json:"role" binding:"required"`
}


func (d *DummyLogin) Validate() error {
	if d.Role == "" {
		return errorsx.ErrEmptyField
	}
	return nil
}

