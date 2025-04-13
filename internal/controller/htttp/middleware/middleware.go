package middleware

import "github.com/ZhdanovichVlad/service-podof/pkg/jwttoken"

type ITokenValidator interface {
	ValidateToken(tokenString string) (*jwttoken.Claims, error)
}

type Metrics interface {
	IncRequestsTotal(method, path, status string)
	ResponseLatency(method, path string, milliseconds float64)
	IncCreatedPvzs(status string)
	IncCreatedReceptions(status string)
	IncAddedProducts(status string)
}

type Middleware struct {
	tokenValidator ITokenValidator
	metrics Metrics
}

func NewMiddleware(tokenValidator ITokenValidator, metrics Metrics) *Middleware {
	return &Middleware{tokenValidator: tokenValidator, metrics: metrics}
}


