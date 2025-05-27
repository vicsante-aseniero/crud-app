package ports

// DbPort is the port for a db adapter
type DbPort interface {
	CloseDbConnection()

	// Company
	InsertNewCompanyToDb(uuid string, name string, contact string, website string, email string, country string, state string, zip string) (int32, error)
}