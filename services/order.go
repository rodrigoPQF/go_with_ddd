package services

import (
	"log"

	"github.com/google/uuid"
	"github.com/rodrigoPQF/go_with_ddd/aggregate"
	"github.com/rodrigoPQF/go_with_ddd/domain/customer"
	"github.com/rodrigoPQF/go_with_ddd/domain/customer/memory"
	"github.com/rodrigoPQF/go_with_ddd/domain/product"
	prodmem "github.com/rodrigoPQF/go_with_ddd/domain/product/memory"
)

type OrderConfiguration func(os *OrderService) error

type OrderService struct {
	customers customer.CustomerRepository
	products  product.ProductRepository
}

func NewOrderService(cfgs ...OrderConfiguration) (*OrderService, error) {
	os := &OrderService{}

	for _, cfg := range cfgs {
		err := cfg(os)

		if err != nil {
			return nil, err
		}
	}
	return os, nil
}

func WithCustomerRepository(cr customer.CustomerRepository) OrderConfiguration {
	return func(os *OrderService) error {
		os.customers = cr
		return nil
	}
}

func WithMemoryCustomerRepository() OrderConfiguration {
	cr := memory.New()
	return WithCustomerRepository(cr)
}

func WithMemoryProductRepository(products []aggregate.Product) OrderConfiguration {
	return func(os *OrderService) error {
		pre := prodmem.New()

		for _, p := range products {
			if err := pre.Add(p); err != nil {
				return err
			}
		}
		os.products = pre
		return nil
	}
}

func (o *OrderService) CreateOrder(customerID uuid.UUID, productsIDs []uuid.UUID) (float64, error) {
	c, err := o.customers.Get(customerID)
	if err != nil {
		return 0, err
	}
	log.Println(c)
	var products []aggregate.Product
	var total float64

	for _, id := range productsIDs {
		p, err := o.products.GetByID(id)

		if err != nil {
			return 0, err
		}

		products = append(products, p)
		total += p.GetPrice()
	}
	log.Printf("Customer: %s has order %d products", c.GetID(), len(products))
	return total, nil
}
