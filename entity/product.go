package entity

type Product struct {
	Id    int
	Name  string
	Price int
	Stock int
}

func (product Product) StockStatus() string {
	var status string
	if product.Stock < 5 {
		status = "Stok hampir habis"
	} else if product.Stock < 10 {
		status = "Stok terbatas"
	}
	return status
}
