package main

import (
	"database/sql"
	"encoding/json"
	"fmt"

	"github.com/marcelofelixsalgado/fullcycle-eda/balances/internal/database"
	getbalance "github.com/marcelofelixsalgado/fullcycle-eda/balances/internal/usecase/get_balance"
	updatebalance "github.com/marcelofelixsalgado/fullcycle-eda/balances/internal/usecase/update_balance"
	"github.com/marcelofelixsalgado/fullcycle-eda/balances/internal/web"
	"github.com/marcelofelixsalgado/fullcycle-eda/balances/internal/web/webserver"
	"github.com/marcelofelixsalgado/fullcycle-eda/balances/pkg/kafka"

	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	_ "github.com/go-sql-driver/mysql"
)

type BalancePayloadInputDTO struct {
	Payload BalanceInputDTO `json:"Payload"`
}

type BalanceInputDTO struct {
	AccountIDFrom        string  `json:"account_id_from"`
	AccountIDTo          string  `json:"account_id_to"`
	BalanceAccountIDFrom float64 `json:"balance_account_id_from"`
	BalanceAccountIDTo   float64 `json:"balance_account_id_to"`
}

func main() {

	databaseConnectionUser := "root"
	databaseConnectionPassword := "root"
	databaseConnectionServerAddress := "mysql-balances"
	databaseConnectionServerPort := 3306
	databaseName := "balances"

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

	balanceDb := database.NewBalanceDB(db)
	updateBalanceUseCase := updatebalance.NewUpdateBalanceUseCase(balanceDb)
	getBalanceUseCase := getbalance.NewGetBalanceUseCase(balanceDb)

	webserver := webserver.NewWebServer(":3003")
	balanceHandler := web.NewWebBalanceHandler(*getBalanceUseCase)
	webserver.AddHandler("/balances/{id}", balanceHandler.GetBalance)
	go webserver.Start()

	fmt.Println("Server is started...")

	topics := []string{"balances"}
	configMap := ckafka.ConfigMap{
		"bootstrap.servers": "kafka:29092",
		"group.id":          "wallet",
		"auto.offset.reset": "earliest",
	}
	kafkaConsumer := kafka.NewConsumer(&configMap, topics)

	balanceChan := make(chan *ckafka.Message)
	go kafkaConsumer.Consume(balanceChan)

	for msg := range balanceChan {

		balancePayloadInputDTO := BalancePayloadInputDTO{}
		err := json.Unmarshal(msg.Value, &balancePayloadInputDTO)
		if err != nil {
			fmt.Println("error: ", err)
			return
		}

		updateBalanceInputFromDTO := updatebalance.UpdateBalanceInputDTO{
			AccountId: balancePayloadInputDTO.Payload.AccountIDFrom,
			Amount:    balancePayloadInputDTO.Payload.BalanceAccountIDFrom,
		}
		updateBalanceOutputFromDTO, err := updateBalanceUseCase.Execute(updateBalanceInputFromDTO)
		if err != nil {
			fmt.Println("error: ", err)
			return
		}
		fmt.Println(updateBalanceOutputFromDTO)

		updateBalanceInputToDTO := updatebalance.UpdateBalanceInputDTO{
			AccountId: balancePayloadInputDTO.Payload.AccountIDTo,
			Amount:    balancePayloadInputDTO.Payload.BalanceAccountIDTo,
		}
		updateBalanceOutputToDTO, err := updateBalanceUseCase.Execute(updateBalanceInputToDTO)
		if err != nil {
			fmt.Println("error: ", err)
			return
		}
		fmt.Println(updateBalanceOutputToDTO)
	}
	fmt.Println("Server is stoping...")
}
