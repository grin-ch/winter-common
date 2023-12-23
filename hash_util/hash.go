package hash_util

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
)

func Md5(src []byte) string {
	h := md5.New()
	h.Write(src)
	return hex.EncodeToString(h.Sum(nil))
}

func Sha256(src []byte, secret []byte) string {
	h := hmac.New(sha256.New, secret)
	h.Write(src)
	return hex.EncodeToString(h.Sum(nil))
}
