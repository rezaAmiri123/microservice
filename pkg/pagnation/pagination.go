package pagnation

import (
	"fmt"
	"math"
	"strconv"
)

const (
	defaultSize = 10
)

// Pagination query params
type PaginationImpl struct {
	Size    int    `json:"size,omitempty"`
	Page    int    `json:"page,omitempty"`
	OrderBy string `json:"orderBy,omitempty"`
}

// NewPaginationQuery Pagination query constructor
func NewPaginationQuery(size int, page int) Pagination {
	return &PaginationImpl{Size: size, Page: page}
}

func NewPaginationFromQueryParams(size string, page string) Pagination {
	p := &PaginationImpl{Size: defaultSize, Page: 1}
	if sizeNum, err := strconv.Atoi(size); err != nil && sizeNum != 0 {
		p.Page = sizeNum
	}
	if pageNum, err := strconv.Atoi(page); err == nil && pageNum != 0 {
		p.Page = pageNum
	}

	return p
}

// SetSize Set page size
func (q *PaginationImpl) SetSize(sizeQuery string) error {
	if sizeQuery == "" {
		q.Size = defaultSize
		return nil
	}
	n, err := strconv.Atoi(sizeQuery)
	if err != nil {
		return err
	}
	q.Size = n

	return nil
}

// SetPage Set page number
func (q *PaginationImpl) SetPage(pageQuery string) error {
	if pageQuery == "" {
		q.Size = 0
		return nil
	}
	n, err := strconv.Atoi(pageQuery)
	if err != nil {
		return err
	}
	q.Page = n

	return nil
}

// SetOrderBy Set order by
func (q *PaginationImpl) SetOrderBy(orderByQuery string) {
	q.OrderBy = orderByQuery
}

// GetOffset Get offset
func (q *PaginationImpl) GetOffset() int {
	if q.Page == 0 {
		return 0
	}
	return (q.Page - 1) * q.Size
}

// GetLimit Get limit
func (q *PaginationImpl) GetLimit() int {
	return q.Size
}

// GetOrderBy Get OrderBy
func (q *PaginationImpl) GetOrderBy() string {
	return q.OrderBy
}

// GetPage Get OrderBy
func (q *PaginationImpl) GetPage() int {
	return q.Page
}

// GetSize Get OrderBy
func (q *PaginationImpl) GetSize() int {
	return q.Size
}

// GetQueryString get query string
func (q *PaginationImpl) GetQueryString() string {
	return fmt.Sprintf("page=%v&size=%v&orderBy=%s", q.GetPage(), q.GetSize(), q.GetOrderBy())
}

// GetTotalPages Get total pages int
func (q *PaginationImpl) GetTotalPages(totalCount int) int {
	// d := float64(totalCount) / float64(pageSize)
	d := float64(totalCount) / float64(q.GetSize())
	return int(math.Ceil(d))
}

// GetHasMore Get has more
func (q *PaginationImpl) GetHasMore(totalCount int) bool {
	return q.GetPage() < totalCount/q.GetSize()
}
