package model

import (
	"net/url"
	"strconv"
)

type Pageable struct {
	Limit  int64 `json:"limit"`
	Offset int64 `json:"offset"`
	Order  string
}

func NewPageable(params url.Values) *Pageable {
	limit := parseInt(params.Get("limit"))
	offset := parseInt(params.Get("offset"))
	return &Pageable{
		Limit:  limit,
		Offset: offset,
		Order:  "id desc",
	}
}

func parseInt(param string) int64 {
	pageableDefault := int64(10)
	str, err := strconv.ParseInt(param, 10, 64)
	if err != nil {
		return pageableDefault
	}
	return str
}
