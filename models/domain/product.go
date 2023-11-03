package domain

// Product adalah model domain yang merepresentasikan entitas produk.
type Product struct {
	ID             uint
	Name           string
	Description    string
	Price          float64
	Stock      	   int
	Category       string
	Image          string 
}