package security

// ISecurity is the interface that wraps the basic security method.
//
// Auth check email and password.
// It returns the jwt token.
// and any error encountered that caused the auth fail.
type ISecurity interface{
	Auth(email string, password string) (string, error)
}