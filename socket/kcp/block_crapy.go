package kcp

import (
	"crypto/sha1"

	kcpgo "github.com/xtaci/kcp-go/v5"
	"golang.org/x/crypto/pbkdf2"
)

const (
	iter   = 1024
	keyLen = 32
)

func MustNewBlockCrypt(passwd, salt []byte) kcpgo.BlockCrypt {
	key := pbkdf2.Key(passwd, salt, iter, keyLen, sha1.New)
	block, err := kcpgo.NewAESBlockCrypt(key)
	if err != nil {
		panic(err)
	}
	return block
}
