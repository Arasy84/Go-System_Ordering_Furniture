package helper

import "golang.org/x/crypto/bcrypt"

// HashPassword adalah fungsi yang digunakan untuk menghasilkan hash dari kata sandi pengguna.
// Parameter password adalah kata sandi yang akan dihash.
// Fungsi ini mengembalikan string yang merupakan hash dari kata sandi.
func HashPassword(password string) string {
	// Menghasilkan hash dari kata sandi dengan menggunakan bcrypt dan biaya (cost) default.
    // Hasilnya berupa slice byte yang kemudian diubah menjadi string sebelum dikembalikan.
	result, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    return string(result)
}

// ComparePassword adalah fungsi yang digunakan untuk membandingkan kata sandi yang dimasukkan
// oleh pengguna dengan hash kata sandi yang tersimpan.
// Parameter hash adalah hash kata sandi yang tersimpan.
// Parameter password adalah kata sandi yang dimasukkan oleh pengguna.
// Fungsi ini mengembalikan error jika kata sandi tidak cocok, atau nil jika cocok.
func ComparePassword(hash, password string) error {
	// Membandingkan kata sandi yang dimasukkan dengan hash kata sandi yang tersimpan.
    // Jika kata sandi cocok, fungsi ini mengembalikan nil. Jika tidak cocok, fungsi ini mengembalikan error.
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		return err
	}
	return nil
}