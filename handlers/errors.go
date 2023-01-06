package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
)

func CustomErrorHandler(e error, c echo.Context) {
	var err error
	if errors.Is(e, ErrBadRequest) {
		err = c.JSON(http.StatusBadRequest,
			ErrorResponse{
				Message: "Invalid request.",
			},
		)
	} else if echoErr, ok := e.(*echo.HTTPError); ok {
		if echoErr.Code == http.StatusNotFound {
			err = c.JSON(http.StatusNotFound,
				ErrorResponse{
					Message: "Not found.",
				},
			)
		} else {
			err = c.JSON(http.StatusInternalServerError,
				ErrorResponse{
					Message: "Internal server error.",
				},
			)
		}
	} else {
		err = c.JSON(http.StatusInternalServerError,
			ErrorResponse{
				Message: "Internal server error.",
			},
		)
	}

	if err != nil {
		c.Logger().Error(err)
	}
}
