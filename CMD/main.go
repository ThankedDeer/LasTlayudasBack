package main

import (
	"database/sql"
	"github/thankeddeer/lastlayudas/api"
	"github/thankeddeer/lastlayudas/config"
	db "github/thankeddeer/lastlayudas/db/sqlc"
	"log"

	_ "github.com/lib/pq"
)

func main() {

	config, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}
	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(config.ServerAddress)

	if err != nil {
		log.Fatal("cannot start server:", err)
	}

}
