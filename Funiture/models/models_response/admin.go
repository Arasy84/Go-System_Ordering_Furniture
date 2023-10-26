package modelsrespons

type AdminReponse struct {
	Id       uint   `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AdminRegister struct {
    Name  string `json:"name"`
	Email string `json:"email"`
	Token string `json:"token"`
}

type AdminLogin struct {
    Name  string `json:"name"`
	Email string `json:"email"`
	Password string `json:"password"`
	Token string `json:"token"`
}		