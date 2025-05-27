package user

// User impliments the User interface
type UserRealm struct {
}

type User struct {
	ID				int
	CompanyID		int
	Email			string
	FirstName		string
	LastName		string
	Status			string
	Active			bool
	UpdatedAt		string
	CreateAt		string
}

var Users = []User {}

// NewUser creates a new User
func New() *UserRealm {
	return &UserRealm{}
}

// Create or Add New User Entry
func (userRealm UserRealm) AddUser(companyId int, eMail string, fName string, lName string) (User) {
	var newUser User

	newUser.CompanyID = companyId
	newUser.Email = eMail
	newUser.FirstName = fName
	newUser.LastName = lName
	newUser.Status = "active"

	Users = append(Users, newUser)

	return newUser
}
