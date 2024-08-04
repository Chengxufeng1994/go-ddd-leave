package auth

import "github.com/Chengxufeng1994/go-ddd-leave/internal/infrastructure/client/auth"

type LocalAuthClient struct {
}

func NewLocalAuthClient() auth.AuthClient {
	return &LocalAuthClient{}
}

// Login implements auth.AuthClient.
func (l *LocalAuthClient) Login(username string, password string) error {
	return nil
}
