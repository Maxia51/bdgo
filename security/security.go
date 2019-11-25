package security

import (
	"fmt"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/maxia51/bdgo/repository"
)

type service struct {
	db repository.IStaffRepo
}

// New instanciate a security service
// Security service implement ISecurity interface
// It return a security pointer
func New(db repository.IStaffRepo) *service {

	return &service{
		db: db,
	}
}

// Auth check the posted credentail
// It return a jwt token and error
func (s *service) Auth(email string, password string) (string, error) {

	staff, err := s.db.GetStaffByEmail(email)
	
	if err != nil {
		return "", err
	}

	// Check password
	if staff.Password != password {
		return "", fmt.Errorf("Incorrect Password")
	}

	token, err := s.createJWT(staff.Id, staff.Role.Name)

	if err != nil {
		return "", err
	}

	return token, nil
}

func (s *service) createJWT(id uint, role string) (string, error) {
	// Create a new token object, specifying signing method and the claims
	// you would like it to contain.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id": id,
		"role": role,
		"nbf":  time.Date(2015, 10, 10, 12, 0, 0, 0, time.UTC).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET_KEY")))

	if err != nil {
		return "", err
	}

	return tokenString, nil
}
