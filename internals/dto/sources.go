package dto

type Usercreate struct {
	Name     string `json:"name" validate:"required,max=100" `
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required, min=8"`
}
