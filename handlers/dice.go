package handlers

import (
	"math/rand"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

func isCorrectRequest(req Request) bool {
	return req.DiceNum > DICE_NUM_MIN && req.Range > RANGE_MIN
}

func generateDiceResult(req *Request) map[int]int {
	result := map[int]int{}

	for cnt := 1; cnt <= req.DiceNum; cnt++ {
		rand.Seed(time.Now().UnixNano())

		// rand.Intn(n) は 0以上n未満の整数を返すので、1足すことで0が出ないようにする
		result[cnt] = rand.Intn(req.Range) + 1
	}

	return result
}

func Dice(c echo.Context) error {
	req := new(Request)
	err := c.Bind(req)
	if err != nil {
		return ErrBadRequest
	}

	if !isCorrectRequest(*req) {
		return ErrBadRequest
	}

	res := new(Response)
	res.Result = generateDiceResult(req)
	return c.JSON(http.StatusOK, res)
}
