package gormhelpers

import (
	"errors"
	"math"
	"reflect"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/lenna-ai/bni-iproc/config"
	"gorm.io/gorm"
)
type PaginatedResponseStruct struct {
    Data         interface{} `json:"data"`
    TotalCount   int64       `json:"total_all_data"`
    TotalPages   int         `json:"total_pages"`
    CurrentPage  int         `json:"current_page"`
    NextPage     int         `json:"next_page,omitempty"`
    PreviousPage int         `json:"previous_page,omitempty"`
}

func Paginate(r *fiber.Ctx) func(db *gorm.DB) *gorm.DB {
	return func (db *gorm.DB) *gorm.DB {
	  page, _ := strconv.Atoi(r.Query("page"))
	  if page <= 0 {
		page = 1
	  }
  
	  pageSize, _ := strconv.Atoi(r.Query("page_size"))
	  switch {
	  case pageSize > 100:
		pageSize = 100
	  case pageSize <= 0:
		pageSize = 10
	  }
  
	  offset := (page - 1) * pageSize
	  data := db.Offset(offset).Limit(pageSize)
	  return data
	}
}

func CountDataModel(data any,totalCount *int64) error {
	if reflect.ValueOf(data).Kind() != reflect.Ptr {
		return errors.New("please fill data with pointer for reference a table")
	}
	config.DB.Model(data).Count(totalCount)
	return nil
}

func PaginatedResponse(page int,pageSize int,totalCount int64, data any) PaginatedResponseStruct {
    if page <= 0 {
        page = 1
    }

    switch {
    case pageSize > 100:
        pageSize = 100
    case pageSize <= 0:
        pageSize = 10
    }

    totalPages := int(math.Ceil(float64(totalCount) / float64(pageSize)))

    // Prepare the paginated response
    response := PaginatedResponseStruct{
        Data:        data,
        TotalCount:  totalCount,
        TotalPages:  totalPages,
        CurrentPage: page,
    }

    if page < totalPages {
        response.NextPage = page + 1
    }
    if page > 1 {
        response.PreviousPage = page - 1
    }

    return response
}