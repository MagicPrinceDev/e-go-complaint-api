package response

type Category struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

func FromUseCaseToResponse(category Category) Category {
	return Category{
		ID:          category.ID,
		Name:        category.Name,
		Description: category.Description,
	}
}
