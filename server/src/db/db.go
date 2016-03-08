package db

import (
	"database/sql"
)

type DB interface {
	Insert(sql string) (uint64, error)
	Delete(sql string) (uint64, error)
	Update(sql stirng) (uint64, error)
	Select(sql string) (uint64, error)
}
