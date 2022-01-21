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

func (cu *categoriesUsecase) DeleteCategorybyId(data categories.Core) (err error) {
	err = cu.categoryData.DeleteCategorybyId(data)
	return
}

func (cu *categoriesUsecase) GetAllCategory(data categories.Core) (resp []categories.Core, err error) {
	resp, err = cu.categoryData.SelectAllCategory(data)
	return
}

func (cu *categoriesUsecase) GetCategorybyId(data categories.Core) (resp categories.Core, err error) {
	resp, err = cu.categoryData.SelectCategorybyId(data)
	return
}

func (cu *categoriesUsecase) GetCategorybyName(data categories.Core) (resp categories.Core, err error) {
	resp, err = cu.categoryData.SelectCategorybyName(data)
	return
}
