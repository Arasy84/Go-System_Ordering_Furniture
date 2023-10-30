package modelsrespons

type UserResponse struct {
	Id uint `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
	Password string `json:"password"`
	Address string `json:"address"`
	Phone uint `json:"phone"`
}

type UserLogin struct {
	ID uint `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
	Password string `json:"password"`
	Token string `json:"token"`
}

type UserCreate struct {
	Name string `json:"name"`
	Email string `json:"email"`
	Password string `json:"password"`
	Address string `json:"address"`
	Phone uint `json:"phone"`
}