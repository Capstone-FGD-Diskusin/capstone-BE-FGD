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

func (a *Follower) toCore() followers.Core {
	return followers.Core{
		FollowingID: a.FollowingID,
		FollowedID:  a.FollowedID,
	}
}

func toCoreList(resp []Follower) []followers.Core {
	a := []followers.Core{}
	for key := range resp {
		a = append(a, resp[key].toCore())
	}
	return a
}
