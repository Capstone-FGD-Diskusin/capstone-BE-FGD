package response

import "github.com/dragranzer/capstone-BE-FGD/features/followers"

type FollowedID struct {
	UserID int
}

func FromCore(res followers.Core) FollowedID {
	return FollowedID{
		UserID: res.FollowedID,
	}
}

func FromCoreSlice(core []followers.Core) []FollowedID {
	var FolArray []FollowedID
	for key := range core {
		FolArray = append(FolArray, FromCore(core[key]))
	}
	return FolArray
}
