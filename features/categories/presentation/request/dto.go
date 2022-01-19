package request

import "github.com/dragranzer/capstone-BE-FGD/features/categories"

type Category struct {
	Name string `json:"name" form:"name"`
}

func ToCore(req Category) categories.Core {
	return categories.Core{
		Name: req.Name,
	}
}
