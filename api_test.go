package kms

import (
	"crypto/md5"
	"encoding/hex"
	"testing"
)

var appkey string
var appsecret string
var pubkey string
var prikey string
var devicecode string
var storage string
var license string

func Test_CreateApp(_t *testing.T) {
	var err error
	appkey, appsecret, pubkey, prikey, err = CreateApp("omonet")
	if nil != err {
		_t.Error(err)
	} else {
		_t.Logf("appkey: %s", appkey)
		_t.Logf("appsecret: %s", appsecret)
		_t.Logf("publickey: %s", pubkey)
		_t.Logf("privatekey: %s", prikey)
	}
}

func Test_MakeLicense(_t *testing.T) {
	hash := md5.New()
	hash.Write([]byte("device code"))
	devicecode = hex.EncodeToString(hash.Sum(nil))
	storage = "{\"key\":\"value\"}"
	var err error
	license, err = MakeLicense(appkey, appsecret, devicecode, storage, 7, pubkey, prikey)
	if nil != err {
		_t.Error(err)
	} else {
		_t.Logf("license: %s", license)
	}
}

func Test_VerifyLicence(_t *testing.T) {
	code, err := VerifyLicense(license, appkey, appsecret, devicecode)
	if nil != err {
		_t.Error(code)
		_t.Error(err)
	}

	code, err = VerifyLicense(license, "12121", appsecret, devicecode)
	if code != 11 {
		_t.Error("test failed")
	}

	code, err = VerifyLicense(license, appkey, "sdd", devicecode)
	if code != 12 {
		_t.Error("test failed")
	}

	code, err = VerifyLicense(license, appkey, appsecret, "sadasd")
	if code != 13 {
		_t.Error("test failed")
	}
}
