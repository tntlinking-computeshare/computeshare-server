package global

import "github.com/samber/lo"

type Page[T any] struct {
	PageNum  int32 `json:"pageNum"`
	PageSize int32 `json:"pageSize"`
	Total    int64 `json:"total"`
	Data     []T   `json:"data"`
}

func Map[T, R any](p *Page[T], iteratee func(item T, index int) R) *Page[R] {

	r := lo.Map[T, R](p.Data, iteratee)

	return &Page[R]{
		PageSize: p.PageSize,
		PageNum:  p.PageNum,
		Total:    p.Total,
		Data:     r,
	}
}
