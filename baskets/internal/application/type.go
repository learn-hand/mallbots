package application

type Store struct {
	ID       string
	Name     string
	Location string
}

type Product struct {
	ID      string
	StoreID string
	Name    string
	Price   float64
}
