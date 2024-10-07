package application

type AccountAppService struct {
	accountService *domainaccount.AccountService
}

func NewAccountAppService(accountService *account.AccountService) *AccountAppService {
	return &AccountAppService{
		accountService: accountService,
	}
}

func (s *AccountAppService) CreateAccount(id, username, password, email string) error {
	newAccount := account.NewAccount(id, username, password, email)
	return s.accountService.CreateAccount(newAccount)
}
