package bussiness

import (
	"fmt"

	"github.com/dragranzer/capstone-BE-FGD/features/categories"
	"github.com/dragranzer/capstone-BE-FGD/features/followers"
	"github.com/dragranzer/capstone-BE-FGD/features/threads"
)

type threadsUsecase struct {
	followerBussiness followers.Bussiness
	threadData        threads.Data
	categoryBussiness categories.Bussiness
}

func NewThreadBussiness(fB followers.Bussiness, tD threads.Data, cB categories.Bussiness) threads.Bussiness {
	return &threadsUsecase{
		followerBussiness: fB,
		threadData:        tD,
		categoryBussiness: cB,
	}
}

func (tu *threadsUsecase) GetThreadHome(data threads.Core) (resp []threads.Core, err error) {
	userID := followers.Core{
		FollowingID: data.OwnerID,
	}
	temp, err := tu.followerBussiness.GetFollowing(userID)
	if err != nil {
		return
	}
	// fmt.Println(temp)
	listFollowedID := []int{}
	for _, value := range temp {
		listFollowedID = append(listFollowedID, value.FollowedID)
	}
	fmt.Println(listFollowedID)
	data.ListFollowedID = listFollowedID
	resp, err = tu.threadData.SelectThreadHome(data)
	for key, value := range resp {
		fmt.Println(value.CategoryID)
		coreCategory := categories.Core{
			ID: value.CategoryID,
		}
		categoryName, err := tu.categoryBussiness.GetCategorybyId(coreCategory)
		if err != nil {
			continue
		}
		// fmt.Println(categoryName)
		resp[key].CategoryName = categoryName.Name
	}
	if err != nil {
		return
	}
	// fmt.Println(resp)
	return
}

func (tu *threadsUsecase) AddThread(data threads.Core) (err error) {
	categoryCore := categories.Core{
		Name: data.CategoryName,
	}
	categoryID, err := tu.categoryBussiness.GetCategorybyName(categoryCore)
	if err != nil {
		return
	}
	data.CategoryID = categoryID.ID
	err = tu.threadData.InsertThread(data)
	return
}

func (tu *threadsUsecase) GetThreadbyID(data threads.Core) (resp threads.Core, err error) {
	resp, err = tu.threadData.SelectThreadbyID(data)
	if err != nil {
		return
	}
	coreCategory := categories.Core{
		ID: resp.CategoryID,
	}
	categoryName, err := tu.categoryBussiness.GetCategorybyId(coreCategory)
	if err != nil {
		return
	}
	// fmt.Println(categoryName)
	resp.CategoryName = categoryName.Name
	return
}

func (tu *threadsUsecase) IncrementLike(data threads.Core) (err error) {
	err = tu.threadData.UpdateLikebyOne(data)
	return
}

func (tu *threadsUsecase) DecrementLike(data threads.Core) (err error) {
	err = tu.threadData.UpdateMinLikebyOne(data)
	return
}

func (tu *threadsUsecase) IncrementComment(data threads.Core) (err error) {
	err = tu.threadData.UpdateCommentbyOne(data)
	return
}

func (tu *threadsUsecase) DeleteThreadbyId(data threads.Core) (err error) {
	err = tu.threadData.DeleteThreadbyId(data)
	return
}

func (tu *threadsUsecase) SearchThread(data threads.Core) (resp []threads.Core, err error) {
	resp, err = tu.threadData.SearchThread(data)
	for key, value := range resp {
		fmt.Println(value.CategoryID)
		coreCategory := categories.Core{
			ID: value.CategoryID,
		}
		categoryName, err := tu.categoryBussiness.GetCategorybyId(coreCategory)
		if err != nil {
			continue
		}
		// fmt.Println(categoryName)
		resp[key].CategoryName = categoryName.Name
	}
	if err != nil {
		return
	}
	return
}

func (tu *threadsUsecase) GetAllThread(data threads.Core) (resp []threads.Core, err error) {
	resp, err = tu.threadData.SelectThreadAll(data)
	for key, value := range resp {
		fmt.Println(value.CategoryID)
		coreCategory := categories.Core{
			ID: value.CategoryID,
		}
		categoryName, err := tu.categoryBussiness.GetCategorybyId(coreCategory)
		if err != nil {
			continue
		}
		// fmt.Println(categoryName)
		resp[key].CategoryName = categoryName.Name
	}
	if err != nil {
		return
	}
	return
}

func (tu *threadsUsecase) GetThreadUser(data threads.Core) (resp []threads.Core, err error) {
	resp, err = tu.threadData.SelectThreadUser(data)
	for key, value := range resp {
		fmt.Println(value.CategoryID)
		coreCategory := categories.Core{
			ID: value.CategoryID,
		}
		categoryName, err := tu.categoryBussiness.GetCategorybyId(coreCategory)
		if err != nil {
			continue
		}
		// fmt.Println(categoryName)
		resp[key].CategoryName = categoryName.Name
	}
	if err != nil {
		return
	}
	return
}
