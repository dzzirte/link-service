package auth

import (
	"awesomeproject/configs"
	"awesomeproject/internal/user"
	"awesomeproject/pkg/db"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func bootstrap() (*AuthHandler, sqlmock.Sqlmock, error) {
	database, mock, err := sqlmock.New()
	if err != nil {
		return nil, nil, err
	}
	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: database,
	}))
	if err != nil {
		return nil, nil, err
	}
	userRepo := user.NewUserRepository(&db.Db{
		DB: gormDB,
	})
	handler := AuthHandler{
		Config: &configs.Config{
			Auth: configs.AuthConfig{
				Secret: "secret",
			},
		},
		AuthService: NewAuthService(userRepo),
	}
	return &handler, mock, nil
}

func TestLoginHandlerSuccess(t *testing.T) {
	handler, mock, err := bootstrap()
	rows := sqlmock.NewRows([]string{"email", "password"}).
		AddRow("a2@f.ru", "$2a$10$UYNkY.XbHG.PS1Xcsl3V4O6ATp4/y42tkBYgZ2.RAIfCix3FiFRxy")
	mock.ExpectQuery(".*SELECT.*").
		WithArgs("a2@f.ru", sqlmock.AnyArg()).
		WillReturnRows(rows)
	if err != nil {
		t.Error(err)
		return
	}
	data, _ := json.Marshal(&LoginRequest{
		Email:    "a2@f.ru",
		Password: "2",
	})
	reader := bytes.NewReader(data)
	w := httptest.NewRecorder()
	req, err := http.NewRequest(http.MethodPost, "/auth/login", reader)
	if err != nil {
		t.Error(err)
	}
	handler.Login()(w, req)
	if w.Code != http.StatusOK {
		t.Errorf("got %d, want %d", w.Code, http.StatusOK)
	}
}

func TestRegisterHandlerSuccess(t *testing.T) {
	handler, mock, err := bootstrap()
	rows := sqlmock.NewRows([]string{"email", "password", "name"})
	mock.ExpectQuery("SELECT").WillReturnRows(rows)
	mock.ExpectBegin()
	mock.ExpectQuery("INSERT").WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow("1"))
	mock.ExpectCommit()
	if err != nil {
		t.Error(err)
		return
	}
	data, _ := json.Marshal(&RegisterRequest{
		Email:    "a2ss@f.ru",
		Password: "4",
		Name:     "Вася",
	})
	reader := bytes.NewReader(data)
	w := httptest.NewRecorder()
	req, err := http.NewRequest(http.MethodPost, "/auth/register", reader)
	handler.Register()(w, req)
	if w.Code != http.StatusCreated {
		t.Errorf("got %d, want %d", w.Code, 201)
	}
}
