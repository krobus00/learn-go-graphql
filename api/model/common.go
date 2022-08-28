package model

import (
	"math"

	"github.com/krobus00/learn-go-graphql/api/constant"
)

const (
	DATASTORE_OPERATION_INSERT = "INSERT"
	DATASTORE_OPERATION_UPDATE = "UPDATE"
	DATASTORE_OPERATION_DELETE = "DELETE"
	DATASTORE_OPERATION_SELECT = "SELECT"
)

type DatastoreSegement struct {
	Collection string
	Operation  string
	Query      string
	Parameter  map[string]interface{}
}

func (req *PaginationRequest) Sanitize() {
	if req.Limit == nil {
		req.Limit = &constant.DEFAULT_PAGINATION_LIMIT
	}
	if req.Limit != nil && *req.Limit <= constant.MIN_PAGINATION_LIMIT {
		req.Limit = &constant.MIN_PAGINATION_LIMIT
	}
	if req.Page == nil {
		req.Page = &constant.DEFAULT_PAGE
	}
	if req.Page != nil && *req.Page <= constant.DEFAULT_PAGE {
		req.Page = &constant.DEFAULT_PAGE
	}

}

func (res *PaginationResponse) BuildResponse(req *PaginationRequest, totalItems uint64) {
	res.TotalItems = int(totalItems)
	res.CurrentPage = *req.Page
	res.Limit = *req.Limit
	if totalItems == 0 {
		res.TotalPages = 0
	} else {
		res.TotalPages = int(math.Ceil(float64(totalItems) / float64(*req.Limit)))
	}
}
