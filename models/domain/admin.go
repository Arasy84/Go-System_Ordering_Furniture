package domain

// Admin adalah model domain yang merepresentasikan entitas admin.
type Admin struct {
	ID 			uint
	Name 		string
	Email 		string
	Password 	string
}