package api

import (
	"database/sql"
	"fmt"
	"frappuccino-alem/internal/config"
	"frappuccino-alem/internal/handlers"
	"frappuccino-alem/internal/handlers/middleware"
	"frappuccino-alem/internal/repository"
	"frappuccino-alem/internal/service"
	"log/slog"
	"net/http"
)

type APIServer struct {
	mux    *http.ServeMux
	cfg    config.Config
	db     *sql.DB
	logger *slog.Logger
}

func NewAPIServer(mux *http.ServeMux, config config.Config, db *sql.DB, logger *slog.Logger) *APIServer {
	return &APIServer{mux, config, db, logger}
}

func (s *APIServer) Run() error {

	// setup three layers for each of the entities
	inventoryRepository := repository.NewInventoryRepository(s.db)
	inventoryService := service.NewInventoryService(inventoryRepository)
	inventoryHandler := handlers.NewInventoryHandler(inventoryService, s.logger)
	inventoryHandler.RegisterEndpoints(s.mux)

	menuRepository := repository.NewMenuRepository(s.db)
	menuService := service.NewMenuService(menuRepository)
	menuHandler := handlers.NewMenuHandler(menuService, s.logger)
	menuHandler.RegisterEndpoints(s.mux)

	orderRepository := repository.NewOrderRepository(s.db)
	orderService := service.NewOrderService(orderRepository)
	orderHandler := handlers.NewOrderHandler(orderService, s.logger)
	orderHandler.RegisterEndpoints(s.mux)

	// add middleware if needed
	timeoutMW := middleware.NewTimoutContextMW(15)
	MWChain := middleware.NewMiddlewareChain(middleware.RecoveryMW, timeoutMW)

	// start server
	serverAddress := fmt.Sprintf("%s:%s", s.cfg.Server.Address, s.cfg.Server.Port)

	return http.ListenAndServe(serverAddress, MWChain(s.mux))
}
