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
	if err != nil {
		return err
	}
	user.ID = data.FollowingID
	err = fu.userBussiness.IncrementFollowing(user)
	return err
}

func (fu *followersUsecase) Unfollow(data followers.Core) (err error) {
	fmt.Println("unfollow ", data)
	err = fu.followerData.DeleteFollow(data)
	if err != nil {
		return err
	}
	user := users.Core{
		ID: data.FollowedID,
	}
	err = fu.userBussiness.DecrementFol(user)
	if err != nil {
		return err
	}
	user.ID = data.FollowingID
	err = fu.userBussiness.DecrementFollowing(user)
	return err
}

func (fu *followersUsecase) GetFollowing(data followers.Core) (resp []followers.Core, err error) {
	fmt.Println(data)
	// listFollowedID :=
	listFollowedID, err := fu.followerData.SelectFollowing(data)
	fmt.Println(listFollowedID)
	listFollower := []followers.Core{}
	for _, value := range listFollowedID {
		coreUser := users.Core{
			ID: value.FollowedID,
		}
		coreUser, err = fu.userBussiness.GetProfileData(coreUser)
		listFollower = append(listFollower, followers.Core{
			FollowingID:  value.FollowingID,
			FollowedID:   value.FollowedID,
			NameFollowed: coreUser.Username,
		})
	}
	resp = listFollower
	return resp, err
}

func (fu *followersUsecase) GetFollowed(data followers.Core) (resp []followers.Core, err error) {
	fmt.Println(data)
	// listFollowedID :=
	listFollowingID, err := fu.followerData.SelectFollowed(data)
	fmt.Println(listFollowingID)
	listFollower := []followers.Core{}
	for _, value := range listFollowingID {
		coreUser := users.Core{
			ID: value.FollowingID,
		}
		coreUser, err = fu.userBussiness.GetProfileData(coreUser)
		listFollower = append(listFollower, followers.Core{
			FollowingID:  value.FollowingID,
			FollowedID:   value.FollowedID,
			NameFollowed: coreUser.Username,
		})
	}
	resp = listFollower
	return resp, err
}
