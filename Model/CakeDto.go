package Model

type (
	CakeRequestDto struct {
		Title       string  `json:"title" validate:"required"`
		Description string  `json:"description" validate:"required"`
		Rating      float64 `json:"rating" validate:"required"`
		Image       string  `json:"image" validate:"required"`
	}
)
