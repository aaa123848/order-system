package usecase

import (
	"context"
	"encoding/json"
	"order/domain"
	"testing"
	"time"

	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson"
)

type MockS struct {
	M mock.Mock
}

func (m *MockS) GetASeat(ctx context.Context, filter bson.M) ([]byte, error) {
	args := m.M.Called(ctx, filter)
	return args[0].([]byte), args.Error(1)

}
func (m *MockS) CreateSeat(car string, num int) error {

	return nil
}
func (m *MockS) UpdateSeats(filter bson.M, update bson.M) error {
	return nil
}

func TestSeatsUserCaseGetASeat(t *testing.T) {
	s := domain.Seat{
		Position:   domain.Position{Car: "A", Num: 1},
		Orderby:    "eric",
		Isorder:    true,
		Updatetime: time.Now(),
	}
	bs, _ := json.Marshal(&s)
	m := MockS{}
	su := NewSeatUseCase(&m)
	m.M.On("GetASeat", context.TODO(), bson.M{}).Return(bs, nil)
	su.GetASeat(context.TODO(), bson.M{})
	m.M.AssertExpectations(t)
}
