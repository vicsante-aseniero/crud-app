package api

import (
	ct "colinahealth_emr-api/internal/application/core/domain/company"
)

type ApiCompany interface {
	AddNewCompany(name ct.Name, contact ct.Contact, website ct.Website, email ct.Email, country ct.Country, state ct.State, zip ct.Zip) (ct.NewCompany, error)
}