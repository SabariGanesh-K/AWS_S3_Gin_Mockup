package api

import (
	"fmt"

	db "github.com/SabariGanesh-K/21BPS1209_Backend.git/db/sqlc"
	"github.com/SabariGanesh-K/21BPS1209_Backend.git/token"
	"github.com/SabariGanesh-K/21BPS1209_Backend.git/util"
	"github.com/gin-gonic/gin"
	// "github.com/gin-gonic/gin/binding"
	// "github.com/go-playground/validator/v10"
)

type Server struct {
	config     util.Config
	store      db.Store
	tokenMaker token.Maker
	router     *gin.Engine
}

func NewServer(config util.Config, store db.Store) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token: %v", err)
	}
	server := &Server{
		config: config,
		store: store,
		 tokenMaker: tokenMaker}
	// if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
	// 	v.RegisterValidation("currency", validCurrency)
	// }
	// //
	server.setupRouter()
	return server, nil

	
}
func (server *Server) setupRouter() {
	router := gin.Default()
	router.POST("/users", server.createUser)
	router.POST("/users/login", server.loginUser)
	router.POST("/tokens/renew_access",server.renewAccessToken)
    router.POST("/file/one", server.uploadFile)
	// authRoutes := router.Group("/").Use(authMiddleware(server.tokenMaker))
	// authRoutes.POST("/accounts", server.createAccount)
	// authRoutes.GET("/accounts/:id", server.getAccount)

	// authRoutes.POST("/transfers", server.createTransfer)

	server.router = router

}

func (server *Server) Start(address string) error {
	// err := cron.startBackgroundJob(ctx, server.config, server.store)
	return server.router.Run(address)
}
func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
