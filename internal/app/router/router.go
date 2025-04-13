package router

import (
	"log/slog"
	"net/http"

	"github.com/ZhdanovichVlad/service-podof/internal/controller/htttp/middleware"
	"github.com/ZhdanovichVlad/service-podof/internal/controller/htttp/handler"
	"github.com/ZhdanovichVlad/service-podof/pkg/errorsx"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

const (
   
    dummyLoginPath = "/dummy-login"
    registerPath   = "/register"
    loginPath      = "/login"

    getPvzListPath = "/pvz/"
    createPvzPath  = "/pvz/"

 
    createReceptionPath = "/reception"
    closeReceptionPath  = "/pvz/:pvzId/close_last_reception"


    
    addProduct     = "/products"
    deleteProduct  = "/pvz/:pvzId/delete_last_product"
)


type Router struct {
    router     *gin.Engine
    handlers   *handler.Handler
    middleware *middleware.Middleware
	logger     *slog.Logger
    addr       string
}

func NewRouter(h *handler.Handler, m *middleware.Middleware, logger *slog.Logger, addr string) *Router {
    router := gin.New()
    
   
    router.Use(gin.Recovery())
    router.Use(gin.Logger())

    return &Router{
        router:     router,
        handlers:   h,
        middleware: m,
        logger:     logger,
        addr:       addr,
    }
}

func (r *Router) Run() error {
	return r.router.Run(r.addr)
}

func (r *Router) SetupRoutes() {
    
    r.router.GET("/metrics", gin.WrapH(promhttp.Handler()))
	
	r.router.NoRoute(func(c *gin.Context) {
        c.JSON(http.StatusNotFound, gin.H{
            "message": errorsx.ErrNotFound.Error(),
        })
    })

    
    r.router.Use(r.middleware.RequestCounter())
    r.router.Use(r.middleware.ResponseLatency())

  
    r.router.POST(dummyLoginPath, r.handlers.DummyLogin)
    r.router.POST(registerPath, r.handlers.Register)
    r.router.POST(loginPath, r.handlers.Login)

    
    authorized := r.router.Group("/")
    authorized.Use(r.middleware.AuthMiddleware())
    {

        authorized.GET(getPvzListPath, r.handlers.GetPvzList)
        authorized.POST(createPvzPath, 
            r.middleware.ModeratorVerification(), 
            r.middleware.IncCreatedPvzs(), 
            r.handlers.CreatePvz)

      
        authorized.POST(createReceptionPath, 
            r.middleware.PvzEmployeeVerification(), 
            r.middleware.IncCreatedReceptions(), 
            r.handlers.CreateReception)
        authorized.POST(closeReceptionPath, 
            r.middleware.PvzEmployeeVerification(), 
            r.handlers.CloseReception)

       
        authorized.POST(addProduct, 
            r.middleware.PvzEmployeeVerification(), 
            r.middleware.IncAddedProducts(), 
            r.handlers.CreateProduct)
        authorized.POST(deleteProduct, 
            r.middleware.PvzEmployeeVerification(), 
            r.handlers.DeleteLastProduct)
    }
}