package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/shaswath/simplebank/api"
	db "github.com/shaswath/simplebank/db/sqlc"
	"github.com/shaswath/simplebank/utils"
)

func main() {
	config, err := utils.LoadConfig(".")
	if err != nil {
		log.Fatal("Cannot load config file", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)

	if err != nil {
		log.Fatal("Cannot connect to db", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(config.ServerAddress)

	if err != nil {
		log.Fatal("Cannot start server", err)
	}

}
