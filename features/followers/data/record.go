package data

import "github.com/dragranzer/capstone-BE-FGD/features/followers"

type Follower struct {
	FollowingID int
	FollowedID  int
}

func fromCore(core followers.Core) Follower {
	return Follower{
		FollowingID: core.FollowingID,
		FollowedID:  core.FollowedID,
	}
}
