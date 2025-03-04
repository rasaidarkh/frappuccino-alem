package main

import (
	"database/sql"
	"frappuccino-alem/internal/api"
	"frappuccino-alem/internal/config"
	"frappuccino-alem/pkg/lib/prettyslog"
	"log"
	"net/http"
	"os"

	_ "github.com/lib/pq"
)

func main() {
	// setup config
	cfg := config.LoadConfig()
	// setup logger
	logger := prettyslog.SetupPrettySlog(os.Stdout) // add level based logging

	//create database object
	connStr := cfg.DB.MakeConnectionString()
	db, _ := sql.Open("postgres", connStr)
	// if err != nil {
	// 	log.Fatalf("could not open database:%s", err)
	// }

	//ping database
	_ = db.Ping()
	// if err != nil {
	// 	log.Fatalf("could not ping database:%s", err)
	// }

	// create serve mux
	mux := http.NewServeMux()

	//define api server and start it
	server := api.NewAPIServer(mux, cfg, db, logger)
	logger.Info("running server")
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
