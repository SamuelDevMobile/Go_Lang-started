package entitys

type OrderRepositoryInterface interface { 
	Save(order *Order) error
	GetTotalTransactions() (int, error)
}