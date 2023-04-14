package main

import (
	"context"
	"database/sql"
	"fmt"

	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	_ "github.com/go-sql-driver/mysql"
	"github.com/marcelofelixsalgado/fullcycle-eda/walletcore/internal/database"
	"github.com/marcelofelixsalgado/fullcycle-eda/walletcore/internal/event"
	"github.com/marcelofelixsalgado/fullcycle-eda/walletcore/internal/event/handler"
	createaccount "github.com/marcelofelixsalgado/fullcycle-eda/walletcore/internal/usecase/create_account"
	"github.com/marcelofelixsalgado/fullcycle-eda/walletcore/internal/usecase/create_client"
	"github.com/marcelofelixsalgado/fullcycle-eda/walletcore/internal/usecase/create_transaction"
	"github.com/marcelofelixsalgado/fullcycle-eda/walletcore/internal/web"
	"github.com/marcelofelixsalgado/fullcycle-eda/walletcore/internal/web/webserver"
	"github.com/marcelofelixsalgado/fullcycle-eda/walletcore/pkg/events"
	"github.com/marcelofelixsalgado/fullcycle-eda/walletcore/pkg/kafka"
	"github.com/marcelofelixsalgado/fullcycle-eda/walletcore/pkg/uow"
)

func main() {
	databaseConnectionUser := "root"
	databaseConnectionPassword := "root"
	databaseConnectionServerAddress := "mysql-walletcore"
	databaseConnectionServerPort := 3306
	databaseName := "wallet"

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		databaseConnectionUser,
		databaseConnectionPassword,
		databaseConnectionServerAddress,
		databaseConnectionServerPort,
		databaseName)

	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	configMap := ckafka.ConfigMap{
		"bootstrap.servers": "kafka:9094",
		"group.id":          "wallet",
	}
	kafkaProducer := kafka.NewKafkaProducer(&configMap)

	eventDispatcher := events.NewEventDispatcher()
	eventDispatcher.Register("TransactionCreated", handler.NewTransactionCreatedKafkaHandler(kafkaProducer))
	eventDispatcher.Register("BalanceUpdated", handler.NewUpdateBalanceKafkaHandler(kafkaProducer))
	transactionCreatedEvent := event.NewTransactionCreated()
	balanceUpdatedEvent := event.NewBalanceUpdated()

	clientDb := database.NewClientDB(db)
	accountDb := database.NewAccountDB(db)

	ctx := context.Background()
	uow := uow.NewUow(ctx, db)

	uow.Register("AccountDB", func(tx *sql.Tx) interface{} {
		return database.NewAccountDB(db)
	})

	uow.Register("TransactionDB", func(tx *sql.Tx) interface{} {
		return database.NewTransactionDB(db)
	})
	createTransactionUseCase := create_transaction.NewCreateTransactionUseCase(uow, eventDispatcher, transactionCreatedEvent, balanceUpdatedEvent)
	createClientUseCase := create_client.NewCreateClientUseCase(clientDb)
	createAccountUseCase := createaccount.NewCreateAccountUseCase(accountDb, clientDb)

	webserver := webserver.NewWebServer(":3002")

	clientHandler := web.NewWebClientHandler(*createClientUseCase)
	accountHandler := web.NewWebAccountHandler(*createAccountUseCase)
	transactionHandler := web.NewWebTransactionHandler(*createTransactionUseCase)

	webserver.AddHandler("/clients", clientHandler.CreateClient)
	webserver.AddHandler("/accounts", accountHandler.CreateAccount)
	webserver.AddHandler("/transactions", transactionHandler.CreateTransaction)

	fmt.Println("Server is running")
	webserver.Start()
}
