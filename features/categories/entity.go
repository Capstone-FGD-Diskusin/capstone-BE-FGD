package categories

type Core struct {
	ID                int
	Name              string
	ModeratorID       int
	ModeratorEmail    string
	ModeratorPassword string
	ModeratorName     string
}

type Bussiness interface {
}

type Data interface {
}
