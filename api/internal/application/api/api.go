package api

import (
	"colinahealth_emr-api/internal/ports"
	ct "colinahealth_emr-api/internal/application/core/domain/company"
)

// Application implements the APIPort interface
type Application struct {
	dbPort		ports.DbPort
	apiCompany	ApiCompany
}

// NewApplication creates a new Application
func NewApplication(db ports.DbPort, apic ApiCompany) *Application {
	return &Application{dbPort: db, apiCompany: apic}
}

func (apic Application) GetNewCompany(name ct.Name, contact ct.Contact, website ct.Website, email ct.Email, country ct.Country, state ct.State, zip ct.Zip) (*ct.NewCompanyResp, error) {

	answer, err := apic.apiCompany.AddNewCompany(name, contact, website, email, country, state, zip)
	if err != nil {
		return nil, err
	}

	id, err := apic.dbPort.InsertNewCompanyToDb(answer.Uuid.String(), string(answer.Name), string(answer.Contact), string(answer.Website), string(answer.Email), string(answer.Country), string(answer.State), string(answer.Zip))
	if err != nil {
		return nil, err
	}

	return &ct.NewCompanyResp{CompanyID: id, NewCompany: answer}, nil
}