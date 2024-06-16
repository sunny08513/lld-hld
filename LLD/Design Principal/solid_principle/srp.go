package solidprinciple

type User struct {
	FirstName string
	LastName  string
}

func (u *User) GetFullName() string {
	return u.FirstName + " " + u.LastName
}

func (u *User) Save() error {
	// Save user to the database
	// ...
	return nil
}

/*In this case, the User struct has two responsibilities: managing user data and saving it to the database.
To adhere to the Single Responsibility Principle, we should separate these concerns:*/

type UserRepository struct {
	// Database connection or other storage-related fields
}

func (r *UserRepository) Save(u *User) error {
	// Save user to the database
	// ...
	return nil
}

/*Now, the User struct is only responsible for managing user data, while the UserRepository handles database operations*/
