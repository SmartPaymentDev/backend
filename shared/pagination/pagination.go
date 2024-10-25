package pagination

import (
	"strconv"

	echo "github.com/labstack/echo/v4"
)

var (
	DefaultPageSize = 20
	MaxPageSize     = 100
	PageVar         = "page"
	PageSizeVar     = "per_page"
)

type Pages struct {
	Pagination
}

type Pagination struct {
	Page       int `json:"page"`
	PerPage    int `json:"per_page"`
	PageCount  int `json:"total_page"`
	TotalCount int `json:"total_data"`
}

func New(page, perPage int) *Pages {
	if perPage <= 0 {
		perPage = DefaultPageSize
	}
	if perPage > MaxPageSize {
		perPage = MaxPageSize
	}
	if page < 1 {
		page = 1
	}

	return &Pages{
		Pagination: Pagination{
			Page:    page,
			PerPage: perPage,
		},
	}
}

func (p *Pages) SetData(totalData int) {
	pageCount := -1
	if totalData >= 0 {
		pageCount = (totalData + p.PerPage - 1) / p.PerPage
	}

	p.PageCount = pageCount
	p.TotalCount = totalData
}

func NewFromRequest(c echo.Context) *Pages {
	page := parseInt(c.QueryParam(PageVar), 1)
	perPage := parseInt(c.QueryParam(PageSizeVar), DefaultPageSize)
	return New(page, perPage)
}

func parseInt(value string, defaultValue int) int {
	if value == "" {
		return defaultValue
	}
	if result, err := strconv.Atoi(value); err == nil {
		return result
	}
	return defaultValue
}
