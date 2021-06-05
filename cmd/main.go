package main

import (
	"context"

	"order/logger"
	"order/mongodb"
	seatHttp "order/seats/delivery/http"
	seatsMongo "order/seats/repository/mongo"
	seatsUseCase "order/seats/usecase"

	"github.com/gin-gonic/gin"
)

func init() {
	logger.InitLogFile()
}

func main() {
	r := gin.Default()
	url := "mongodb://root:1234@order-mongo-1:27017/?replicaSet=rs1"
	client := mongodb.Connect(context.TODO(), url)
	MongoSeatRepository := seatsMongo.NewSeatsRepository(client)
	SeatsUseCase := seatsUseCase.NewSeatUseCase(MongoSeatRepository)
	seatHttp.NewSeatHandler(r, SeatsUseCase)
	r.Run(":1944")
}
