package modelsrespons

// UserResponse adalah model respons yang digunakan untuk mengirim data pengguna ke klien.
type UserResponse struct {
	Id uint `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
	Password string `json:"password"`
	Address string `json:"address"`
	Phone uint `json:"phone"`
}

// UserLogin adalah model respons yang digunakan untuk mengirim data pengguna yang berhasil masuk ke klien.
type UserLogin struct {
	ID uint `json:"id"`
	Email string `json:"email"`
	Password string `json:"password"`
	Token string `json:"token"`
}

// UserCreate adalah model respons yang digunakan untuk mengirim data pengguna yang berhasil dibuat ke klien.
type UserCreate struct {
	Name string `json:"name"`
	Email string `json:"email"`
	Password string `json:"password"`
	Address string `json:"address"`
	Phone uint `json:"phone"`
}