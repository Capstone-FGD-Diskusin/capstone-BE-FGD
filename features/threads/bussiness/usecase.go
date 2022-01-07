package bussiness

import (
	"fmt"

	"github.com/dragranzer/capstone-BE-FGD/features/followers"
	"github.com/dragranzer/capstone-BE-FGD/features/threads"
)

type threadsUsecase struct {
	followerBussiness followers.Bussiness
}

func NewThreadBussiness(fB followers.Bussiness) threads.Bussiness {
	return &threadsUsecase{
		followerBussiness: fB,
	}
}

func (tu *threadsUsecase) GetThreadHome(data threads.Core) (resp []threads.Core, err error) {
	userID := followers.Core{
		FollowingID: data.OwnerID,
	}
	listID, err := tu.followerBussiness.GetFollowing(userID)
	fmt.Println(listID)
	return
}
