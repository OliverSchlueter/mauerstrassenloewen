package authentication

type Store struct {
	globalToken string
}

type Configuration struct {
	GlobalToken string
}

func NewStore(cfg *Configuration) (*Store, error) {
	return &Store{
		globalToken: cfg.GlobalToken,
	}, nil
}

func (s *Store) IsAuthTokenValid(token string) bool {
	return token == s.globalToken
}
