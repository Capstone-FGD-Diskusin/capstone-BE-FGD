package request

import "github.com/dragranzer/capstone-BE-FGD/features/categories"

type Category struct {
	ID   int    `json:"id" form:"id"`
	Name string `json:"name" form:"name"`
}

func ToCore(req Category) categories.Core {
	return categories.Core{
		ID:   req.ID,
		Name: req.Name,
	}
}
