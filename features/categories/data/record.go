package data

import (
	"time"

	"github.com/dragranzer/capstone-BE-FGD/features/categories"
)

type Category struct {
	ID        int
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (a *Category) toCore() categories.Core {
	return categories.Core{
		ID:   int(a.ID),
		Name: a.Name,
	}
}

func ToCoreSlice(data []Category) []categories.Core {
	resp := []categories.Core{}
	for _, value := range data {
		resp = append(resp, value.toCore())
	}
	return resp
}
