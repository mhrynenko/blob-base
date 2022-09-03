package handlers

import (
	"net/http"
	"strconv"

	"blob-base/resources"
	"gitlab.com/distributed_lab/kit/pgdb"
)

const (
	pageParamLimit  = "page[limit]"
	pageParamNumber = "page[number]"
	pageParamCursor = "page[cursor]"
	pageParamOrder  = "page[order]"
)

func GetOffsetLinks(r *http.Request, p pgdb.OffsetPageParams) *resources.Links {
	prevPage := ""
	if p.PageNumber > 0 {
		prevPage = getOffsetLink(r, p.PageNumber-1, p.Limit, p.Order)
	}

	result := resources.Links{
		Next: getOffsetLink(r, p.PageNumber+1, p.Limit, p.Order),
		Prev: prevPage,
		Self: getOffsetLink(r, p.PageNumber, p.Limit, p.Order),
	}

	return &result
}

func getOffsetLink(r *http.Request, pageNumber, limit uint64, order string) string {
	u := r.URL
	query := u.Query()
	query.Set(pageParamNumber, strconv.FormatUint(pageNumber, 10))
	query.Set(pageParamLimit, strconv.FormatUint(limit, 10))
	query.Set(pageParamOrder, order)
	u.RawQuery = query.Encode()
	return u.String()
}
