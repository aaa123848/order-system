package main

import (
	"context"

	"order/logger"
	"order/mongodb"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"

	seatHttp "order/seats/delivery/http"
	seatsMongo "order/seats/repository/mongo"
	seatsUseCase "order/seats/usecase"
)

func init() {
	logger.InitLogFile()
}

var Coll *mongo.Collection

func main() {
	r := gin.Default()
	url := "mongodb://root:1234@order-mongo-1:27017/?replicaSet=rs1"
	client := mongodb.Connect(context.TODO(), url)
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()
	MongoSeatRepository := seatsMongo.NewSeatsRepository(client)
	SeatsUseCase := seatsUseCase.NewSeatUseCase(MongoSeatRepository)
	seatHttp.NewSeatHandler(r, SeatsUseCase)
	r.Run(":1944")
}
