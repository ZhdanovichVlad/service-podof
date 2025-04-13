package entity

import (
	"time"

	"github.com/ZhdanovichVlad/service-podof/pkg/errorsx"
)

const (
	defaultLimit  = 10
	defaultOffset = 0
	maxLimit      = 30
)

type Filter struct {
	StartDate *time.Time `form:"start_date"`
	EndDate   *time.Time `form:"end_date"`
	Limit     *int       `form:"limit"`
	Offset     *int      `form:"page"`
}

func (f *Filter) Validate() error {
 
    if f.StartDate != nil && f.EndDate != nil {
        if f.StartDate.After(*f.EndDate) {
            return errorsx.ErrStartDateAfterEndDate
        }
    } else {
        endDate := time.Now()
        f.EndDate = &endDate
        startDate := endDate.AddDate(0, 0, -30)
        f.StartDate = &startDate
    }


    if f.Limit == nil {
        limit := defaultLimit
        f.Limit = &limit
    } else if *f.Limit <= 0 {
        return errorsx.ErrInvalidLimit
    } else if *f.Limit > maxLimit {
        return errorsx.ErrInvalidLimit
    }

    if f.Offset == nil {
        offset := defaultOffset
        f.Offset = &offset
    } else if *f.Offset <= 0 {
      
        page := 1
        f.Offset = &page
    }

  
    offset := (*f.Offset - 1) * *f.Limit
    f.Offset = &offset

    return nil
}
