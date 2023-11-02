type Storage interface {
	CreateAccount(*model.Account) error
	DeleteAccount(int) error
	UpdateAccount(*model.Account) error
	GetAccountById(int) (*model.Account, error)
}

type PostgresStore struct {
	db *sql.DB
}

