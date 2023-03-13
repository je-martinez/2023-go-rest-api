package DTOs

type RegisterUserDTO struct {
	Username string `json:"username"`
	Fullname string `json:"fullname"`
	Password string `json:"password"`
}
