package repository

import (
	db "github.com/MamangRust/monolith-ecommerce-pkg/database/schema"
)

type Repositories struct {
	CategoryQuery   CategoryQueryRepository
	CategoryCommand CategoryCommandRepository
}

func NewRepositories(DB *db.Queries) *Repositories {
	return &Repositories{
		CategoryQuery:   NewCategoryQueryRepository(DB),
		CategoryCommand: NewCategoryCommandRepository(DB),
	}
}
