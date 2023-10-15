package repository

import (
	"github.com/sakupay-apps/internal/model/dto"
)

type BaseRepository[T any] interface {
	Create(payload *T) (*T, error)
	List(requestPaging dto.PaginationParam, queries ...string) ([]*T, *dto.Paging, error)
	Get(id string) (*T, error)
	Update(id string, payload *T) (*T, error)
	Delete(id string) (*T, error)
}

