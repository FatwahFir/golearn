package categoryRequest

type CategoryCreateRequest struct {
	Name string `json:"name" validate:"required"`
}
