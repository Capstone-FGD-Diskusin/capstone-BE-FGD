package bussiness

import (
	"github.com/dragranzer/capstone-BE-FGD/features/categories"
)

type categoriesUsecase struct {
	categoryData categories.Data
}

func NewCategoryBussiness(cD categories.Data) categories.Bussiness {
	return &categoriesUsecase{
		categoryData: cD,
	}
}

func (cu *categoriesUsecase) AddCategory(data categories.Core) (err error) {
	err = cu.categoryData.InsertCategory(data)
	return
}

func (cu *categoriesUsecase) EditCategory(data categories.Core) (err error) {
	err = cu.categoryData.UpdateCategory(data)
	return
}
