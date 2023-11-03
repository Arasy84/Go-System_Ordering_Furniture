package config

import (
	"fmt"
	"furniture/models/schema"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

// ConnectDB adalah fungsi untuk menghubungkan ke basis data MySQL dan mengembalikan instance DB GORM.
func ConnectDB() *gorm.DB {
	err := godotenv.Load(filepath.Join(".", ".env"))
	if err != nil {
		fmt.Println("Error loading .env file")
		os.Exit(1)
	}
	// Mengambil informasi koneksi basis data dari variabel lingkungan (environment variables)
	dbUser := os.Getenv("DBUSER")
	dbPass := os.Getenv("DBPASS")
	dbHost := os.Getenv("DBHOST")
	dbPort := os.Getenv("DBPORT")
	dbName := os.Getenv("DBNAME")
	// aiKey  := os.Getenv("AIKEY")

	// Membentuk Data Source Name (DSN) untuk koneksi MySQL
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbUser, dbPass, dbHost, dbPort, dbName)


	var errDB error
	DB, errDB = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if errDB != nil {
		panic("Failed to Connect Database")
	}

	// Melakukan migrasi model basis data
	Migrate()

	fmt.Println("Connected to Database")

	return DB
}

// Migrate adalah fungsi untuk melakukan migrasi basis data dengan menghasilkan tabel-tabel yang sesuai dengan model (schema) yang telah didefinisikan.
func Migrate() {
	DB.AutoMigrate(schema.Admin{}, schema.Product{}, schema.User{}, schema.Order{}, schema.FurnitureRequest{})
}
