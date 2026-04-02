package auth

import (
	"awesomeproject/internal/user"
	"testing"
)

type MockUserRepository struct{}

func (repo *MockUserRepository) Create(u *user.User) (*user.User, error) {
	return &user.User{
		Email: "aa2@a.ru",
	}, nil
}

func (repo *MockUserRepository) FindByEmail(email string) (*user.User, error) {
	return nil, nil
}

func TestRegisterSucces(t *testing.T) {
	const initialEmail = "aa2@a.ru"
	authService := NewAuthService(&MockUserRepository{})
	email, err := authService.Register(initialEmail, "1", "Даня")
	if err != nil {
		t.Fatal(err)
	}
	if email != initialEmail {
		t.Fatalf("Email %s does not match %s", email, initialEmail)
	}
}
