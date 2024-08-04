package auth

type User struct {
	Username     string
	Password     string
	Secret       string
	TwoFAEnabled bool
}
