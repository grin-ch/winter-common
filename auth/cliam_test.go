package auth_test

import (
	"fmt"
	"testing"

	"github.com/grin-ch/winter-common/auth"
)

const secret = "mock"

func TestGenerateJWT(t *testing.T) {
	jwt, err := auth.GenerateJWT(60*60*2, secret, "mock", auth.RoleBase{
		Uid:    1024,
		Avatar: "mock/img",
		Sex:    "mock",
		Ip:     "127.0.0.1",
	})
	if err != nil {
		t.Fatal(err)

	}
	c, err := auth.ParseJWT(secret, jwt)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%+v", c)
}
