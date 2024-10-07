package account

type Account struct {
	ID       string
	Username string
	Password string
	Email    string
}

func NewAccount(id, username, password, email string) *Account {
	return &Account{
		ID:       id,
		Username: username,
		Password: password,
		Email:    email,
	}
}

func (a *Account) UpdatePassword(newPassword string) {
	a.Password = newPassword
}
