package domain

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

type Seat struct {
	Position   Position
	Orderby    string
	Isorder    bool
	Updatetime time.Time
}

type Position struct {
	Car string
	Num int
}

type SeatRepository interface {
	GetASeat(ctx context.Context, filter bson.M) ([]byte, error)
	CreateSeat(car string, num int) error
	UpdateSeats(filter bson.M, update bson.M) error
}

type SeatUseCase interface {
	GetASeat(ctx context.Context, filter bson.M) ([]byte, error)
	CreateSeat()
	UpdateSeats(car string, num int, oderer string) error
}
