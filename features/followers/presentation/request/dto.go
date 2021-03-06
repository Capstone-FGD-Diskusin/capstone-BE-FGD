package request

import "github.com/dragranzer/capstone-BE-FGD/features/followers"

type Follow struct {
	FollowedID  int `json:"followed_id" form:"followed_id"`
	FollowingID int `json:"following_id" form:"following_id"`
}

func ToCore(req Follow, followingId int) followers.Core {

	return followers.Core{
		FollowedID:  req.FollowedID,
		FollowingID: followingId,
	}
}

func ToCoreFollowed(req Follow, followingId int) followers.Core {

	return followers.Core{
		FollowedID:  followingId,
		FollowingID: req.FollowedID,
	}
}

func ToCoreFollowing(req Follow, followingId int) followers.Core {

	return followers.Core{
		FollowedID:  followingId,
		FollowingID: req.FollowingID,
	}
}
