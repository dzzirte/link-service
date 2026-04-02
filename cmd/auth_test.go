package main

import (
	"awesomeproject/internal/auth"
	"awesomeproject/internal/user"
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func initDb() *gorm.DB {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}
	db, err := gorm.Open(postgres.Open(os.Getenv("DSN")), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db
}

func initData(db *gorm.DB) {
	db.Create(&user.User{
		Email:    "a2@f.ru",
		Password: "$2a$10$UYNkY.XbHG.PS1Xcsl3V4O6ATp4/y42tkBYgZ2.RAIfCix3FiFRxy",
		Name:     "Даня",
	})
}

func removeData(db *gorm.DB) {
	db.Unscoped().
		Where("email = ?", "a2@f.ru").
		Delete(&user.User{})
}

func TestLoginSuccess(t *testing.T) {
	db := initDb()
	initData(db)
	ts := httptest.NewServer(App(".env"))
	defer ts.Close()
	data, _ := json.Marshal(&auth.LoginRequest{
		Email:    "a2@f.ru",
		Password: "1",
	})
	resp, err := http.Post(ts.URL+"/auth/login", "application/json", bytes.NewBuffer(data))
	if err != nil {
		t.Fatal(err)
	}
	if resp.StatusCode != 200 {
		t.Fatalf("Expected %d got %d", 200, resp.StatusCode)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}
	var resData auth.LoginResponse
	err = json.Unmarshal(body, &resData)
	if err != nil {
		t.Fatal(err)
	}
	if resData.Token == "" {
		t.Fatal("Token is empty")
	}
	removeData(db)
}

func TestLoginFail(t *testing.T) {
	db := initDb()
	initData(db)
	ts := httptest.NewServer(App(".env"))
	defer ts.Close()
	data, _ := json.Marshal(&auth.LoginRequest{
		Email:    "a2ff@f.ru",
		Password: "1",
	})
	resp, err := http.Post(ts.URL+"/auth/login", "application/json", bytes.NewBuffer(data))
	if err != nil {
		t.Fatal(err)
	}
	if resp.StatusCode != 401 {
		t.Fatalf("Expected %d got %d", 401, resp.StatusCode)
	}
	removeData(db)
}
