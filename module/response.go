package module

import (
	"math"

	"github.com/labstack/echo/v4"
)

type responseDataPayload struct {
	Data    interface{} `json:"data"`
	Error   interface{} `json:"error,omitempty"`
	Message string      `json:"message"`
}

type responsePaginatePayload struct {
	Data     interface{}     `json:"data"`
	Error    interface{}     `json:"error,omitempty"`
	Message  string          `json:"message"`
	Paginate paginatePayload `json:"paginate,omitempty"`
}

type paginatePayload struct {
	CurrentPage int64 `json:"current_page"`
	PerPage     int64 `json:"per_page"`
	TotalPage   int   `json:"total_page"`
	TotalData   int64 `json:"total_data"`
}

func ResponseData(c echo.Context, status int, data, err interface{}, msg string) error {
	return c.JSON(status, responseDataPayload{
		Data:    data,
		Error:   err,
		Message: msg,
	})
}

func ResponsePaginate(c echo.Context, status int, data, err interface{}, msg string, paginate paginatePayload) error {
	return c.JSON(status, responsePaginatePayload{
		Data:     data,
		Error:    err,
		Message:  msg,
		Paginate: paginate,
	})
}

func ToPaginatePayload(currentPage, perPage, totalData int64) paginatePayload {
	totalPage := math.Ceil(float64(totalData) / float64(perPage))

	return paginatePayload{
		CurrentPage: currentPage,
		PerPage:     perPage,
		TotalPage:   int(totalPage),
		TotalData:   totalData,
	}
}
