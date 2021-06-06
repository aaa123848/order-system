package repository

import (
	"context"
	"encoding/json"
	"log"
	"order/domain"
	"order/logger"
	"time"

	myerrors "order/errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type SeatsRepository struct {
	Client *mongo.Client
}

func NewSeatsRepository(client *mongo.Client) SeatsRepository {
	return SeatsRepository{
		Client: client,
	}
}

func (sr SeatsRepository) GetASeat(ctx context.Context, filter bson.M) ([]byte, error) {
	col := sr.Client.Database("test").Collection("seats")
	seat := domain.Seat{}
	err := col.FindOne(ctx, filter).Decode(&seat)
	if err != nil {
		logger.LogStdError(err)
		return nil, err
	}
	res, err := json.Marshal(seat)
	if err != nil {
		logger.LogStdError(err)
		return nil, err
	}
	return res, nil
}

func (sr SeatsRepository) UpdateSeats(filter bson.M, update bson.M) error {
	col := sr.Client.Database("test").Collection("seats")
	log.Println(filter)
	res, err := col.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		panic(err)
	}
	log.Println(res)
	if res.ModifiedCount < 1 {
		return myerrors.GetSeatOccupiedError()
	}
	col = sr.Client.Database("test").Collection("transaction")
	col.InsertOne(
		context.TODO(),
		&domain.Transaction{
			Price:      300,
			CreateTime: time.Now(),
		},
	)
	return nil
}

func (sr SeatsRepository) CreateSeat(car string, num int) error {
	s := domain.Seat{
		Position:   domain.Position{Car: car, Num: num},
		Orderby:    "",
		Isorder:    false,
		Updatetime: time.Now(),
	}
	coll := sr.Client.Database("test").Collection("seats")
	_, err := coll.InsertOne(context.TODO(), &s)
	if err != nil {
		return err
	}
	return nil
}
