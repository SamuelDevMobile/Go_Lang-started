package usecase

import "github.com/SamuelDevMobile/Go_Lang-started/internal/entitys"

type OrderInput struct {
	ID    string
	Price float64
	Tax   float64
}

type OrderOutput struct {
	ID         string
	Price      float64
	Tax        float64
	FinalPrice float64
}

// SOLID - "D" - Dependency Inversion Principle
type CalculateFinalPrice struct {
	OrderRepository entitys.OrderRepositoryInterface
}

func NewCalculateFinalPrice(orderRepository entitys.OrderRepositoryInterface) *CalculateFinalPrice { 
	return &CalculateFinalPrice{ 
		OrderRepository: orderRepository,
	}
}

func (c *CalculateFinalPrice) Execute(input OrderInput) (*OrderOutput, error) {
	order, err := entitys.NewOrder(input.ID, input.Price, input.Tax)
	if err != nil {
		return nil, err
	}
	err = order.CalculateFinalPrice()
	if err != nil {
		return nil, err
	}
	err = c.OrderRepository.Save(order)
	if err != nil {
		return nil, err
	}
	return &OrderOutput{
		ID:         order.ID,
		Price:      order.Price,
		Tax:        order.Tax,
		FinalPrice: order.FinalPrice,
	}, nil
}
