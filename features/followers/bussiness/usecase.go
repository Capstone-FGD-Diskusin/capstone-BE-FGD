package bussiness

import (
	"fmt"

	"github.com/dragranzer/capstone-BE-FGD/features/followers"
)

type followersUsecase struct {
	followerData followers.Data
}

func NewFollowerBussiness(followerData followers.Data) followers.Bussiness {
	return &followersUsecase{
		followerData: followerData,
	}
}

func (fu *followersUsecase) Follow(data followers.Core) (err error) {
	err = fu.followerData.InsertFollow(data)
	return err
}

func (fu *followersUsecase) Unfollow(data followers.Core) (err error) {
	err = fu.followerData.DeleteFollow(data)
	return err
}

func (fu *followersUsecase) GetFollowing(data followers.Core) (resp []followers.Core, err error) {
	fmt.Println(data)
	// listFollowedID :=
	listFollowedID, err := fu.followerData.SelectFollowing(data)
	// fmt.Println(listFollowedID)
	resp = listFollowedID
	return resp, err
}
