package entities

// Product represnts a single product
type Product struct {
	ID    int     `json:"id,omitempty"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}
