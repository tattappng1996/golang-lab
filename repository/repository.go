package repository

import (
	"golang-lab/usecase"
	"os"
)

// NewRepository ..
func NewRepository(db *os.File) usecase.Repository {
	return &repository{
		db: db,
	}
}

type repository struct {
	db *os.File
}
