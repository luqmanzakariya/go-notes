package repository

import "4-section-go-lang-unit-test/entity"

type CategoryRepository interface {
	FindById(id string) *entity.Category
}