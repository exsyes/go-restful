package simple

type Customer interface {
	find(id int) int
}

type CustomerService struct {
	Customer Customer
}

type CustomerImpl struct {
}

func (c *CustomerImpl) find(id int) int {
	return id
}

func NewCustomerImpl() *CustomerImpl {
	return &CustomerImpl{}
}

func NewCustomerService(customer Customer) *CustomerService {
	return &CustomerService{Customer: customer}
}
