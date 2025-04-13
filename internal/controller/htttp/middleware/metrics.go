package middleware

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func (m *Middleware) RequestCounter() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		m.metrics.IncRequestsTotal(c.Request.Method, c.Request.URL.Path, strconv.Itoa(c.Writer.Status()))
	}
}

func (m *Middleware) ResponseLatency() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()
		latency := time.Since(start).Milliseconds()
		m.metrics.ResponseLatency(c.Request.Method, c.Request.URL.Path, float64(latency))
	}
}

func (m *Middleware) IncCreatedPvzs() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		if c.Writer.Status() == http.StatusCreated {	
			m.metrics.IncCreatedPvzs(strconv.Itoa(c.Writer.Status()))
		}
	}
}

func (m *Middleware) IncCreatedReceptions() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		if c.Writer.Status() == http.StatusCreated {
			m.metrics.IncCreatedReceptions(strconv.Itoa(c.Writer.Status()))
		}
	}
}

func (m *Middleware) IncAddedProducts() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		if c.Writer.Status() == http.StatusCreated {
			m.metrics.IncAddedProducts(strconv.Itoa(c.Writer.Status()))
		}
	}
}

