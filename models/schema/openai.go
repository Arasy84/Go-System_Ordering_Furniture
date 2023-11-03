package schema

type FurnitureRequest struct {
	Budget   float64 `json:"budget"`
	Purpose  string  `json:"purpose"`
	Brand    string  `json:"brand"`
	Material string  `json:"material"`
	Color    string  `json:"color"`
}
