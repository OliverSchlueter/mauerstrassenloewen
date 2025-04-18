package authentication

type Store struct {
	globalToken string
}

type StoreConfiguration struct {
	GlobalToken string
}

func NewStore(cfg StoreConfiguration) *Store {
	return &Store{
		globalToken: cfg.GlobalToken,
	}
}

func (s *Store) IsAuthTokenValid(token string) bool {
	return token == s.globalToken
}
