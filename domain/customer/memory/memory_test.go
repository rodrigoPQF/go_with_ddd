package memory

import (
	"errors"
	"testing"

	"github.com/google/uuid"
	"github.com/rodrigoPQF/go_with_ddd/aggregate"
	"github.com/rodrigoPQF/go_with_ddd/domain/customer"
)

func TestMemory_GetCustom(t *testing.T) {
	type testCase struct {
		name        string
		id          uuid.UUID
		expectedErr error
	}

	cust, err := aggregate.NewCustomer("percy")
	if err != nil {
		t.Fatal(err)
	}

	id := cust.GetID()

	repo := MemoryRepository{
		customers: map[uuid.UUID]aggregate.Customer{
			id: cust,
		},
	}

	testCases := []testCase{
		{
			name:        "no custom by id",
			id:          uuid.MustParse("e080b518-95ee-4457-90ef-6c562576b7ae"),
			expectedErr: customer.ErrCustomerNotFound,
		}, {
			name:        "custom by id",
			id:          id,
			expectedErr: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			_, err := repo.Get(tc.id)
			if !errors.Is(err, tc.expectedErr) {
				t.Errorf("expected error %v, got %v", tc.expectedErr, err)
			}
		})
	}

}
