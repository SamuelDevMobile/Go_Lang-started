package entitys

type OrderRepositoryInterface interface { 
	Save(order *Order) error
	GetTotal() (int, error)
}