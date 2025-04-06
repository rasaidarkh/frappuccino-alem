package api

import (
	"database/sql"
	"fmt"
	"log/slog"
	"net/http"

	"frappuccino-alem/internal/config"
	"frappuccino-alem/internal/handlers"
	"frappuccino-alem/internal/handlers/middleware"
	"frappuccino-alem/internal/service"
	"frappuccino-alem/internal/store"
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
	inventoryStore := store.NewInventoryStore(s.db)
	inventoryService := service.NewInventoryService(inventoryStore)
	inventoryHandler := handlers.NewInventoryHandler(inventoryService, s.logger)
	inventoryHandler.RegisterEndpoints(s.mux)

	menuStore := store.NewMenuStore(s.db)
	menuService := service.NewMenuService(menuStore, inventoryStore)
	menuHandler := handlers.NewMenuHandler(menuService, s.logger)
	menuHandler.RegisterEndpoints(s.mux)

	orderStore := store.NewOrderStore(s.db)
	orderService := service.NewOrderService(inventoryStore, menuStore, orderStore)
	orderHandler := handlers.NewOrderHandler(orderService, s.logger)
	orderHandler.RegisterEndpoints(s.mux)

	reportStore := store.NewReportStore(s.db)
	reportService := service.NewReportService(reportStore)
	reportHandler := handlers.NewReportHandler(reportService, s.logger)
	reportHandler.RegisterEndpoints(s.mux)

	// add middleware if needed
	timeoutMW := middleware.NewTimoutContextMW(15)
	// WholeMwChain
	MWChain := middleware.NewMiddlewareChain(middleware.RecoveryMW, timeoutMW)

	// start server
	serverAddress := fmt.Sprintf("%s:%s", s.cfg.Server.Address, s.cfg.Server.Port)

	s.logger.Info("starting server", slog.String("host", serverAddress))
	return http.ListenAndServe(serverAddress, MWChain(s.mux))
}
