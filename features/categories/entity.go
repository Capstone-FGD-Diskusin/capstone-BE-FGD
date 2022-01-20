package categories

type Core struct {
	ID   int
	Name string
}

type Bussiness interface {
	AddCategory(data Core) (err error)
	EditCategory(data Core) (err error)
	DeleteCategorybyId(data Core) (err error)
	GetAllCategory(data Core) (resp []Core, err error)
}

type Data interface {
	InsertCategory(data Core) (err error)
	UpdateCategory(data Core) (err error)
	DeleteCategorybyId(data Core) (err error)
	SelectAllCategory(data Core) (resp []Core, err error)
}
