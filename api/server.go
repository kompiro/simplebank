package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	db "github.com/techschool/simplebank/db/sqlc"
	"github.com/techschool/simplebank/token"
	"github.com/techschool/simplebank/util"
)

// Server serves HTTP requests for our banking service.
type Server struct {
	config     util.Config
	store      db.Store
	tokenMaker token.Maker
	router     *gin.Engine
}

// NewServer creates a new HTTP server and setup routing.
func NewServer(config util.Config, store db.Store) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}

	server := &Server{
		config:     config,
		store:      store,
		tokenMaker: tokenMaker,
	}

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("currency", validCurrency)
	}

	server.setupRouter()
	return server, nil
}

func (server *Server) setupRouter() {
	router := gin.Default()

	router.GET("/", server.root)

	router.GET("/healthz", server.healthCheck)

	router.POST("/users", server.CreateUser)
	router.POST("/users/login", server.loginUser)
	router.POST("/tokens/renew_access", server.renewAccessToken)

	authRouters := router.Group("/").Use(authMiddleware(server.tokenMaker))

	authRouters.POST("/accounts", server.CreateAccount)
	authRouters.GET("/accounts/:account_id", server.GetAccount)
	authRouters.GET("/accounts", server.ListAccount)

	authRouters.GET("/entries/:id", server.GetEntry)
	authRouters.GET("/accounts/:account_id/entries", server.ListEntry)

	authRouters.POST("/transfers", server.CreateTransfer)

	server.router = router
}

func (server *Server) root(ctx *gin.Context) {
	ctx.String(200, "Welcome!")
}

// Start runs the HTTP server on a specific address.
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
