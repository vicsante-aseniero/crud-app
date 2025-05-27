package company

import (
	"github.com/google/uuid"
)

// Company impliments the company interface
type Company struct {
}

// NewCompany creates a new Company
func New() *Company {
	return &Company{}
}

type NewCompany struct {
	Uuid		uuid.UUID
	Name		Name
	Contact		Contact
	Website		Website
	Email		Email
	Country		Country
	State		State
	Zip			Zip
}

// GetCompany gets the result of new company entry
type NewCompanyResp struct {
	CompanyID	int32
	NewCompany	NewCompany 
}

func (c Company) AddNewCompany(cn Name, pn Contact, ws Website, em Email, cc Country, sc State, zc Zip) (NewCompany, error) {
	return NewCompany{
		Uuid:		uuid.New(),
		Name:		cn,
		Contact: 	pn,
		Website: 	ws,
		Email: 		em,
		Country:	cc,
		State: 		sc,
		Zip: 		zc,
	}, nil
}