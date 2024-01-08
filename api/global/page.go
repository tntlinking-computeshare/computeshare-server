package global

import "github.com/samber/lo"

type Page[T any] struct {
	Page  int32 `json:"pageNum"`
	Size  int32 `json:"pageSize"`
	Total int64 `json:"total"`
	Data  []T   `json:"data"`
}

func Map[T, R any](p *Page[T], iteratee func(item T, index int) R) *Page[R] {

	r := lo.Map[T, R](p.Data, iteratee)

	return &Page[R]{
		Page:  p.Page,
		Size:  p.Size,
		Total: p.Total,
		Data:  r,
	}
}
