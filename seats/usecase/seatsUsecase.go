package usecase

import (
	"context"
	"order/domain"

	"go.mongodb.org/mongo-driver/bson"
)

type SeatUseCase struct {
	SeatRepository domain.SeatRepository
}

func NewSeatUseCase(sr domain.SeatRepository) SeatUseCase {
	return SeatUseCase{
		SeatRepository: sr,
	}
}

func (s SeatUseCase) GetASeat(ctx context.Context, filter bson.M) ([]byte, error) {
	b, err := s.SeatRepository.GetASeat(ctx, filter)
	if err != nil {
		return nil, err
	}
	return b, nil
}

func (s SeatUseCase) UpdateSeats(car string, num int, orderer string) error {
	filter := bson.M{"position": bson.M{"car": car, "num": num}, "isorder": false}
	update := bson.M{"$set": bson.M{"isorder": true, "orderby": orderer}}
	return s.SeatRepository.UpdateSeats(filter, update)

}

func (s SeatUseCase) CreateSeat() {
	for i := 0; i < 100; i++ {
		err := s.SeatRepository.CreateSeat("A", i)
		if err != nil {
			panic(err)
		}
	}
}
