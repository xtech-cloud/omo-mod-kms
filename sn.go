package kms

import (
	"crypto/md5"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"io"
	"strings"
)

func newSN() (string, error) {
	b := make([]byte, 48)
	if _, err := io.ReadFull(rand.Reader, b); nil != err {
		return "", err
	}
	str := base64.URLEncoding.EncodeToString(b)
	hash := md5.New()
	hash.Write([]byte(str))
	guid := hex.EncodeToString(hash.Sum(nil))
	return strings.ToUpper(guid), nil
}
