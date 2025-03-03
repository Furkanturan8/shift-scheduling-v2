package dto

// AddCragRequest Model of CreateCragRequestHandler
type AddUserRequest struct {
	Name    string `json:"name" validate:"required"`
	Desc    string `json:"desc,omitempty"`
	Country string `json:"country" validate:"required"`
}
