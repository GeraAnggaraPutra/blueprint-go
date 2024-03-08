package module

import (
	"fmt"

	"github.com/GeraAnggaraPutra/blueprint-go/constant"
)

type Pagination struct {
	Search string `query:"search" json:"search"`
	Limit  int64  `query:"limit" json:"limit"`
	Page   int64  `query:"page" json:"page"`
	Order  string `query:"order" json:"order"`
	Sort   string `query:"sort" json:"sort"`
}

func MakeSearch(search string) (bool, string) {
	if search != "" {
		return true, fmt.Sprintf("%%%s%%", search)
	}

	return false, search
}

func MakeLimit(limit int64) int64 {
	if limit <= 0 {
		return constant.DefaultLimit
	}

	return limit
}

func MakeOffset(limit, page int64) int64 {
	if page <= 0 {
		return (1 * limit) - limit
	}

	return (page * limit) - limit
}

func MakeOrder(order, sort string) string {
	if order == "" || sort == "" {
		return constant.DefaultOrder
	}

	return fmt.Sprintf("%s %s", order, sort)
}
