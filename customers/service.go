package customers

type Service interface {
	ListOfCustomer(findCust FindCustomer) ([]Customers, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) ListOfCustomer(find FindCustomer) ([]Customers, error) {
	agreementno := find.AgreementNo

	//nilai baliknya Customer struct
	loc, err := s.repository.ListOfCustomer(agreementno)

	if err != nil {
		return loc, err
	}

	return loc, nil
}
