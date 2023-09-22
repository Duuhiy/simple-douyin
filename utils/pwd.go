package utils

import (
	"crypto/sha512"
	"fmt"
	"github.com/anaskhan96/go-password-encoder"
	"strings"
)

func PwdEncode(pwd string) string {
	// Using the default options
	options := &password.Options{10, 100, 32, sha512.New}
	salt, encodedPwd := password.Encode(pwd, options)
	finalPwd := fmt.Sprintf("%s$%s", salt, encodedPwd)
	return finalPwd
}

func PwdCheck(rawpwd, pwd string) bool {
	options := &password.Options{10, 100, 32, sha512.New}
	pwdInfo := strings.Split(pwd, "$")
	check := password.Verify(rawpwd, pwdInfo[0], pwdInfo[1], options)
	return check
}
