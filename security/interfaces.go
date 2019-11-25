package security

type ISecurity interface{
	Auth(email string, password string) (bool, error)
	IsLogged() bool
	GetSession() (Session, error)
}