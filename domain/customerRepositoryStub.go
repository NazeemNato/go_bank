package domain

type CustomerRepositoryStub struct {
	customers []Customer
}

func (cr CustomerRepositoryStub) FindAll() ([]Customer, error) {
	return cr.customers, nil
}

func NewCustomerRepositoryStub() CustomerRepositoryStub {
	customers := []Customer{
		{"1", "Nazeem", "Alappuzha", "688007", "04/05/1999", "1"},
		{"2", "Jose", "Thampa", "88007", "10/12/1995", "1"},
	}
	return CustomerRepositoryStub{customers}
}
