package modelsrespons

// AdminResponse adalah model respons yang digunakan untuk mengirim data admin ke klien.
type AdminReponse struct {
	Id       uint   `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// AdminRegister adalah model respons yang digunakan untuk mengirim data admin yang berhasil terdaftar ke klien.
type AdminRegister struct {
    Name  string `json:"name"`
	Email string `json:"email"`
	Token string `json:"token"`
}

// AdminLogin adalah model respons yang digunakan untuk mengirim data admin yang berhasil login ke klien.
type AdminLogin struct {
    Name  string `json:"name"`
	Email string `json:"email"`
	Password string `json:"password"`
	Token string `json:"token"`
}		