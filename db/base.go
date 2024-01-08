package db

// DBService is the abstraction of any kind of database
type DBService interface {
	CreateUser(email, password string) error
	GetUserByEmail(email string) (*User, error)
	CreateOrders(userId string, bookIds []string) error
	ListUserOrders(userId string) ([]Order, error)
}
