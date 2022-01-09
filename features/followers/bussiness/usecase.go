package bussiness

import (
	"fmt"

	"github.com/dragranzer/capstone-BE-FGD/features/followers"
	"github.com/dragranzer/capstone-BE-FGD/features/users"
)

type followersUsecase struct {
	followerData  followers.Data
	userBussiness users.Bussiness
}

func NewFollowerBussiness(followerData followers.Data, uB users.Bussiness) followers.Bussiness {
	return &followersUsecase{
		followerData:  followerData,
		userBussiness: uB,
	}
}

func (fu *followersUsecase) Follow(data followers.Core) (err error) {
	err = fu.followerData.InsertFollow(data)
	if err != nil {
		return err
	}
	user := users.Core{
		ID: data.FollowedID,
	}
	err = fu.userBussiness.IncrementFol(user)
	return err
}

func (fu *followersUsecase) Unfollow(data followers.Core) (err error) {
	err = fu.followerData.DeleteFollow(data)
	if err != nil {
		return err
	}
	user := users.Core{
		ID: data.FollowedID,
	}
	err = fu.userBussiness.DecrementFol(user)
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
