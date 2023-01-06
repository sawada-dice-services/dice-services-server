package handlers

import "errors"

type Request struct {
	DiceNum int `json:"dice_num"`
	Range   int `json:"range"`
}

type Response struct {
	Result map[int]int `json:"result"`
}

type ErrorResponse struct {
	Message string `json:"message"`
}

var ErrBadRequest = errors.New("bad request")
