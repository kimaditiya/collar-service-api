package customers

type Service interface {
	// ListOfCustomer(findCust FindCustomer) ([]Customers, error)
	ListOfCustomer() ([]Customers, error)
	FindLisOfCustomerByProduct(input FilterCustomerByProduct) ([]Customers, error)
	FindListOfCustomerByOvdDays(input FilterListCustomerByOvd) ([]Customers, error)
	FindListOfCustomerByDueDate(input FilterListCustomerByDueDate) ([]Customers, error)
	// Service Mobile
	ListOfCustomerColl() ([]Customers, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

// func (s *service) ListOfCustomer(find FindCustomer) ([]Customers, error) {
func (s *service) ListOfCustomer() ([]Customers, error) {
	// agreementno := find.AgreementNo

	//nilai baliknya Customer struct
	// loc, err := s.repository.ListOfCustomer(agreementno)
	loc, err := s.repository.ListOfCustomer()

	if err != nil {
		return loc, err
	}

	return loc, nil
}

func (s *service) FindLisOfCustomerByProduct(input FilterCustomerByProduct) ([]Customers, error) {
	product := input.Product

	locFindByProduct, err := s.repository.FindLisOfCustomerByProduct(product)

	if err != nil {
		return locFindByProduct, err
	}

	return locFindByProduct, nil
}
func (s *service) FindListOfCustomerByOvdDays(input FilterListCustomerByOvd) ([]Customers, error) {
	ovdDays := input.OvdDays

	locFindByOvdDays, err := s.repository.FindListOfCustomerByOvd(ovdDays)

	if err != nil {
		return locFindByOvdDays, err
	}
	return locFindByOvdDays, nil
}
func (s *service) FindListOfCustomerByDueDate(input FilterListCustomerByDueDate) ([]Customers, error) {
	dueDate := input.DueDate

	locFindByDueDate, err := s.repository.FindListOfCustomerByDueDate(dueDate)

	if err != nil {
		return locFindByDueDate, err
	}
	return locFindByDueDate, nil
}

func (s *service) ListOfCustomerColl() ([]Customers, error) {

	loc, err := s.repository.ListOfCustomerColl()

	if err != nil {
		return loc, err
	}

	return loc, nil
}
