package threads

type Core struct {
	ID            int
	Title         string
	Description   string
	UserID        int
	Like          int
	JumlahComment int
	ImgUrl        string
}

type Bussiness interface {
}

type Data interface {
}
