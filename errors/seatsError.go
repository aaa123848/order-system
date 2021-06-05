package error

type SeatOccupiedError struct {
	msg string
}

func (e *SeatOccupiedError) Error() string {
	return e.msg
}

func GetSeatOccupiedError() *SeatOccupiedError {
	return &SeatOccupiedError{"Seats is Occupied"}
}
