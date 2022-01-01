package bussiness

import "github.com/dragranzer/capstone-BE-FGD/features/followers"

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
