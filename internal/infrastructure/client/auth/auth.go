package auth

type AuthClient interface {
	Login(username, password string) error
}
