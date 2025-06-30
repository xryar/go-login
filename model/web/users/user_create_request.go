package web

type UserCreateRequest struct {
	Username string `validate:"required,min=4" json:"username"`
	Fullname string `validate:"required,min=4" json:"fullname"`
	Password string `validate:"required,min=6" json:"password"`
}
