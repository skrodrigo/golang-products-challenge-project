package shared

type PageMeta struct {
	Limit   int64 `json:"limit"`
	Offset  int64 `json:"offset"`
	Total   int64 `json:"total"`
	Length  int64 `json:"length"`
	HasNext bool  `json:"has_next"`
	HasPrev bool  `json:"has_prev"`
}

type Page[T any] struct {
	Meta    PageMeta `json:"meta"`
	Records []T      `json:"records"`
}

type MapFunc[TOrigin any, TDest any] func(*TOrigin) TDest

func MapPage[TOrigin any, TDest any](page *Page[TOrigin], mapper MapFunc[TOrigin, TDest]) *Page[TDest] {
	if page == nil || page.Records == nil {
		return &Page[TDest]{}
	}
	records := make([]TDest, len(page.Records))
	for i, record := range page.Records {
		records[i] = mapper(&record)
	}

	return &Page[TDest]{
		Meta:    page.Meta,
		Records: records,
	}
}