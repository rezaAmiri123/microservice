package pagnation

type Pagination interface {
	SetSize(sizeQuery string) error
	SetPage(pageQuery string) error
	SetOrderBy(orderByQuery string)
	GetOffset() int
	GetLimit() int
	GetOrderBy() string
	GetPage() int
	GetSize() int
	GetQueryString() string
	GetTotalPages(totalCount int) int
	GetHasMore(totalCount int) bool
}
