package main

import (
	"Tech_Blog/api"
	db "Tech_Blog/db/sqlc"
	"Tech_Blog/util"
	"database/sql"

	_ "github.com/lib/pq"

	"log"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("connot laod config:", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("Connot connect to the database:", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)
	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("connot start server:", err)
	}

}
