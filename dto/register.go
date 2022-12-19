package dto

type RegisterDto struct {
	Username  string `form:"username"`
	Password  string `form:"password"`
	Firstname string `form:"firstname"`
	Lastname  string `form:"lastname"`
}

type LoginDto struct {
	Username string `form:"username"`
	Password string `form:"password"`
}
