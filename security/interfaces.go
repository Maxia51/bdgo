package security

// ISecurity is the interface that wraps the basic security method.
//
// Auth check email and password.
// It returns a boolean value.
// and any error encountered that caused the auth fail.
type ISecurity interface{
	Auth(email string, password string) (bool, error)
}