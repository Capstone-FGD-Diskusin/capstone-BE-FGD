package categories

type Core struct {
	ID   int
	Name string
}

type Bussiness interface {
	AddCategory(data Core) (err error)
}

type Data interface {
	InsertCategory(data Core) (err error)
}
