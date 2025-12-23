package postgres

type CheckOutRepository interface {
	Checkout(UserID int) error
}
