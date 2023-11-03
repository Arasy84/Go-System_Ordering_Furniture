package domain

// user adalah model domain yang merepresentasikan entitas admin.
type User struct {
	ID        uint
	Name      string
	Email     string
	Password  string
	Address   string
	Phone     uint
}