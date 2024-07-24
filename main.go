package main

import (
	"database/sql"
	"github.com/m-bashari-m/simplebank/api"
	db "github.com/m-bashari-m/simplebank/db/sqlc"
	"github.com/m-bashari-m/simplebank/utils"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	config, err := utils.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot read config:", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(config.Address)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}
