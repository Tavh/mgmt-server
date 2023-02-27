package main

import (
	"database/sql"
	"mgmt-server/api"
	"mgmt-server/dal"

	_ "github.com/lib/pq"
)

func main() {
	connStr := "postgres://postgres:postgres@localhost:5431/chat_app_mgmt?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}

	userDAL := dal.NewUserDAL(db)
	chatRoomDAL := dal.NewChatRoomDAL(db)
	if err := api.NewAPI(userDAL, chatRoomDAL).Start(1000); err != nil {
		panic(err)
	}
}
