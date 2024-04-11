package auth_test

import (
	"fmt"
	"testing"

	"github.com/grin-ch/winter-common/auth"
)

const secret = "mock"

func TestGenerateJWT(t *testing.T) {
	jwt, err := auth.GenerateJWT(60*60*2, secret, "mock", auth.MakeRoleBase(
		"mock",
		"mock/img",
		"mock",
		"127.0.0.1",
	), "salt")
	if err != nil {
		t.Fatal(err)

	}
	c, err := auth.ParseJwtWithSalt(secret, jwt, "salt")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%+v", c)
}
