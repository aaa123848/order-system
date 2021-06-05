package http

import (
	"context"
	"fmt"
	"log"
	"order/domain"
	myerror "order/errors"
	"order/logger"
	"reflect"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

type SeatsHandler struct {
	SeatUseCase domain.SeatUseCase
}

func NewSeatHandler(e *gin.Engine, s domain.SeatUseCase) {
	handler := SeatsHandler{
		SeatUseCase: s,
	}
	e.GET("/seat/:seatid", handler.GetASeat)
	e.POST("/seat/:seatid", handler.UpdateASeats)
}

func (s SeatsHandler) GetASeat(c *gin.Context) {
	ctx := context.Background()
	seatId := c.Param("seatid")
	sarr := strings.Split(seatId, "-")
	num, err := strconv.Atoi(sarr[1])
	if err != nil {
		log.Println(err)
		c.JSON(501, gin.H{"msg": "wrong"})
		return
	}
	filter := bson.M{
		"position": bson.M{
			"car": strings.ToUpper(sarr[0]), "num": num,
		},
	}
	b, err := s.SeatUseCase.GetASeat(
		ctx,
		filter,
	)
	if err != nil {
		logger.GetErrorLog(err)
		c.JSON(501, gin.H{"msg": "wrong"})
		return
	}
	c.JSON(200, string(b))
}

func (s SeatsHandler) UpdateASeats(c *gin.Context) {
	seatId := c.Param("seatid")
	sarr := strings.Split(seatId, "-")
	num, err := strconv.Atoi(sarr[1])
	if err != nil {
		c.JSON(501, gin.H{
			"msg": "Wrong",
		})
	}
	orderer, ok := c.GetPostForm("orderer")
	if !ok {
		c.JSON(501, gin.H{"msg": "NoOrderer"})
		return
	}
	err = s.SeatUseCase.UpdateSeats(strings.ToUpper(sarr[0]), num, orderer)
	if reflect.TypeOf(err) == reflect.TypeOf(myerror.GetSeatOccupiedError()) {
		c.JSON(200, gin.H{
			"msg": "Seats have been occupied",
		})
		return
	}
	c.JSON(200, gin.H{
		"msg": fmt.Sprintf("reserve seats %v - %v successfully", strings.ToUpper(sarr[0]), num),
	})
}
