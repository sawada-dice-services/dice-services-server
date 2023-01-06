package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

// 正常系 1d6
func TestDice_1d6(t *testing.T) {
	// setup
	var (
		dice_num   = 1
		dice_range = 6
		mockReq    = `{
			"dice_num": ` + strconv.Itoa(dice_num) + `, 
			"range": ` + strconv.Itoa(dice_range) +
			`}`
	)

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/dice", strings.NewReader(mockReq))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// exec and validate
	if assert.NoError(t, Dice(c)) {
		actually := new(Response)
		err := json.Unmarshal(rec.Body.Bytes(), &actually)
		if err != nil {
			panic(err)
		}

		assert.Exactly(t, http.StatusOK, rec.Code)
		assert.Exactly(t, dice_num, len(actually.Result))
		for _, v := range actually.Result {
			assert.Greater(t, v, RANGE_MIN)
			assert.LessOrEqual(t, v, dice_range)
		}
	}
}

// 正常系 3d100
func TestDice_3d100(t *testing.T) {
	// setup
	var (
		dice_num   = 3
		dice_range = 100
		mockReq    = `{
			"dice_num": ` + strconv.Itoa(dice_num) + `, 
			"range": ` + strconv.Itoa(dice_range) +
			`}`
	)

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/dice", strings.NewReader(mockReq))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// exec and validate
	if assert.NoError(t, Dice(c)) {
		actually := new(Response)
		err := json.Unmarshal(rec.Body.Bytes(), &actually)
		if err != nil {
			panic(err)
		}

		assert.Exactly(t, http.StatusOK, rec.Code)
		assert.Exactly(t, dice_num, len(actually.Result))
		for _, v := range actually.Result {
			assert.Greater(t, v, RANGE_MIN)
			assert.LessOrEqual(t, v, dice_range)
		}
	}
}

// 異常系 ダイスの個数が指定されていない
func TestDice_NoDice(t *testing.T) {
	// setup
	var (
		dice_range = 100
		mockReq    = `{
			"range": ` + strconv.Itoa(dice_range) +
			`}`
	)

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/dice", strings.NewReader(mockReq))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// exec and validate
	err := Dice(c)
	if assert.Error(t, err) {
		assert.ErrorIs(t, ErrBadRequest, err)
	}
}

// 異常系 範囲が指定されていない
func TestDice_NoRange(t *testing.T) {
	// setup
	var (
		dice_num = 3
		mockReq  = `{
				"dice_num": ` + strconv.Itoa(dice_num) + `
			}`
	)

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/dice", strings.NewReader(mockReq))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// exec and validate
	err := Dice(c)
	if assert.Error(t, err) {
		assert.ErrorIs(t, ErrBadRequest, err)
	}
}
