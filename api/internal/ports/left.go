package ports

import (
	ct "colinahealth_emr-api/internal/application/core/domain/company"
)

// APIPort is the technology neutral port for driving adapters
type APICompanyPort interface {
	GetNewCompany(name ct.Name, contact ct.Contact, website ct.Website, email ct.Email, country ct.Country, state ct.State, zip ct.Zip) (*ct.NewCompanyResp, error)
}