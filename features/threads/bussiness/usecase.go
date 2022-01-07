package bussiness

import (
	"fmt"

	"github.com/dragranzer/capstone-BE-FGD/features/followers"
	"github.com/dragranzer/capstone-BE-FGD/features/threads"
)

type threadsUsecase struct {
	followerBussiness followers.Bussiness
	threadData        threads.Data
}

func NewThreadBussiness(fB followers.Bussiness, tD threads.Data) threads.Bussiness {
	return &threadsUsecase{
		followerBussiness: fB,
		threadData:        tD,
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
	return
}
