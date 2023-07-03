package customer

import (
	"errors"

	"github.com/google/uuid"
	"github.com/rodrigoPQF/go_with_ddd/aggregate"
)

var (
	ErrCustomerNotFound    = errors.New("the costumer was found in the repository")
	ErrFailedToAddCostumer = errors.New("failed to add the customer")
	ErrUpdateCustomer      = errors.New("failed to update the customer")
)

type CustomerRepository interface {
	Get(uuid.UUID) (aggregate.Customer, error)
	Add(aggregate.Customer) error
	Update(aggregate.Customer) error
}
