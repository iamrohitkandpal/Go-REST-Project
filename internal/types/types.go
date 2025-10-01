package types

type Student struct {
	Id    int64 `json:"id"`
	Name  string `json:"name" validate:"required"`
	Email string `json:"email" validate:"required"`
	Age   int    `json:"age" validate:"required"`
}

type Admin struct {
    Id       int64  `json:"id"`
    Username string `json:"username" validate:"required"`
    Email    string `json:"email" validate:"required,email"`
    Password string `json:"password" validate:"required,min=6"`
}

type AdminResponse struct {
    Id       int64  `json:"id"`
    Username string `json:"username"`
    Email    string `json:"email"`
}