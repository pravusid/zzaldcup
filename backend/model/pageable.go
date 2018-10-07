package model

import (
	"golang-server/helper"
	"net/url"
)

type Pageable struct {
	Limit  uint64 `json:"limit"`
	Offset uint64 `json:"offset"`
	Order  string
}

func (pageable *Pageable) Of(params url.Values) *Pageable {
	limit := helper.ParseInt(params.Get("limit"), 10)
	offset := helper.ParseInt(params.Get("offset"), 10)
	pageable.Limit = limit
	pageable.Offset = offset
	pageable.Order = "id desc"
	return pageable
}
