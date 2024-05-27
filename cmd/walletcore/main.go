package main

import (
	"database/sql"
	"fmt"

	"github.com/caiofsr/walletcore/internal/database"
	"github.com/caiofsr/walletcore/internal/event"
	"github.com/caiofsr/walletcore/internal/usecases"
	"github.com/caiofsr/walletcore/internal/web"
	"github.com/caiofsr/walletcore/internal/web/webserver"
	"github.com/caiofsr/walletcore/pkg/events"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", "root", "root", "localhost", "3306", "wallet"))
	if err != nil {
		panic(err)
	}
	defer db.Close()

	eventDispatcher := events.NewEventDispatcher()
	transferCreatedEvent := event.NewTransferCreated()
	// eventDispatcher.Register("TransferCreated", handler)

	clientDb := database.NewClientDB(db)
	accountDb := database.NewAccountDB(db)
	transferDb := database.NewTransferDB(db)

	createClientUseCase := usecases.NewCreateClientUseCase(clientDb)
	createAccountUseCase := usecases.NewCreateAccountUseCase(accountDb, clientDb)
	createTransferUseCase := usecases.NewCreateTransferUseCase(transferDb, accountDb, eventDispatcher, transferCreatedEvent)

	webserver := webserver.NewWebServer(":3333")

	clientHandler := web.NewWebClientHandler(*createClientUseCase)
	accountHandler := web.NewWebAccountHandler(*createAccountUseCase)
	transferHandler := web.NewWebTransferHandler(*createTransferUseCase)

	webserver.AddHandler("/clients", clientHandler.CreateClient)
	webserver.AddHandler("/accounts", accountHandler.CreateAccount)
	webserver.AddHandler("/transfers", transferHandler.CreateTransfer)

	webserver.Start()
}
