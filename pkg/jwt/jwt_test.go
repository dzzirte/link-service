package jwt

import "testing"

func TestJWTCreateAndParse(t *testing.T) {
	const email = "a2aa@f.ru"
	jwtService := NewJWT("Ya8nDfgSFVwFpAIyxKFtIkzt8DN34hl3Aozrc7XfIEd")
	token, err := jwtService.Create(JWTData{
		Email: email,
	})
	if err != nil {
		t.Fatal(err)
	}
	isValid, data := jwtService.Parse(token)
	if !isValid {
		t.Fatal(data)
	}
	if data.Email != email {
		t.Fatalf("Email %s not equal to %s", data.Email, email)
	}
}
